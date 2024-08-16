package rabbitmq

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	WorkerStatusExchange = "worker_status_exchange"
)

func workerStatusExchange(conn *amqp.Connection) (*amqp.Channel, error) {
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	err = ch.ExchangeDeclare(
		WorkerStatusExchange,
		"fanout",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return ch, nil
}

func ListenWorkerStatus(conn *amqp.Connection) error {
	ch, err := workerStatusExchange(conn)
	if err != nil {
		return err
	}

	q, err := ch.QueueDeclare(
		"",
		false,
		false,
		true,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	err = ch.QueueBind(
		q.Name,
		"",
		WorkerStatusExchange,
		false,
		nil,
	)
	if err != nil {
		return err
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

	for msg := range msgs {
		fmt.Println(string(msg.Body))
	}

	return err
}

type workerStatusClient struct {
	ch *amqp.Channel
}

func WorkerStatusClient(conn *amqp.Connection) (*workerStatusClient, error) {
	ch, err := workerStatusExchange(conn)
	if err != nil {
		return nil, err
	}

	return &workerStatusClient{
		ch: ch,
	}, nil
}

func (c *workerStatusClient) Close() error {
	return c.ch.Close()
}

func (c *workerStatusClient) Publish(msg []byte) error {
	return c.ch.Publish(
		WorkerStatusExchange,
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        msg,
		},
	)
}
