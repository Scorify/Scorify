package cache

import (
	"context"
	"encoding/json"

	"github.com/redis/go-redis/v9"
	"github.com/scorify/scorify/pkg/graph/model"
)

const (
	// kothScoreboardUpdateChannel is the channel to publish scoreboard updates
	kothScoreboardUpdateChannel = "koth_scoreboard_round_update_channel"
)

func PublishKothScoreboardUpdate(ctx context.Context, redisClient *redis.Client, scoreboardUpdate *model.KothScoreboard) (*redis.IntCmd, error) {
	out, err := json.Marshal(scoreboardUpdate)
	if err != nil {
		return nil, err
	}
	err = SetLatestKothScoreboard(ctx, redisClient, scoreboardUpdate)
	if err != nil {
		return nil, err
	}

	return redisClient.Publish(ctx, kothScoreboardUpdateChannel, out), nil
}

func SubscribeKothScoreboardUpdate(ctx context.Context, redisClient *redis.Client) *redis.PubSub {
	return redisClient.Subscribe(ctx, kothScoreboardUpdateChannel)
}
