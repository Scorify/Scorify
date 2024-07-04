package cache

import (
	"context"
	"encoding/json"

	"github.com/redis/go-redis/v9"
	"github.com/scorify/scorify/pkg/graph/model"
)

const (
	// GlobalNotificationChannel is the key for the global notifications pub/sub in redis
	globalNotificationChannel = "global_notification_channel"
)

func PublishNotification(ctx context.Context, redisClient *redis.Client, message string, notification_type model.NotificationType) (*redis.IntCmd, error) {
	out, err := json.Marshal(model.Notification{
		Message: message,
		Type:    notification_type,
	})
	if err != nil {
		return nil, err
	}

	return redisClient.Publish(ctx, globalNotificationChannel, out), nil
}

func SubscribeNotification(ctx context.Context, redisClient *redis.Client) *redis.PubSub {
	return redisClient.Subscribe(ctx, globalNotificationChannel)
}
