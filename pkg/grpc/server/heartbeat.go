package server

import (
	"context"

	"github.com/scorify/scorify/pkg/grpc/proto"
)

func (*minionServer_s) Heartbeat(ctx context.Context, req *proto.HeartbeatRequest) (*proto.HeartbeatResponse, error) {
	return &proto.HeartbeatResponse{}, nil
}
