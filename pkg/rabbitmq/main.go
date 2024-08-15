package rabbitmq

import (
	"context"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/scorify/scorify/pkg/config"
	"github.com/sirupsen/logrus"
)

func openClient() (*amqp.Connection, error) {
	connStr := fmt.Sprintf(
		"amqp://%s:%s@%s:%d/",
		config.RabbitMQ.User,
		config.RabbitMQ.Password,
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
		err := ListenSubmitTask(conn, ctx)
		if err != nil {
			logrus.WithError(err).Error("failed to listen submit task")
		}
	}()

	select {}
}

func Client() (*amqp.Connection, error) {
	return openClient()
}
