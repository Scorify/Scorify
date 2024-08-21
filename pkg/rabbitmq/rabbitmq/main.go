package rabbitmq

import (
	"context"
	"fmt"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/scorify/scorify/pkg/config"
	"github.com/sirupsen/logrus"
)

func openClient() (*amqp.Connection, error) {
	connStr := fmt.Sprintf(
		"amqp://%s:%s@%s:%d/",
		config.RabbitMQ.Server.User,
		config.RabbitMQ.Server.Password,
		config.RabbitMQ.Host,
		config.RabbitMQ.Port,
	)

	conn, err := amqp.Dial(connStr)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func Serve(ctx context.Context) error {
	conn, err := openClient()
	if err != nil {
		return err
	}
	defer conn.Close()

	logrus.Info("RabbitMQ server started")

	go func() {
		err := ListenTaskResponse(conn, ctx)
		if err != nil {
			logrus.WithError(err).Fatal("failed to listen to task response")
		}
	}()

	go func() {
		err := ListenHeartbeat(conn, ctx)
		if err != nil {
			logrus.WithError(err).Fatal("failed to listen heartbeat")
		}
	}()

	go func() {
		workerStatusClient, err := WorkerStatusClient(conn)
		if err != nil {
			logrus.WithError(err).Fatal("failed to create worker status client")
		}

		ticker := time.NewTicker(time.Second)

		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				err := workerStatusClient.Publish([]byte(time.Now().String()))
				if err != nil {
					logrus.WithError(err).Fatal("failed to send worker status")
				}
			}
		}
	}()

	select {}
}

func Client() error {
	conn, err := openClient()
	if err != nil {
		return err
	}

	defer conn.Close()

	logrus.Info("RabbitMQ client started")

	go func() {
		err := ListenWorkerStatus(conn)
		if err != nil {
			logrus.WithError(err).Error("failed to listen worker status")
		}
	}()

	select {}
}
