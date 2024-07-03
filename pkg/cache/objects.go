package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type ObjectKey string

const (
	ScoreboardObjectKey  ObjectKey = "object-scoreboard"
	LatestRoundObjectKey ObjectKey = "object-latest-round"
)

func GetScoreboardObjectKey(round int) ObjectKey {
	return ObjectKey(fmt.Sprintf("object-scoreboard-%d", round))
}

func GetCheckObjectKey(checkID uuid.UUID) ObjectKey {
	return ObjectKey("object-check-" + checkID.String())
}

func GetCheckConfigObjectKey(checkConfigID uuid.UUID) ObjectKey {
	return ObjectKey("object-check-config-" + checkConfigID.String())
}

func GetRoundObjectKey(roundID uuid.UUID) ObjectKey {
	return ObjectKey("object-round-" + roundID.String())
}

func GetScoreCacheObjectKey(roundID uuid.UUID) ObjectKey {
	return ObjectKey("object-score-cache-" + roundID.String())
}

func GetStatusObjectKey(statusID uuid.UUID) ObjectKey {
	return ObjectKey("object-status-" + statusID.String())
}

func GetUserObjectKey(userID uuid.UUID) ObjectKey {
	return ObjectKey("object-user-" + userID.String())
}

func SetObject(ctx context.Context, redisClient *redis.Client, key ObjectKey, obj interface{}, expiration time.Duration) error {
	out, err := json.Marshal(obj)
	if err != nil {
		return err
	}

	return redisClient.Set(ctx, string(key), out, expiration).Err()
}

func GetObject(ctx context.Context, redisClient *redis.Client, key ObjectKey, obj interface{}) bool {
	out, err := redisClient.Get(ctx, string(key)).Result()
	if err != nil {
		return false
	}

	return json.Unmarshal([]byte(out), obj) == nil
}
