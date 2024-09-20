package rabbitmq

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/scorify/scorify/pkg/config"
	"github.com/scorify/scorify/pkg/structs"
)

const (
	WorkerEnrollQueue = "task_response_queue"
	WorkerEnrollVhost = "task_response_vhost"
)

var (
	// Permissions for minions in task_response vhost
	WorkerEnrollConfigurePermissions   = regex(WorkerEnrollQueue)
	WorkerEnrollMinionWritePermissions = regex_amq_default(WorkerEnrollQueue)
	WorkerEnrollMinionReadPermissions  = regex("")
)

func workerEnrollQueue(conn *amqp.Connection) (*amqp.Channel, amqp.Queue, error) {
	ch, err := conn.Channel()
	if err != nil {
		return nil, amqp.Queue{}, err
	}

	q, err := ch.QueueDeclare(
		WorkerEnrollQueue,
		false,
		false,
		false,
		false,
		nil,
	)

	return ch, q, err
}

type workerEnrollListener struct {
	ch   *amqp.Channel
	msgs <-chan amqp.Delivery
}

func (r *RabbitMQConnections) WorkerEnrollListener(ctx context.Context) (*workerEnrollListener, error) {
	ch, q, err := workerEnrollQueue(r.WorkerEnroll)
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

	return &workerEnrollListener{
		ch:   ch,
		msgs: msgs,
	}, nil
}

func (r *workerEnrollListener) Close() error {
	return r.ch.Close()
}

func (r *workerEnrollListener) Consume(ctx context.Context) (*structs.WorkerEnroll, error) {
	msg, ok := <-r.msgs
	if !ok {
		return nil, fmt.Errorf("channel closed")
	}

	workerEnroll := &structs.WorkerEnroll{}
	err := json.Unmarshal(msg.Body, workerEnroll)
	if err != nil {
		return nil, err
	}

	return workerEnroll, nil
}

type workerEnrollClient struct {
	ch *amqp.Channel
	q  amqp.Queue
}

func (c *RabbitMQConnections) WorkerEnrollClient() (*workerEnrollClient, error) {
	ch, q, err := workerEnrollQueue(c.WorkerEnroll)
	if err != nil {
		return nil, err
	}

	return &workerEnrollClient{
		ch: ch,
		q:  q,
	}, nil
}

func (c *workerEnrollClient) Close() error {
	return c.ch.Close()
}

func (c *workerEnrollClient) EnrollMinion(ctx context.Context) error {
	hostname, err := os.Hostname()
	if err != nil {
		return err
	}

	workerEnroll := structs.WorkerEnroll{
		MinionID: config.Minion.ID,
		Hostname: hostname,
	}

	body, err := json.Marshal(workerEnroll)
	if err != nil {
		return err
	}

	return c.ch.Publish(
		"",
		c.q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		},
	)
}
