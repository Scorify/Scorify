package rabbitmq

import (
	"context"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	TaskResponseQueue = "task_response_queue"
	TaskResponseVhost = "task_response_vhost"

	// Permissions for minions in task_response vhost
	TaskResponseConfigurePermissions   = TaskResponseQueue
	TaskResponseMinionWritePermissions = TaskResponseQueue
	TaskResponseMinionReadPermissions  = ""
)

func taskResponseQueue(conn *amqp.Connection) (*amqp.Channel, amqp.Queue, error) {
	ch, err := conn.Channel()
	if err != nil {
		return nil, amqp.Queue{}, err
	}

	q, err := ch.QueueDeclare(
		TaskResponseQueue,
		false,
		false,
		false,
		false,
		nil,
	)

	return ch, q, err
}

func ListenTaskResponse(conn *amqp.Connection, ctx context.Context) error {
	ch, q, err := taskResponseQueue(conn)
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

type taskResponseClient struct {
	ch *amqp.Channel
	q  amqp.Queue
}

func TaskResponseClient(conn *amqp.Connection, ctx context.Context) (*taskResponseClient, error) {
	ch, q, err := taskResponseQueue(conn)
	if err != nil {
		return nil, err
	}

	return &taskResponseClient{
		ch: ch,
		q:  q,
	}, nil
}

func (c *taskResponseClient) Close() error {
	return c.ch.Close()
}

func (c *taskResponseClient) SubmitTaskResponse(ctx context.Context, content string) error {
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
