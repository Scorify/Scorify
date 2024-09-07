package cache

import (
	"context"
	"encoding/json"

	"github.com/redis/go-redis/v9"
	"github.com/scorify/scorify/pkg/graph/model"
)

const (
	// scoreboardUpdateChannel is the channel to publish scoreboard updates
	scoreboardUpdateChannel = "scoreboard_round_update_channel"
)

func PublishScoreboardUpdate(ctx context.Context, redisClient *redis.Client, scoreboardUpdate *model.Scoreboard) (*redis.IntCmd, error) {
	out, err := json.Marshal(scoreboardUpdate)
	if err != nil {
		return nil, err
	}
	err = SetScoreboard(ctx, redisClient, scoreboardUpdate)
	if err != nil {
		return nil, err
	}

	return redisClient.Publish(ctx, scoreboardUpdateChannel, out), nil
}

func SubscribeScoreboardUpdate(ctx context.Context, redisClient *redis.Client) *redis.PubSub {
	return redisClient.Subscribe(ctx, scoreboardUpdateChannel)
}
