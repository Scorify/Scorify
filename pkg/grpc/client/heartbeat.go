package client

import (
	"context"
	"encoding/json"
	"time"

	"github.com/scorify/scorify/pkg/grpc/proto"
	"github.com/scorify/scorify/pkg/structs"
	"github.com/sirupsen/logrus"
)

func (c *MinionClient) Heartbeat(ctx context.Context) error {
	start := time.Now()

	// TODO: handle getting metrics

	metrics := structs.MinionMetrics{
		MinionID:    c.MinionID,
		MemoryUsage: 0,
		MemoryTotal: 0,
		CPUUsage:    0,
		CPUTotal:    0,
		DiskUsage:   0,
		DiskTotal:   0,
	}

	metrics_out, err := json.Marshal(metrics)
	if err != nil {
		return err
	}

	_, err = c.client.Heartbeat(ctx, &proto.HeartbeatRequest{
		MinionId: c.MinionID.String(),
		Metrics:  string(metrics_out),
	})
	if err != nil {
		return err
	}

	now := time.Now()
	logrus.WithField("time", now).Infof("Heartbeat sent to server in %s", now.Sub(start))

	return nil
}
