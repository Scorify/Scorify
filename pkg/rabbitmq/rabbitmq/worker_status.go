package rabbitmq

import (
	"context"
	"encoding/json"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/scorify/scorify/pkg/rabbitmq/types"
)

const (
	WorkerStatusExchange = "worker_status_exchange"
	WorkerStatusVhost    = "worker_status_vhost"
)

var (
	// Permissions for minions in worker status vhosts
	WorkerStatusConfigurePermissions   = regex_amq_gen(WorkerStatusExchange)
	WorkerStatusMinionWritePermissions = regex("amq\\.gen-.*")
	WorkerStatusMinionReadPermissions  = regex_amq_gen(WorkerStatusExchange)
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

type workerStatusListener struct {
	ch   *amqp.Channel
	msgs <-chan amqp.Delivery
}

func WorkerStatusListener(ctx context.Context, conn *amqp.Connection) (*workerStatusListener, error) {
	ch, err := workerStatusExchange(conn)
	if err != nil {
		return nil, err
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
		return nil, err
	}

	err = ch.QueueBind(
		q.Name,
		"",
		WorkerStatusExchange,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

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

	return &workerStatusListener{
		ch:   ch,
		msgs: msgs,
	}, nil
}

func (l *workerStatusListener) Close() error {
	return l.ch.Close()
}

func (l *workerStatusListener) Consume(ctx context.Context) (*types.WorkerStatus, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case msg, ok := <-l.msgs:
		if !ok {
			return nil, fmt.Errorf("worker status channel closed")
		}

		var workerStatus types.WorkerStatus
		err := json.Unmarshal(msg.Body, &workerStatus)
		if err != nil {
			return nil, err
		}

		return &workerStatus, nil
	}
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

func (c *workerStatusClient) Publish(ctx context.Context, workerStatus *types.WorkerStatus) error {
	out, err := json.Marshal(workerStatus)
	if err != nil {
		return err
	}

	return c.ch.PublishWithContext(
		ctx,
		WorkerStatusExchange,
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        out,
		},
	)
}
