package rabbitmq

import (
	"context"
	"encoding/json"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/scorify/scorify/pkg/config"
	"github.com/scorify/scorify/pkg/structs"
)

const (
	KothTaskResponseQueue = "koth_task_response_queue"
	KothTaskResponseVhost = "koth_task_response_vhost"
)

var (
	// Permissions for minions in koth_task_response vhosts
	KothTaskResponseConfigurePermissions = regex(KothTaskResponseQueue)
	KothTaskResponseWritePermissions     = regex_amq_default(KothTaskResponseQueue)
	KothTaskResponseReadPermissions      = regex("")
)

func kothTaskResponseQueue(conn *amqp.Connection) (*amqp.Channel, amqp.Queue, error) {
	ch, err := conn.Channel()
	if err != nil {
		return nil, amqp.Queue{}, err
	}

	q, err := ch.QueueDeclare(
		KothTaskResponseQueue,
		false,
		false,
		false,
		false,
		nil,
	)

	return ch, q, err
}

type kothTaskResponseListener struct {
	ch   *amqp.Channel
	msgs <-chan amqp.Delivery
}

func (r *RabbitMQConnections) KothTaskResponseListener(ctx context.Context) (*kothTaskResponseListener, error) {
	ch, q, err := kothTaskResponseQueue(r.KothTaskResponse)
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

	return &kothTaskResponseListener{
		ch:   ch,
		msgs: msgs,
	}, nil
}

func (l *kothTaskResponseListener) Close() error {
	return l.ch.Close()
}

func (l *kothTaskResponseListener) Consume(ctx context.Context) (*structs.KothTaskResponse, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case msg := <-l.msgs:
		var kothTaskResponse structs.KothTaskResponse

		err := json.Unmarshal(msg.Body, &kothTaskResponse)
		if err != nil {
			return nil, err
		}

		return &kothTaskResponse, nil
	}
}

type kothTaskResponseClient struct {
	ch *amqp.Channel
	q  amqp.Queue
}

func (r *RabbitMQConnections) KothTaskResponseClient() (*kothTaskResponseClient, error) {
	ch, q, err := kothTaskResponseQueue(r.KothTaskResponse)
	if err != nil {
		return nil, err
	}

	return &kothTaskResponseClient{
		ch: ch,
		q:  q,
	}, nil
}

func (c *kothTaskResponseClient) Close() error {
	return c.ch.Close()
}

func (c *kothTaskResponseClient) SubmitKothTaskResponse(ctx context.Context, kothTaskResponse *structs.KothTaskResponse) error {
	kothTaskResponse.MinionID = config.Minion.ID

	out, err := json.Marshal(kothTaskResponse)
	if err != nil {
		return err
	}

	return c.ch.PublishWithContext(
		ctx,
		"",
		c.q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        out,
		},
	)
}
