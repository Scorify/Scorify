package rabbitmq

import (
	"context"
	"fmt"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/redis/go-redis/v9"
	"github.com/scorify/scorify/pkg/cache"
	"github.com/scorify/scorify/pkg/config"
	"github.com/scorify/scorify/pkg/ent"
	"github.com/scorify/scorify/pkg/rabbitmq/types"
	"github.com/scorify/scorify/pkg/structs"
	"github.com/sirupsen/logrus"
)

type RabbitMQConnections struct {
	Heartbeat    *amqp.Connection
	TaskRequest  *amqp.Connection
	TaskResponse *amqp.Connection
	WorkerStatus *amqp.Connection
}

func (r *RabbitMQConnections) Close() error {
	err := r.Heartbeat.Close()
	if err != nil {
		return err
	}

	err = r.TaskRequest.Close()
	if err != nil {
		return err
	}

	err = r.TaskResponse.Close()
	if err != nil {
		return err
	}

	return r.WorkerStatus.Close()
}

func openConnection(vhost string, username string, password string) (*amqp.Connection, error) {
	connStr := fmt.Sprintf(
		"amqp://%s:%s@%s:%d/%s",
		username,
		password,
		config.RabbitMQ.Host,
		config.RabbitMQ.Port,
		vhost,
	)

	conn, err := amqp.Dial(connStr)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func Client(username string, password string) (*RabbitMQConnections, error) {
	heartbeatConn, err := openConnection(HeartbeatVhost, username, password)
	if err != nil {
		return nil, err
	}

	taskRequestsConn, err := openConnection(TaskRequestVhost, username, password)
	if err != nil {
		return nil, err
	}

	taskResponsesConn, err := openConnection(TaskResponseVhost, username, password)
	if err != nil {
		return nil, err
	}

	workerStatusConn, err := openConnection(WorkerStatusVhost, username, password)
	if err != nil {
		return nil, err
	}

	return &RabbitMQConnections{
		Heartbeat:    heartbeatConn,
		TaskRequest:  taskRequestsConn,
		TaskResponse: taskResponsesConn,
		WorkerStatus: workerStatusConn,
	}, nil
}

func Serve(ctx context.Context, taskRequestChan chan *types.TaskRequest, taskResponseChan chan *types.TaskResponse, workerStatusChan chan *types.WorkerStatus, redisClient *redis.Client, entClient *ent.Client) error {
	conn, err := Client(config.RabbitMQ.Server.User, config.RabbitMQ.Server.Password)
	if err != nil {
		return err
	}
	defer conn.Close()

	logrus.Info("Connected to RabbitMQ server")

	go func() {
		heartbeatListener, err := HeartbeatListener(ctx, conn.Heartbeat)
		if err != nil {
			logrus.WithError(err).Fatal("failed to create heartbeat listener")
		}
		defer heartbeatListener.Close()

		for {
			heartbeat, err := heartbeatListener.Consume(ctx)
			if err != nil {
				logrus.WithError(err).Fatal("failed to consume heartbeat")
			}

			heartbeat.Timestamp = time.Now()

			//TODO: switch SetMinionMetrics to use types.Heartbeat
			err = cache.SetMinionMetrics(ctx, heartbeat.MinionID, redisClient, &structs.MinionMetrics{
				MinionID:    heartbeat.MinionID,
				Timestamp:   heartbeat.Timestamp,
				MemoryUsage: heartbeat.MemoryUsage,
				MemoryTotal: heartbeat.MemoryTotal,
				CPUUsage:    heartbeat.CPUUsage,
				Goroutines:  heartbeat.Goroutines,
			})
			if err != nil {
				logrus.WithError(err).Error("failed to set minion metrics")
			}

			//TODO: switch PublishMinionMetrics to use types.Heartbeat
			_, err = cache.PublishMinionMetrics(ctx, redisClient, &structs.MinionMetrics{
				MinionID:    heartbeat.MinionID,
				Timestamp:   heartbeat.Timestamp,
				MemoryUsage: heartbeat.MemoryUsage,
				MemoryTotal: heartbeat.MemoryTotal,
				CPUUsage:    heartbeat.CPUUsage,
				Goroutines:  heartbeat.Goroutines,
			})
			if err != nil {
				logrus.WithError(err).Error("failed to publish minion metrics")
			}
		}
	}()

	go func() {
		taskRequestClient, err := TaskRequestClient(conn.TaskRequest, ctx)
		if err != nil {
			logrus.WithError(err).Fatal("failed to create task request client")
		}

		for {
			select {
			case <-ctx.Done():
				return
			case taskRequest := <-taskRequestChan:
				err := taskRequestClient.Publish(ctx, taskRequest)
				if err != nil {
					logrus.WithError(err).Fatal("failed to send task request")
				}
			}
		}
	}()

	go func() {
		taskResponseListener, err := TaskResponseListener(ctx, conn.TaskResponse)
		if err != nil {
			logrus.WithError(err).Fatal("failed to create task response listener")
		}
		defer taskResponseListener.Close()

		for {
			taskResponse, err := taskResponseListener.Consume(ctx)
			if err != nil {
				logrus.WithError(err).Fatal("failed to consume task response")
			}

			taskResponseChan <- taskResponse
		}
	}()

	go func() {
		workerStatusClient, err := WorkerStatusClient(conn.WorkerStatus)
		if err != nil {
			logrus.WithError(err).Fatal("failed to create worker status client")
		}

		for {
			select {
			case <-ctx.Done():
				return
			case workerStatus := <-workerStatusChan:
				err := workerStatusClient.Publish(ctx, workerStatus)
				if err != nil {
					logrus.WithError(err).Fatal("failed to send worker status")
				}
			}
		}
	}()

	<-ctx.Done()

	return nil
}
