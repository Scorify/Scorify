package rabbitmq

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/scorify/scorify/pkg/config"
	"github.com/sirupsen/logrus"
)

func Serve() {
	connStr := fmt.Sprintf(
		"amqp://%s:%s@%s:%d/",
		config.RabbitMQ.User,
		config.RabbitMQ.Password,
		config.RabbitMQ.Host,
		config.RabbitMQ.Port,
	)

	conn, err := amqp.Dial(connStr)
	if err != nil {
		logrus.WithError(err).Fatal("failed to connect to RabbitMQ")
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		logrus.WithError(err).Fatal("failed to open a channel")
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		logrus.WithError(err).Fatal("failed to declare a queue")
	}

	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("Hello World"),
		},
	)
	if err != nil {
		logrus.WithError(err).Fatal("failed to publish a message")
	}

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		logrus.WithError(err).Fatal("failed to consume a message")
	}

	for msg := range msgs {
		logrus.Infof("Received a message: %s", msg.Body)
	}

	logrus.Info("RabbitMQ server started")
	select {}
}
