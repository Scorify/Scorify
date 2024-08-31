package rabbitmq

import (
	"context"
	"encoding/json"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/scorify/scorify/pkg/structs"
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

func (r *RabbitMQConnections) WorkerStatusListener(ctx context.Context) (*workerStatusListener, error) {
	ch, err := workerStatusExchange(r.WorkerStatus)
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

func (l *workerStatusListener) Consume(ctx context.Context) <-chan *structs.WorkerStatus {
	out := make(chan *structs.WorkerStatus)

	go func() {
		defer close(out)

		for {
			select {
			case <-ctx.Done():
				return
			case msg := <-l.msgs:
				var workerStatus structs.WorkerStatus
				err := json.Unmarshal(msg.Body, &workerStatus)
				if err != nil {
					continue
				}

				select {
				case out <- &workerStatus:
				case <-ctx.Done():
					return
				}
			}
		}
	}()

	return out
}

type workerStatusClient struct {
	ch *amqp.Channel
}

func (r *RabbitMQConnections) WorkerStatusClient() (*workerStatusClient, error) {
	ch, err := workerStatusExchange(r.WorkerStatus)
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

func (c *workerStatusClient) Publish(ctx context.Context, workerStatus *structs.WorkerStatus) error {
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
