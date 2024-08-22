package rabbitmq

import (
	"context"
	"fmt"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/scorify/scorify/pkg/config"
	"github.com/sirupsen/logrus"
)

type rabbitMQConnections struct {
	Heartbeat    *amqp.Connection
	TaskRequest  *amqp.Connection
	TaskResponse *amqp.Connection
	WorkerStatus *amqp.Connection
}

func (r *rabbitMQConnections) Close() error {
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

func openClient(username string, password string) (*rabbitMQConnections, error) {
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

	return &rabbitMQConnections{
		Heartbeat:    heartbeatConn,
		TaskRequest:  taskRequestsConn,
		TaskResponse: taskResponsesConn,
		WorkerStatus: workerStatusConn,
	}, nil
}

func Serve(ctx context.Context) error {
	conn, err := openClient(config.RabbitMQ.Server.User, config.RabbitMQ.Server.Password)
	if err != nil {
		return err
	}
	defer conn.Close()

	logrus.Info("Connected to RabbitMQ server")

	go func() {
		err := ListenHeartbeat(conn.Heartbeat, ctx)
		if err != nil {
			logrus.WithError(err).Fatal("failed to listen heartbeat")
		}
	}()

	go func() {
		taskRequestClient, err := TaskRequestClient(conn.TaskRequest, ctx)
		if err != nil {
			logrus.WithError(err).Fatal("failed to create task request client")
		}

		ticker := time.NewTicker(time.Second)

		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				err := taskRequestClient.Publish(ctx, time.Now().String())
				if err != nil {
					logrus.WithError(err).Fatal("failed to send task request")
				}
			}
		}
	}()

	go func() {
		err := ListenTaskResponse(conn.TaskResponse, ctx)
		if err != nil {
			logrus.WithError(err).Fatal("failed to listen to task response")
		}
	}()

	go func() {
		workerStatusClient, err := WorkerStatusClient(conn.WorkerStatus)
		if err != nil {
			logrus.WithError(err).Fatal("failed to create worker status client")
		}

		ticker := time.NewTicker(time.Second)

		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				fmt.Println("sending worker status")
				err := workerStatusClient.Publish([]byte(time.Now().String()))
				if err != nil {
					logrus.WithError(err).Fatal("failed to send worker status")
				}
			}
		}
	}()

	<-ctx.Done()

	return nil
}

func Client(ctx context.Context) error {
	conn, err := openClient(config.RabbitMQ.Minion.User, config.RabbitMQ.Minion.Password)
	if err != nil {
		return err
	}

	logrus.Info("RabbitMQ client started")

	go func() {
		heartbeatClient, err := HeartbeatClient(conn.Heartbeat, ctx)
		if err != nil {
			logrus.WithError(err).Error("failed to create heartbeat client")
		}

		ticker := time.NewTicker(time.Second)

		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				err := heartbeatClient.SendHeartbeat(ctx)
				if err != nil {
					logrus.WithError(err).Error("failed to send heartbeat")
				}
			}
		}
	}()

	go func() {
		err := ListenTaskRequest(conn.TaskRequest, ctx)
		if err != nil {
			logrus.WithError(err).Error("failed to listen task request")
		}
	}()

	go func() {
		taskResponseClient, err := TaskResponseClient(conn.TaskResponse, ctx)
		if err != nil {
			logrus.WithError(err).Error("failed to create task response client")
		}

		ticker := time.NewTicker(time.Second)

		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				err := taskResponseClient.SubmitTaskResponse(ctx, time.Now().String())
				if err != nil {
					logrus.WithError(err).Error("failed to send task response")
				}
			}
		}
	}()

	go func() {
		err := ListenWorkerStatus(conn.WorkerStatus)
		if err != nil {
			logrus.WithError(err).Error("failed to listen worker status")
		}
	}()

	<-ctx.Done()

	return nil
}
