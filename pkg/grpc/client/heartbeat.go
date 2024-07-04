package client

import (
	"context"
	"time"

	"github.com/scorify/scorify/pkg/grpc/proto"
	"github.com/sirupsen/logrus"
)

func (c *MinionClient) Heartbeat(ctx context.Context) error {
	start := time.Now()

	_, err := c.client.Heartbeat(ctx, &proto.HeartbeatRequest{
		MinionId: c.MinionID.String(),
	})
	if err != nil {
		return err
	}

	now := time.Now()
	logrus.WithField("time", now).Infof("Heartbeat sent to server in %s", now.Sub(start))

	return nil
}
