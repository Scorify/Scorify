package rabbitmq

import (
	"context"
	"encoding/json"
	"fmt"
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

	// Permissions for minions in heartbeat vhosts
	HeartbeatConfigurePermissions   = HeartbeatQueue
	HeartbeatMinionWritePermissions = HeartbeatQueue
	HeartbeatMinionReadPermissions  = ""
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

func ListenHeartbeat(conn *amqp.Connection, ctx context.Context) error {
	ch, q, err := heartbeatQueue(conn)
	if err != nil {
		return err
	}
	defer ch.Close()

	msgs, err := ch.Consume(
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

	for {
		select {
		case <-ctx.Done():
			return nil
		case msg := <-msgs:
			fmt.Println(string(msg.Body))
		}
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

	metrics := structs.MinionMetrics{
		MinionID:    config.Minion.ID,
		MemoryUsage: int64(memoryStats.Active),
		MemoryTotal: int64(memoryStats.Total),
		CPUUsage:    cpuUsage[0],
		Goroutines:  int64(runtime.NumGoroutine()),
	}

	metrics_out, err := json.Marshal(metrics)
	if err != nil {
		return err
	}

	err = c.ch.Publish(
		"",
		c.q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        metrics_out,
		},
	)
	if err != nil {
		return err
	}

	now := time.Now()
	logrus.WithField("time", now).Infof("Heartbeat sent to server in %s", now.Sub(start))

	return nil
}
