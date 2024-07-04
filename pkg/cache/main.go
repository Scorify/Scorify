package cache

import (
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/scorify/scorify/pkg/config"
)

func NewRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Redis.Host, config.Redis.Port),
		Password: config.Redis.Password,
		DB:       0,
	})
}
