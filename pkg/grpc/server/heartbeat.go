package server

import (
	"context"
	"encoding/json"
	"net"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"github.com/scorify/scorify/pkg/cache"
	"github.com/scorify/scorify/pkg/grpc/proto"
	"github.com/scorify/scorify/pkg/structs"
	"google.golang.org/grpc/peer"
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

	p, ok := peer.FromContext(ctx)
	if ok {
		clientIP, _, err := net.SplitHostPort(p.Addr.String())
		if err != nil {
			return nil, err
		}

		metrics.IP = clientIP
	}

	metrics.Timestamp = time.Now()
	metrics.MinionID = minion_id

	err = setHeartbeatData(ctx, minion_id, m.redisClient, &metrics)
	if err != nil {
		return nil, err
	}

	return &proto.HeartbeatResponse{}, nil
}

func setHeartbeatData(ctx context.Context, minion_id uuid.UUID, redisClient *redis.Client, metrics *structs.MinionMetrics) error {
	err := cache.SetMinionLiveness(ctx, minion_id, redisClient, metrics)
	if err != nil {
		return err
	}

	return cache.SetMinionMetrics(ctx, minion_id, redisClient, metrics)
}
