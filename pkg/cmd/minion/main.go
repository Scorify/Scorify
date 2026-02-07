package minion

import (
	"context"
	"sync/atomic"
	"time"

	"github.com/google/uuid"
	"github.com/scorify/scorify/pkg/checks"
	"github.com/scorify/scorify/pkg/config"
	"github.com/scorify/scorify/pkg/ent/minion"
	"github.com/scorify/scorify/pkg/ent/status"
	"github.com/scorify/scorify/pkg/rabbitmq/rabbitmq"
	"github.com/scorify/scorify/pkg/structs"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "minion",
	Short:   "Start scoring minion worker",
	Long:    "Start scoring minion worker",
	Aliases: []string{"m", "worker", "w"},
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		config.InitMinion()
	},

	Run: run,
}

func run(cmd *cobra.Command, args []string) {
	for {
		logrus.Info("Starting minion worker")
		mainLoop()
		logrus.Info("minion worker stopped")
		time.Sleep(time.Second * 5)
	}
}

func mainLoop() {
	rabbitmqClient, err := rabbitmq.Client(
		config.RabbitMQ.Minion.User,
		config.RabbitMQ.Minion.Password,
	)
	if err != nil {
		logrus.WithError(err).Error("failed to create rabbitmq client")
		return
	}
	defer rabbitmqClient.Close()

	var backOff atomic.Int64
	backOff.Store(int64(time.Second))
	heartbeatSuccess := make(chan struct{}, 1)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Reset backoff on successful heartbeat
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-heartbeatSuccess:
				backOff.Store(int64(time.Second))
			}
		}
	}()

	// Run minion loop, retry with backoff on failure
	// Exit mainLoop when backoff reaches max (connection likely unhealthy)
	minionLoop(ctx, rabbitmqClient, heartbeatSuccess)
	for {
		currentBackOff := time.Duration(backOff.Load())
		if currentBackOff >= time.Minute {
			logrus.Warn("max backoff reached, reconnecting to rabbitmq")
			return
		}

		time.Sleep(currentBackOff)
		minionLoop(ctx, rabbitmqClient, heartbeatSuccess)

		newBackOff := backOff.Load() * 2
		backOff.Store(min(newBackOff, int64(time.Minute)))
	}
}

func minionLoop(ctx context.Context, rabbitmqClient *rabbitmq.RabbitMQConnections, heartbeatSuccess chan struct{}) {
	minionCtx, minionCancel := context.WithCancel(ctx)
	defer minionCancel()

	workerEnrollClient, err := rabbitmqClient.WorkerEnrollClient()
	if err != nil {
		logrus.WithError(err).Fatal("failed to create worker enroll client")
	}

	err = workerEnrollClient.EnrollMinion(minionCtx, minion.RoleService)
	if err != nil {
		logrus.WithError(err).Fatal("failed to enroll minion")
	}
	workerEnrollClient.Close()

	heartbeatClient, err := rabbitmqClient.HeartbeatClient()
	if err != nil {
		logrus.WithError(err).Fatal("failed to create heartbeat client")
	}
	defer heartbeatClient.Close()

	workerStatusListener, err := rabbitmqClient.WorkerStatusListener(minionCtx)
	if err != nil {
		logrus.WithError(err).Fatal("failed to create worker status listener")
	}
	defer workerStatusListener.Close()

	workerStatusChannel := workerStatusListener.Consume(minionCtx)

	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	err = heartbeatClient.SendHeartbeat(minionCtx)
	if err != nil {
		logrus.WithError(err).Error("failed to send heartbeat")
		return
	}

	heartbeatSuccess <- struct{}{}

	scoringCtx, scoringCtxCancel := context.WithCancel(minionCtx)
	defer func() { scoringCtxCancel() }()

	scoring := true

	go score(scoringCtx, rabbitmqClient)

	for {
		select {
		case <-minionCtx.Done():
			return
		case <-ticker.C:
			err := heartbeatClient.SendHeartbeat(minionCtx)
			if err != nil {
				logrus.WithError(err).Error("failed to send heartbeat")
				return
			}

			timeout := time.NewTimer(5 * time.Second)
			select {
			case <-timeout.C:
				logrus.Error("failed to publish to heartbeatSuccess channel")
				return
			case <-minionCtx.Done():
				timeout.Stop()
				return
			case heartbeatSuccess <- struct{}{}:
				timeout.Stop()
			}

		case workerStatus := <-workerStatusChannel:
			disabled := workerStatus.Disabled(config.Minion.ID)

			if disabled && scoring {
				scoringCtxCancel()
			} else if !disabled && !scoring {
				scoringCtx, scoringCtxCancel = context.WithCancel(minionCtx)
				go score(scoringCtx, rabbitmqClient)
				scoring = true
			}

		case <-scoringCtx.Done():
			scoring = false
		}
	}
}

