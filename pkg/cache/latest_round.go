package cache

import (
	"context"
	"encoding/json"

	"github.com/redis/go-redis/v9"
	"github.com/scorify/scorify/pkg/ent"
)

const (
	// LatestRoundChannel is the key for the latest round pub/sub in redis
	latestRoundChannel = "latest_round_channel"
)

func PublishLatestRound(ctx context.Context, redisClient *redis.Client, entRound *ent.Round) (*redis.IntCmd, error) {
	err := SetObject(ctx, redisClient, LatestRoundObjectKey, entRound, 0)
	if err != nil {
		return nil, err
	}

	out, err := json.Marshal(entRound)
	if err != nil {
		return nil, err
	}

	return redisClient.Publish(ctx, latestRoundChannel, out), nil
}

func SubscribeLatestRound(ctx context.Context, redisClient *redis.Client) *redis.PubSub {
	return redisClient.Subscribe(ctx, latestRoundChannel)
}
