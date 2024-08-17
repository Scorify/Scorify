package rabbitmq

import (
	"context"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	SubmitTaskQueue = "submit_task_queue"
)

func submitTaskQueue(conn *amqp.Connection) (*amqp.Channel, amqp.Queue, error) {
	ch, err := conn.Channel()
	if err != nil {
		return nil, amqp.Queue{}, err
	}

	q, err := ch.QueueDeclare(
		SubmitTaskQueue,
		false,
		false,
		false,
		false,
		nil,
	)

	return ch, q, err
}

func ListenSubmitTask(conn *amqp.Connection, ctx context.Context) error {
	ch, q, err := submitTaskQueue(conn)
	if err != nil {
		return err
	}
	defer ch.Close()

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
		return err
	}

	for {
		select {
		case <-ctx.Done():
			return nil
		case msg := <-msgs:
			fmt.Println(string(msg.Body))
		}
	}
}

type submitTaskClient struct {
	ch *amqp.Channel
	q  amqp.Queue
}

func SubmitTaskClient(conn *amqp.Connection, ctx context.Context) (*submitTaskClient, error) {
	ch, q, err := submitTaskQueue(conn)
	if err != nil {
		return nil, err
	}

	return &submitTaskClient{
		ch: ch,
		q:  q,
	}, nil
}

func (c *submitTaskClient) Close() error {
	return c.ch.Close()
}

func (c *submitTaskClient) SubmitTask(ctx context.Context, content string) error {
	return c.ch.Publish(
		"",
		c.q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(content),
		},
	)
}
