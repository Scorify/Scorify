package rabbitmq

import (
	"context"
	"encoding/json"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/scorify/scorify/pkg/structs"
)

const (
	KothTaskRequestExchange = "koth_task_request_exchange"
	KothTaskRequestVhost    = "koth_task_request_vhost"
)

/* ┌──────────────────────────────────────────────────────────────────┐
 * │                             topic_key:                           │
 * │                         koth.[check_name]                        │
 * │                                                                  │
 * │                 koth.check_1                                     │
 * │                 koth.check_2                                     │
 * │                      ┌─────────────►  Queue 1 ─────► Koth Minion │
 * │                      │                                           │
 * │                      │                                           │
 * │                      │                                           │
 * │                        koth.check_3                              │
 * │ Server ─────► Exchange ────────────►  Queue 2 ─────► Koth Minion │
 * │                                                                  │
 * │                      │                                           │
 * │                      │                                           │
 * │                      │                                           │
 * │                      └─────────────►  Queue 3 ─────► Koth Minion │
 * │                 koth.check_4                                     │
 * └──────────────────────────────────────────────────────────────────┘
 */

var (
	// Permissions for minions in koth_task_request vhosts
	KothTaskRequestConfigurePermissions   = regex_amq_gen(KothTaskRequestExchange)
	KothTaskRequestMinionWritePermissions = regex("amq\\.gen-.*")
	KothTaskRequestMinionReadPermissions  = regex_amq_gen(KothTaskRequestExchange)
)

func kothTaskRequestExchange(conn *amqp.Connection) (*amqp.Channel, error) {
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	err = ch.ExchangeDeclare(
		KothTaskRequestExchange,
		"topic",
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

type kothTaskRequestListener struct {
	ch   *amqp.Channel
	msgs <-chan amqp.Delivery
}

func (r *RabbitMQConnections) KothTaskRequestListener(ctx context.Context, check_names ...string) (*kothTaskRequestListener, error) {
	ch, err := kothTaskRequestExchange(r.KothTaskRequest)
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

	for _, check_name := range check_names {
		err = ch.QueueBind(
			q.Name,
			check_name,
			KothTaskRequestExchange,
			false,
			nil,
		)
		if err != nil {
			return nil, err
		}
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

	return &kothTaskRequestListener{
		ch:   ch,
		msgs: msgs,
	}, nil
}

func (l *kothTaskRequestListener) Close() error {
	return l.ch.Close()
}

func (l *kothTaskRequestListener) Consume(ctx context.Context) (*structs.KothTaskRequest, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case msg := <-l.msgs:
		var task_request structs.KothTaskRequest

		err := json.Unmarshal(msg.Body, &task_request)
		if err != nil {
			return nil, err
		}

		return &task_request, nil
	}
}

type KothTaskRequestClient struct {
	ch *amqp.Channel
}

func (r *RabbitMQConnections) KothTaskRequestClient() (*KothTaskRequestClient, error) {
	ch, err := kothTaskRequestExchange(r.KothTaskRequest)
	if err != nil {
		return nil, err
	}

	return &KothTaskRequestClient{
		ch: ch,
	}, nil
}

func (c *KothTaskRequestClient) Close() error {
	return c.ch.Close()
}

func (c *KothTaskRequestClient) Publish(ctx context.Context, check_name string, task_request *structs.KothTaskRequest) error {
	out, err := json.Marshal(task_request)
	if err != nil {
		return nil
	}

	return c.ch.PublishWithContext(
		ctx,
		KothTaskRequestExchange,
		check_name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        out,
		},
	)
}
