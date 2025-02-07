package koth

import (
	"context"
	"time"

	"github.com/scorify/scorify/pkg/config"
	"github.com/scorify/scorify/pkg/rabbitmq/rabbitmq"
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
}
