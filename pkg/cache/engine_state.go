package cache

import (
	"context"
	"errors"

	"github.com/redis/go-redis/v9"
	"github.com/scorify/scorify/pkg/graph/model"
)

const (
	// EngineStateChannel is the key for the engine state pub/sub in redis
	engineStateChannel = "engine_state_channel"
)

var (
	// ErrInvalidEngineState is the error message for invalid engine state
	ErrInvalidEngineState = errors.New("invalid engine state")
)

func PublishEngineState(ctx context.Context, redisClient *redis.Client, state model.EngineState) (*redis.IntCmd, error) {
	out := redisClient.Publish(ctx, engineStateChannel, string(state))

	return out, out.Err()
}

func SubscribeEngineState(ctx context.Context, redisClient *redis.Client) *redis.PubSub {
	return redisClient.Subscribe(ctx, engineStateChannel)
}
