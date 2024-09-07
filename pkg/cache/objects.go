package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"github.com/scorify/scorify/pkg/ent"
)

const (
	short   time.Duration = 10 * time.Second
	medium  time.Duration = 1 * time.Minute
	long    time.Duration = 5 * time.Minute
	forever time.Duration = 0
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

func GetUser(ctx context.Context, redisClient *redis.Client, entClient *ent.Client, userID uuid.UUID) (*ent.User, error) {
	var entUser *ent.User
	if GetObject(ctx, redisClient, GetUserObjectKey(userID), entUser) {
		return entUser, nil
	}

	entUser, err := entClient.User.Get(ctx, userID)
	if err != nil {
		return nil, err
	}

	err = SetObject(ctx, redisClient, GetUserObjectKey(userID), entUser, medium)
	if err != nil {
		return nil, err
	}

	return entUser, nil
}

func GetRound(ctx context.Context, redisClient *redis.Client, entClient *ent.Client, roundID uuid.UUID) (*ent.Round, error) {
	var entRound *ent.Round
	if GetObject(ctx, redisClient, GetRoundObjectKey(roundID), entRound) {
		return entRound, nil
	}

	entRound, err := entClient.Round.Get(ctx, roundID)
	if err != nil {
		return nil, err
	}

	if entRound.Complete {
		err = SetObject(ctx, redisClient, GetRoundObjectKey(roundID), entRound, long)
	} else {
		err = SetObject(ctx, redisClient, GetRoundObjectKey(roundID), entRound, short)
	}
	if err != nil {
		return nil, err
	}

	return entRound, nil
}

func GetLatestRound(ctx context.Context, redisClient *redis.Client) (*ent.Round, bool) {
	var entRound *ent.Round
	ok := GetObject(ctx, redisClient, LatestRoundObjectKey, entRound)

	return entRound, ok
}
