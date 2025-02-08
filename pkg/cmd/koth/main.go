package koth

import (
	"bytes"
	"context"
	"os"
	"time"

	"github.com/scorify/scorify/pkg/config"
	"github.com/scorify/scorify/pkg/ent/minion"
	"github.com/scorify/scorify/pkg/rabbitmq/rabbitmq"
	"github.com/scorify/scorify/pkg/structs"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "koth",
	Short:   "Start scoring koth worker",
	Long:    "Start scoring koth worker",
	Aliases: []string{"k", "worker", "w"},
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		config.InitKoth()
	},

	Run: run,
}

var kothCheckNames []string

func init() {
	Cmd.Flags().StringArrayVar(&kothCheckNames, "check", []string{}, "Name of koth checks")
}

func run(cmd *cobra.Command, args []string) {
	if len(kothCheckNames) == 0 {
		err := cmd.Help()
		if err != nil {
			logrus.WithError(err).Fatal("failed to print help")
		}

		return
	}

	rabbitmqClient, err := rabbitmq.Client(
		config.RabbitMQ.Minion.User,
		config.RabbitMQ.Minion.Password,
	)
	if err != nil {
		logrus.WithError(err).Fatal("failed to create rabbitmq client")
	}
	defer rabbitmqClient.Close()

	backOff := time.Second
	heartbeatSuccess := make(chan struct{})
	ctx := context.Background()

	// Reset backoff on successful heartbeat
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-heartbeatSuccess:
				backOff = time.Second
			}
		}
	}()

	// Run koth loop first with no backoff first
	kothLoop(context.Background(), rabbitmqClient, heartbeatSuccess)
	for {
		time.Sleep(backOff)
		kothLoop(context.Background(), rabbitmqClient, heartbeatSuccess)
		backOff = min(backOff*2, time.Minute)
	}
}

func kothLoop(ctx context.Context, rabbitmqClient *rabbitmq.RabbitMQConnections, heartbeatSuccess chan struct{}) {
	minionCtx, minionCancel := context.WithCancel(ctx)
	defer minionCancel()

	workerEnrollClient, err := rabbitmqClient.WorkerEnrollClient()
	if err != nil {
		logrus.WithError(err).Fatal("failed to create worker enroll client")
	}

	err = workerEnrollClient.EnrollMinion(minionCtx, minion.RoleKoth)
	if err != nil {
		logrus.WithError(err).Fatal("failed to enroll koth")
	}

	workerEnrollClient.Close()

	heartbeatClient, err := rabbitmqClient.HeartbeatClient()
	if err != nil {
		logrus.WithError(err).Fatal("failed to create heartbeat client")
	}
	defer heartbeatClient.Close()

	workerStatusListener, err := rabbitmqClient.WorkerStatusListener(ctx)
	if err != nil {
		logrus.WithError(err).Fatal("failed to create worker status listener")
	}
	defer workerStatusListener.Close()

	workerStatusChannel := workerStatusListener.Consume(minionCtx)

	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	heartbeatSuccess <- struct{}{}

	kothCtx, kothCancel := context.WithCancel(ctx)
	defer kothCancel()

	scoring := true

	go scoreKoth(kothCtx, rabbitmqClient)

	for {
		select {
		case <-minionCtx.Done():
			return
		case <-ticker.C:
			err = heartbeatClient.SendHeartbeat(minionCtx)
			if err != nil {
				logrus.WithError(err).Error("failed to send heartbeat")
				return
			}

			select {
			case <-time.After(5 * time.Second):
				logrus.Error("failed to publish to heartbeatSuccess channel")
				return
			case <-minionCtx.Done():
				return
			case heartbeatSuccess <- struct{}{}:
			}
		case workerStatus := <-workerStatusChannel:
			disabled := workerStatus.Disabled(config.Minion.ID)

			if disabled && scoring {
				kothCancel()
				scoring = false
			} else if !disabled && !scoring {
				go scoreKoth(kothCtx, rabbitmqClient)
				scoring = true
			}
		case <-kothCtx.Done():
			scoring = true
		}
	}
}

func scoreKoth(ctx context.Context, RabbitMQClient *rabbitmq.RabbitMQConnections) {
	kothTaskRquestListener, err := RabbitMQClient.KothTaskRequestListener(ctx)
	if err != nil {
		logrus.WithError(err).Error("failed to create koth task request listener")
		return
	}
	defer kothTaskRquestListener.Close()

	kothTaskResponseClient, err := RabbitMQClient.KothTaskResponseClient()
	if err != nil {
		logrus.WithError(err).Error("failed to create koth task response client")
		return
	}
	defer kothTaskResponseClient.Close()

	for {
		// recieve task request
		taskRequest, err := kothTaskRquestListener.Consume(ctx)
		if err != nil {
			logrus.WithError(err).Error("failed to consume task request")
			return
		}

		content, err := os.ReadFile(taskRequest.Filename)
		if err != nil {
			logrus.WithError(err).Error("failed to read file")
			err = kothTaskResponseClient.SubmitKothTaskResponse(ctx, &structs.KothTaskResponse{
				StatusID: taskRequest.StatusID,
				Error:    err.Error(),
			})
			if err != nil {
				logrus.WithError(err).Error("failed to submit koth task response")
			}
			continue
		}

		err = kothTaskResponseClient.SubmitKothTaskResponse(ctx, &structs.KothTaskResponse{
			StatusID: taskRequest.StatusID,
			Content:  string(bytes.TrimSpace(content)),
		})
		if err != nil {
			logrus.WithError(err).Error("failed to submit koth task response")
		}
	}
}
