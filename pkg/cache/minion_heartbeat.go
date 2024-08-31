package cache

import (
	"context"
	"encoding/json"

	"github.com/redis/go-redis/v9"
	"github.com/scorify/scorify/pkg/structs"
)

const (
	// MinionStatusChannel is the key for the minion metrics pub/sub in redis
	MinionStatusChannel = "minion_heartbeat_channel"
)

func PublishMinionHeartbeat(ctx context.Context, redisClient *redis.Client, metrics *structs.Heartbeat) (*redis.IntCmd, error) {
	out, err := json.Marshal(metrics)
	if err != nil {
		return nil, err
	}

	resp := redisClient.Publish(ctx, MinionStatusChannel, string(out))

	return resp, resp.Err()
}

func SubscribeMiniontMetrics(ctx context.Context, redisClient *redis.Client) *redis.PubSub {
	return redisClient.Subscribe(ctx, MinionStatusChannel)
}
