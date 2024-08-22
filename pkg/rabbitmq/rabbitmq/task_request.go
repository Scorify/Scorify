package rabbitmq

import (
	"context"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	TaskRequestQueue = "task_request_queue"
	TaskRequestVhost = "task_request_vhost"
)

var (
	// Permissions for minions in task_request vhost
	TaskRequestConfigurePermissions   = regex(TaskRequestQueue)
	TaskRequestMinionWritePermissions = regex("")
	TaskRequestMinionReadPermissions  = regex(TaskRequestQueue)
)

func taskRequestQueue(conn *amqp.Connection) (*amqp.Channel, amqp.Queue, error) {
	ch, err := conn.Channel()
	if err != nil {
		return nil, amqp.Queue{}, err
	}

	q, err := ch.QueueDeclare(
		TaskRequestQueue,
		false,
		false,
		false,
		false,
		nil,
	)

	return ch, q, err
}

func ListenTaskRequest(conn *amqp.Connection, ctx context.Context) error {
	ch, q, err := taskRequestQueue(conn)
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
	if err != nil {
		return err
	}

	for {
		select {
		case <-ctx.Done():
			return nil
		case msg := <-msgs:
			fmt.Println("task_request: ", string(msg.Body))
		}
	}
}

type taskRequestClient struct {
	ch *amqp.Channel
	q  amqp.Queue
}

func TaskRequestClient(conn *amqp.Connection, ctx context.Context) (*taskRequestClient, error) {
	ch, q, err := taskRequestQueue(conn)
	if err != nil {
		return nil, err
	}

	return &taskRequestClient{
		ch: ch,
		q:  q,
	}, nil
}

func (c *taskRequestClient) Close() error {
	return c.ch.Close()
}

func (c *taskRequestClient) Publish(ctx context.Context, message string) error {
	return c.ch.Publish(
		"",
		TaskRequestQueue,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
}
