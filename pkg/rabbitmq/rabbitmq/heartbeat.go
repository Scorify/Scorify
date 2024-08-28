package rabbitmq

import (
	"context"
	"encoding/json"
	"fmt"
	"runtime"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/scorify/scorify/pkg/config"
	"github.com/scorify/scorify/pkg/rabbitmq/types"
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

func HeartbeatListener(ctx context.Context, conn *amqp.Connection) (*heartbeatListener, error) {
	ch, q, err := heartbeatQueue(conn)
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

func (l *heartbeatListener) Consume(ctx context.Context) (*types.Heartbeat, error) {
	select {
	case <-ctx.Done():
		return nil, nil
	case msg, ok := <-l.msgs:
		if !ok {
			return nil, fmt.Errorf("heartbeat channel closed")
		}

		var heartbeat types.Heartbeat
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

func HeartbeatClient(conn *amqp.Connection, ctx context.Context) (*heartbeatClient, error) {
	ch, q, err := heartbeatQueue(conn)
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

	metrics := types.Heartbeat{
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
