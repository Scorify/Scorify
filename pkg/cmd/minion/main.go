package minion

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/scorify/scorify/pkg/checks"
	"github.com/scorify/scorify/pkg/config"
	"github.com/scorify/scorify/pkg/ent/status"
	"github.com/scorify/scorify/pkg/grpc/client"
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
	backOff := time.Second
	heartbeatSuccess := make(chan struct{})

	// Reset backoff on successful heartbeat
	go func() {
		for range heartbeatSuccess {
			backOff = time.Second
		}
	}()

	// Run minion loop first with no backoff
	minionLoop(context.Background(), heartbeatSuccess)
	for {
		time.Sleep(backOff)
		minionLoop(context.Background(), heartbeatSuccess)
		backOff = min(backOff*2, time.Minute)
	}

}

func minionLoop(ctx context.Context, heartbeatSuccess chan struct{}) {
	grpcClient, err := client.Open(ctx)
	if err != nil {
		logrus.WithError(err).Error("encountered error while opening gRPC client")
		return
	}
	defer grpcClient.Close()

	logrus.Info("gRPC client opened successfully")

	err = grpcClient.Enroll(ctx)
	if err != nil {
		logrus.WithError(err).Error("encountered error while enrolling")
		return
	}

	logrus.WithField("minion_id", grpcClient.MinionID).Info("enrolled successfully")

	err = grpcClient.Heartbeat(ctx)
	if err != nil {
		logrus.WithError(err).Error("encountered error while sending heartbeat")
		ctx.Done()
		return
	}
	heartbeatSuccess <- struct{}{}

	go func() {
		ticker := time.NewTicker(10 * time.Second)
		defer ticker.Stop()

		for range ticker.C {
			err := grpcClient.Heartbeat(ctx)
			if err != nil {
				logrus.WithError(err).Error("encountered error while sending heartbeat")
				ctx.Done()
				return
			}
			heartbeatSuccess <- struct{}{}
		}
	}()

	for {
		// recieved score task
		task, err := grpcClient.GetScoreTask(ctx)
		if err != nil {
			logrus.WithError(err).Error("encountered error while getting score task")
			ctx.Done()
			return
		}

		checkDeadline := time.Now().Add(time.Duration(float64(config.Interval) * 0.9))
		submissionDeadline := time.Now().Add(time.Duration(float64(config.Interval) * 0.95))

		// parse UUID
		uuid, err := uuid.Parse(task.GetStatusId())
		if err != nil {
			logrus.WithError(err).Error("encountered error while parsing UUID")
			continue
		}

		// check if source exists
		check, ok := checks.Checks[task.GetSourceName()]
		if !ok {
			logrus.WithField("source", task.GetSourceName()).
				Error("source not found")

			_, err = grpcClient.SubmitScoreTask(ctx, uuid, "source not found", status.StatusDown)
			if err != nil {
				logrus.WithError(err).Error("encountered error while submitting score task")
				ctx.Done()
				return
			}
			continue
		}

		// run check
		go runCheck(checkDeadline, submissionDeadline, grpcClient, uuid, check, task.GetConfig())
	}
}

func runCheck(checkDeadline time.Time, submissionDeadline time.Time, grpcClient *client.MinionClient, uuid uuid.UUID, check checks.Check, config string) {
	checkCtx, checkCancel := context.WithDeadline(context.Background(), checkDeadline)
	submissionCtx, submissionCancel := context.WithDeadline(context.Background(), submissionDeadline)
	defer checkCancel()
	defer submissionCancel()

	// run check and close check context
	err := check.Func(checkCtx, config)
	checkCtx.Done()

	// submit score task and close submission context
	if err != nil {
		_, err = grpcClient.SubmitScoreTask(submissionCtx, uuid, err.Error(), status.StatusDown)
	} else {
		_, err = grpcClient.SubmitScoreTask(submissionCtx, uuid, "", status.StatusUp)
	}
	submissionCtx.Done()

	// log error if submission failed
	if err != nil {
		logrus.WithError(err).Error("encountered error while submitting score task")
	}
}
