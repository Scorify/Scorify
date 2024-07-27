package client

import (
	"context"
	"os"

	"github.com/scorify/scorify/pkg/grpc/proto"
)

func (c *MinionClient) Enroll(ctx context.Context) error {
	hostname, err := os.Hostname()
	if err != nil {
		return err
	}

	_, err = c.client.Enroll(ctx, &proto.EnrollRequest{
		MinionId: c.MinionID.String(),
		Name:     hostname,
	})

	return err
}
