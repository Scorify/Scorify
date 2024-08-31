package rabbitmq

import (
	"context"
	"encoding/json"
	"runtime"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/scorify/scorify/pkg/config"
	"github.com/scorify/scorify/pkg/structs"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	"github.com/sirupsen/logrus"
)

const (
	HeartbeatQueue = "heartbeat_queue"
	HeartbeatVhost = "heartbeat_vhost"
)

var (
	// Permissions for minions in heartbeat vhosts
	HeartbeatConfigurePermissions   = regex(HeartbeatQueue)
	HeartbeatMinionWritePermissions = regex_amq_default(HeartbeatQueue)
	HeartbeatMinionReadPermissions  = regex("")
)

func heartbeatQueue(conn *amqp.Connection) (*amqp.Channel, amqp.Queue, error) {
	ch, err := conn.Channel()
	if err != nil {
		return nil, amqp.Queue{}, err
	}

	q, err := ch.QueueDeclare(
		HeartbeatQueue,
		false,
		false,
		false,
		false,
		nil,
	)

	return ch, q, err
}

type heartbeatListener struct {
	ch   *amqp.Channel
	msgs <-chan amqp.Delivery
}

func (r *RabbitMQConnections) HeartbeatListener(ctx context.Context) (*heartbeatListener, error) {
	ch, q, err := heartbeatQueue(r.Heartbeat)
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

	return &heartbeatListener{
		ch:   ch,
		msgs: msgs,
	}, nil
}

func (l *heartbeatListener) Close() error {
	return l.ch.Close()
}

func (l *heartbeatListener) Consume(ctx context.Context) (*structs.Heartbeat, error) {
	select {
	case <-ctx.Done():
		return nil, nil
	case msg := <-l.msgs:
		var heartbeat structs.Heartbeat
		err := json.Unmarshal(msg.Body, &heartbeat)
		if err != nil {
			return nil, err
		}

		return &heartbeat, nil
	}
}

type heartbeatClient struct {
	ch *amqp.Channel
	q  amqp.Queue
}

func (r *RabbitMQConnections) HeartbeatClient() (*heartbeatClient, error) {
	ch, q, err := heartbeatQueue(r.Heartbeat)
	if err != nil {
		return nil, err
	}

	return &heartbeatClient{
		ch: ch,
		q:  q,
	}, nil
}

func (c *heartbeatClient) Close() error {
	return c.ch.Close()
}

func (c *heartbeatClient) SendHeartbeat(ctx context.Context) error {
	start := time.Now()

	cpuUsage, err := cpu.Percent(0, false)
	if err != nil {
		return err
	}

	memoryStats, err := mem.VirtualMemory()
	if err != nil {
		return err
	}

	metrics := structs.Heartbeat{
		MinionID:    config.Minion.ID,
		MemoryUsage: int64(memoryStats.Active),
		MemoryTotal: int64(memoryStats.Total),
		CPUUsage:    cpuUsage[0],
		Goroutines:  int64(runtime.NumGoroutine()),
	}

	out, err := json.Marshal(metrics)
	if err != nil {
		return err
	}

	err = c.ch.PublishWithContext(
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
	if err != nil {
		return err
	}

	now := time.Now()
	logrus.WithField("time", now).Infof("Heartbeat sent to server in %s", now.Sub(start))

	return nil
}
