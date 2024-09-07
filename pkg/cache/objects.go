package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"github.com/scorify/scorify/pkg/ent"
	"github.com/scorify/scorify/pkg/graph/model"
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

func getRoundObjectKey(roundID uuid.UUID) ObjectKey {
	return ObjectKey("object-round-" + roundID.String())
}

func getUserObjectKey(userID uuid.UUID) ObjectKey {
	return ObjectKey("object-user-" + userID.String())
}

func setObject(ctx context.Context, redisClient *redis.Client, key ObjectKey, obj interface{}, expiration time.Duration) error {
	out, err := json.Marshal(obj)
	if err != nil {
		return err
	}

	return redisClient.Set(ctx, string(key), out, expiration).Err()
}

func getObject(ctx context.Context, redisClient *redis.Client, key ObjectKey, obj interface{}) bool {
	out, err := redisClient.Get(ctx, string(key)).Result()
	if err != nil {
		return false
	}

	return json.Unmarshal([]byte(out), obj) == nil
}

func GetUser(ctx context.Context, redisClient *redis.Client, entClient *ent.Client, userID uuid.UUID) (*ent.User, error) {
	var entUser *ent.User
	if getObject(ctx, redisClient, getUserObjectKey(userID), entUser) {
		return entUser, nil
	}

	entUser, err := entClient.User.Get(ctx, userID)
	if err != nil {
		return nil, err
	}

	err = setObject(ctx, redisClient, getUserObjectKey(userID), entUser, medium)
	if err != nil {
		return nil, err
	}

	return entUser, nil
}

func SetUser(ctx context.Context, redisClient *redis.Client, entUser *ent.User) error {
	return setObject(ctx, redisClient, getUserObjectKey(entUser.ID), entUser, medium)
}

func GetRound(ctx context.Context, redisClient *redis.Client, entClient *ent.Client, roundID uuid.UUID) (*ent.Round, error) {
	var entRound *ent.Round
	if getObject(ctx, redisClient, getRoundObjectKey(roundID), entRound) {
		return entRound, nil
	}

	entRound, err := entClient.Round.Get(ctx, roundID)
	if err != nil {
		return nil, err
	}

	if entRound.Complete {
		err = setObject(ctx, redisClient, getRoundObjectKey(roundID), entRound, long)
	} else {
		err = setObject(ctx, redisClient, getRoundObjectKey(roundID), entRound, short)
	}
	if err != nil {
		return nil, err
	}

	return entRound, nil
}

func SetRound(ctx context.Context, redisClient *redis.Client, entRound *ent.Round) error {
	return setObject(ctx, redisClient, getRoundObjectKey(entRound.ID), entRound, long)
}

func GetLatestRound(ctx context.Context, redisClient *redis.Client) (*ent.Round, bool) {
	var entRound *ent.Round
	ok := getObject(ctx, redisClient, LatestRoundObjectKey, entRound)

	return entRound, ok
}

func SetLatestRound(ctx context.Context, redisClient *redis.Client, entRound *ent.Round) error {
	return setObject(ctx, redisClient, LatestRoundObjectKey, entRound, long)
}

func GetScoreboard(ctx context.Context, redisClient *redis.Client) (*model.Scoreboard, bool) {
	var entScoreboard *model.Scoreboard
	ok := getObject(ctx, redisClient, ScoreboardObjectKey, entScoreboard)

	return entScoreboard, ok
}

func SetScoreboard(ctx context.Context, redisClient *redis.Client, scoreboard *model.Scoreboard) error {
	return setObject(ctx, redisClient, ScoreboardObjectKey, scoreboard, long)
}
