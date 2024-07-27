package server

import (
	"context"
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/scorify/scorify/pkg/cache"
	"github.com/scorify/scorify/pkg/grpc/proto"
	"github.com/scorify/scorify/pkg/structs"
)

func (m *minionServer_s) Heartbeat(ctx context.Context, req *proto.HeartbeatRequest) (*proto.HeartbeatResponse, error) {
	minion_id_str := req.GetMinionId()
	metrics_str := req.GetMetrics()

	minion_id, err := uuid.Parse(minion_id_str)
	if err != nil {
		return nil, err
	}

	var metrics structs.MinionMetrics
	err = json.Unmarshal([]byte(metrics_str), &metrics)
	if err != nil {
		return nil, err
	}

	metrics.Timestamp = time.Now()
	metrics.MinionID = minion_id

	err = cache.SetMinionMetrics(ctx, minion_id, m.redisClient, &metrics)
	if err != nil {
		return nil, err
	}

	return &proto.HeartbeatResponse{}, nil
}
