package rabbitmq

import (
	"context"
	"encoding/json"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/scorify/scorify/pkg/rabbitmq/management/types"
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

func ListenWorkerStatus(conn *amqp.Connection, workerStatusHandler func(*types.WorkerStatus), ctx context.Context) error {
	ch, err := workerStatusExchange(conn)
	if err != nil {
		return err
	}
	defer ch.Close()

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
		return err
	}

	for msg := range msgs {
		var workerStatus types.WorkerStatus
		err := json.Unmarshal(msg.Body, &workerStatus)
		if err != nil {
			return err
		}

		workerStatusHandler(&workerStatus)
	}

	return nil
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
