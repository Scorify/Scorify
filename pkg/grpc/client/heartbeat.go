package client

import (
	"context"
	"encoding/json"
	"runtime"
	"time"

	"github.com/scorify/scorify/pkg/grpc/proto"
	"github.com/scorify/scorify/pkg/structs"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	"github.com/sirupsen/logrus"
)

func (c *MinionClient) Heartbeat(ctx context.Context) error {
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
		MinionID:    c.MinionID,
		MemoryUsage: int64(memoryStats.Active),
		MemoryTotal: int64(memoryStats.Total),
		CPUUsage:    cpuUsage[0],
		Goroutines:  int64(runtime.NumGoroutine()),
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
