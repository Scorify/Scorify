package rabbitmq

import (
	"context"
	"encoding/json"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/scorify/scorify/pkg/config"
	"github.com/scorify/scorify/pkg/structs"
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

	// Set QoS to limit concurrent tasks per minion
	err = ch.Qos(
		config.RabbitMQ.Minion.QoS, // prefetch count
		0,                           // prefetch size (0 = no limit)
		false,                       // global (false = per consumer)
	)
	if err != nil {
		return nil, err
	}

	msgs, err := ch.ConsumeWithContext(
		ctx,
		q.Name,
		"",
		false, // auto-ack disabled - we'll manually ack after task completion
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return &taskRequestListener{
		ch:   ch,
		msgs: msgs,
	}, nil
}

func (l *taskRequestListener) Close() error {
	return l.ch.Close()
}

func (l *taskRequestListener) Consume(ctx context.Context) (*structs.TaskRequest, func() error, error) {
	select {
	case <-ctx.Done():
		return nil, nil, ctx.Err()
	case msg := <-l.msgs:
		var taskRequest structs.TaskRequest
		err := json.Unmarshal(msg.Body, &taskRequest)
		if err != nil {
			// Nack the message if we can't parse it
			msg.Nack(false, false)
			return nil, nil, err
		}

		// Return the task request and an ack function
		ackFunc := func() error {
			return msg.Ack(false)
		}

		return &taskRequest, ackFunc, nil
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

func (c *taskRequestClient) Publish(ctx context.Context, taskRequest *structs.TaskRequest) error {
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
