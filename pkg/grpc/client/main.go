package client

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/scorify/scorify/pkg/config"
	"github.com/scorify/scorify/pkg/grpc/proto"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type MinionClient struct {
	MinionID uuid.UUID
	conn     *grpc.ClientConn
	client   proto.MinionServiceClient
}

func Open(ctx context.Context) (*MinionClient, error) {
	_conn, err := grpc.DialContext(
		ctx,
		fmt.Sprintf("%s:%d", config.GRPC.Host, config.GRPC.Port),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	return &MinionClient{
		MinionID: config.Minion.ID,
		conn:     _conn,
		client:   proto.NewMinionServiceClient(_conn),
	}, nil
}

func (c *MinionClient) Close() {
	err := c.conn.Close()
	if err != nil {
		logrus.WithError(err).Fatal("encountered error while closing gRPC client")
	}
}
