package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"github.com/scorify/scorify/pkg/static"
	"github.com/scorify/scorify/pkg/structs"
	"github.com/sirupsen/logrus"
)

const (
	// minion_metrics_key_prefix is the prefix for minion metrics
	minion_metrics_key_prefix = "minion_metrics"

	// minion_metrics_key_prefix is the prefix for minion metrics
	minion_liveness_key_prefix = "minion_liveness"
)

func GetMinionMetrics(ctx context.Context, minionID uuid.UUID, redisClient *redis.Client) (minionMetrics *structs.MinionMetrics, err error) {
	data, err := redisClient.Get(ctx, fmt.Sprintf("%s:%s", minion_metrics_key_prefix, minionID.String())).Result()
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(data), minionMetrics)
	if err != nil {
		return nil, err
	}

	return minionMetrics, nil
}

func SetMinionMetrics(ctx context.Context, minionID uuid.UUID, redisClient *redis.Client, metrics *structs.MinionMetrics) error {
	data, err := json.Marshal(metrics)
	if err != nil {
		return err
	}

	return redisClient.Set(ctx, fmt.Sprintf("%s:%s", minion_metrics_key_prefix, minionID.String()), data, 0).Err()
}

func GetMinionMetricsGroup(ctx context.Context, redisClient *redis.Client) ([]*structs.MinionMetrics, error) {
	keys, err := redisClient.Keys(ctx, fmt.Sprintf("%s:*", minion_metrics_key_prefix)).Result()
	if err != nil {
		return nil, err
	}

	data, err := redisClient.MGet(ctx, keys...).Result()
	if err != nil {
		return nil, err
	}

	return static.MapSlice(data, func(_ int, value interface{}) *structs.MinionMetrics {
		data, ok := value.(string)
		if !ok {
			logrus.WithField("value", value).Error("failed to convert value to string")
			return nil
		}

		var minionMetrics structs.MinionMetrics
		err := json.Unmarshal([]byte(data), &minionMetrics)
		if err != nil {
			logrus.WithError(err).Error("failed to unmarshal data")
			return nil
		}

		return &minionMetrics
	}), nil
}

func GetMinionLiveness(ctx context.Context, minionID uuid.UUID, redisClient *redis.Client) (minionMetrics *structs.MinionMetrics, err error) {
	data, err := redisClient.Get(ctx, fmt.Sprintf("%s:%s", minion_liveness_key_prefix, minionID.String())).Result()
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(data), minionMetrics)
	if err != nil {
		return nil, err
	}

	return minionMetrics, nil
}

func SetMinionLiveness(ctx context.Context, minionID uuid.UUID, redisClient *redis.Client, minionMetrics *structs.MinionMetrics) error {
	data, err := json.Marshal(minionMetrics)
	if err != nil {
		return err
	}

	return redisClient.Set(ctx, fmt.Sprintf("%s:%s", minion_liveness_key_prefix, minionID.String()), data, time.Minute).Err()
}

func GetMinionLivenessGroup(ctx context.Context, redisClient *redis.Client) ([]*structs.MinionMetrics, error) {
	keys, err := redisClient.Keys(ctx, fmt.Sprintf("%s:*", minion_liveness_key_prefix)).Result()
	if err != nil {
		return nil, err
	}

	data, err := redisClient.MGet(ctx, keys...).Result()
	if err != nil {
		return nil, err
	}

	return static.MapSlice(data, func(_ int, value interface{}) *structs.MinionMetrics {
		data, ok := value.(string)
		if !ok {
			logrus.WithField("value", value).Error("failed to convert value to string")
			return nil
		}

		var minionMetrics structs.MinionMetrics
		err := json.Unmarshal([]byte(data), &minionMetrics)
		if err != nil {
			logrus.WithError(err).Error("failed to unmarshal data")
			return nil
		}

		return &minionMetrics
	}), nil
}
