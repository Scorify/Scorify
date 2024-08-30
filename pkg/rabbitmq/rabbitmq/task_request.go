package rabbitmq

import (
	"context"
	"encoding/json"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/scorify/scorify/pkg/rabbitmq/types"
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

type taskRequestListener struct {
	ch   *amqp.Channel
	msgs <-chan amqp.Delivery
}

func (r *RabbitMQConnections) TaskRequestListener(ctx context.Context) (*taskRequestListener, error) {
	ch, q, err := taskRequestQueue(r.TaskRequest)
	if err != nil {
		return nil, err
	}
	defer ch.Close()

	msgs, err := ch.ConsumeWithContext(
		ctx,
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	fmt.Println("Listening for task requests")
	for msg := range msgs {
		fmt.Println(msg)
	}

	return &taskRequestListener{
		ch:   ch,
		msgs: msgs,
	}, nil
}

func (l *taskRequestListener) Close() error {
	return l.ch.Close()
}

func (l *taskRequestListener) Consume(ctx context.Context) (*types.TaskRequest, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case msg := <-l.msgs:
		fmt.Println(msg.Body)
		var taskRequest types.TaskRequest
		err := json.Unmarshal(msg.Body, &taskRequest)
		if err != nil {
			return nil, err
		}

		return &taskRequest, nil
	}
}

type taskRequestClient struct {
	ch *amqp.Channel
	q  amqp.Queue
}

func (r *RabbitMQConnections) TaskRequestClient() (*taskRequestClient, error) {
	ch, q, err := taskRequestQueue(r.TaskRequest)
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

func (c *taskRequestClient) Publish(ctx context.Context, taskRequest *types.TaskRequest) error {
	out, err := json.Marshal(taskRequest)
	if err != nil {
		return err
	}

	return c.ch.PublishWithContext(
		ctx,
		"",
		TaskRequestQueue,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        out,
		},
	)
}