func score(ctx context.Context, rabbitmqClient *rabbitmq.RabbitMQConnections) {
	taskRequestListener, err := rabbitmqClient.TaskRequestListener(ctx)
	if err != nil {
		logrus.WithError(err).Fatal("failed to create task request listener")
	}
	defer taskRequestListener.Close()

	taskResponseClient, err := rabbitmqClient.TaskResponseClient()
	if err != nil {
		logrus.WithError(err).Fatal("failed to create task response client")
	}
	defer taskResponseClient.Close()

	for {
		// received score task
		task, ackFunc, err := taskRequestListener.Consume(ctx)
		if err != nil {
			logrus.WithError(err).Error("failed to consume task request")
			return
		}

		checkDeadline := time.Now().Add(time.Duration(float64(config.Interval) * 0.8))
		submissionDeadline := time.Now().Add(time.Duration(float64(config.Interval) * 0.85))

		// check if source exists
		check, ok := checks.Checks[task.SourceName]
		if !ok {
			logrus.WithField("source", task.SourceName).
				Error("source not found")

			err = taskResponseClient.SubmitTaskResponse(
				ctx,
				&structs.TaskResponse{
					StatusID: task.StatusID,
					MinionID: config.Minion.ID,
					Status:   status.StatusDown,
					Error:    "source not found",
				},
			)
			if err != nil {
				logrus.WithError(err).Error("encountered error while submitting score task")
				return
			}

			// Ack the message even though the source was not found
			if err := ackFunc(); err != nil {
				logrus.WithError(err).Error("failed to ack message")
			}
			continue
		}

		// run check
		go func(checkDeadline time.Time, submissionDeadline time.Time, statusID uuid.UUID, check checks.Check, taskConfig string, ackFunc func() error) {
			checkCtx, checkCancel := context.WithDeadline(context.Background(), checkDeadline)
			submissionCtx, submissionCancel := context.WithDeadline(context.Background(), submissionDeadline)
			defer checkCancel()
			defer submissionCancel()

			// run check
			err := check.Func(checkCtx, taskConfig)

			// submit score task
			if err != nil {
				err = taskResponseClient.SubmitTaskResponse(
					submissionCtx,
					&structs.TaskResponse{
						StatusID: statusID,
						MinionID: config.Minion.ID,
						Status:   status.StatusDown,
						Error:    err.Error(),
					},
				)
			} else {
				err = taskResponseClient.SubmitTaskResponse(
					submissionCtx,
					&structs.TaskResponse{
						StatusID: statusID,
						MinionID: config.Minion.ID,
						Status:   status.StatusUp,
					},
				)
			}

			// log error if submission failed
			if err != nil {
				logrus.WithError(err).Error("encountered error while submitting score task")
			}

			// Ack the message after task completion
			if err := ackFunc(); err != nil {
				logrus.WithError(err).Error("failed to ack message")
			}
		}(checkDeadline, submissionDeadline, task.StatusID, check, task.Config, ackFunc)
	}
}
