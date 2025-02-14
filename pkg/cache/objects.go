package cache

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"github.com/scorify/scorify/pkg/ent"
	"github.com/scorify/scorify/pkg/ent/round"
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
	LatestKothScoreboardObjectKey ObjectKey = "object-latest-koth-scoreboard"
	LatestScoreboardObjectKey     ObjectKey = "object-latest-scoreboard"
	LatestRoundObjectKey          ObjectKey = "object-latest-round"
)

func getRoundObjectKey(roundID uuid.UUID) ObjectKey {
	return ObjectKey("object-round-" + roundID.String())
}

func getUserObjectKey(userID uuid.UUID) ObjectKey {
	return ObjectKey("object-user-" + userID.String())
}

func getCheckObjectKey(checkID uuid.UUID) ObjectKey {
	return ObjectKey("object-check-" + checkID.String())
}

func getKothCheckObjectKey(checkID uuid.UUID) ObjectKey {
	return ObjectKey("object-kothcheck-" + checkID.String())
}

func getScoreboardObjectKey(round int) ObjectKey {
	return ObjectKey(fmt.Sprintf("object-scoreboard-%d", round))
}

func getKothScoreboardObjectKey(round int) ObjectKey {
	return ObjectKey(fmt.Sprintf("object-koth-scoreboard-%d", round))
}

func getMinionObjectKey(minionID uuid.UUID) ObjectKey {
	return ObjectKey("object-minion-" + minionID.String())
}

func getJWTObjectKey(token_hash string) ObjectKey {
	return ObjectKey("object-jwt-" + token_hash)
}

func deleteObject(ctx context.Context, redisClient *redis.Client, key ObjectKey) error {
	return redisClient.Del(ctx, string(key)).Err()
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
	entUser := &ent.User{}
	if getObject(ctx, redisClient, getUserObjectKey(userID), entUser) {
		return entUser, nil
	}

	entUser, err := entClient.User.Get(ctx, userID)
	if err != nil {
		return nil, err
	}

	return entUser, setObject(ctx, redisClient, getUserObjectKey(userID), entUser, medium)
}

func SetUser(ctx context.Context, redisClient *redis.Client, entUser *ent.User) error {
	return setObject(ctx, redisClient, getUserObjectKey(entUser.ID), entUser, medium)
}

func GetRound(ctx context.Context, redisClient *redis.Client, entClient *ent.Client, roundID uuid.UUID) (*ent.Round, error) {
	entRound := &ent.Round{}
	if getObject(ctx, redisClient, getRoundObjectKey(roundID), entRound) {
		return entRound, nil
	}

	entRound, err := entClient.Round.Get(ctx, roundID)
	if err != nil {
		return nil, err
	}

	if entRound.Complete {
		return entRound, setObject(ctx, redisClient, getRoundObjectKey(roundID), entRound, long)
	}
	return entRound, setObject(ctx, redisClient, getRoundObjectKey(roundID), entRound, short)
}

func SetRound(ctx context.Context, redisClient *redis.Client, entRound *ent.Round) error {
	return setObject(ctx, redisClient, getRoundObjectKey(entRound.ID), entRound, long)
}

func GetLatestRound(ctx context.Context, redisClient *redis.Client, entClient *ent.Client) (*ent.Round, error) {
	entRound := &ent.Round{}
	if getObject(ctx, redisClient, LatestRoundObjectKey, entRound) {
		return entRound, nil
	}

	entRound, err := entClient.Round.Query().
		Where(
			round.Complete(true),
		).
		Order(
			ent.Desc(
				round.FieldCreateTime,
			),
		).
		First(ctx)
	if err != nil {
		return nil, err
	}

	return entRound, setObject(ctx, redisClient, LatestRoundObjectKey, entRound, long)
}

func SetLatestRound(ctx context.Context, redisClient *redis.Client, entRound *ent.Round) error {
	return setObject(ctx, redisClient, LatestRoundObjectKey, entRound, forever)
}

func GetLatestScoreboard(ctx context.Context, redisClient *redis.Client) (*model.Scoreboard, bool) {
	scoreboard := &model.Scoreboard{}
	return scoreboard, getObject(ctx, redisClient, LatestScoreboardObjectKey, scoreboard)
}

func SetLatestScoreboard(ctx context.Context, redisClient *redis.Client, scoreboard *model.Scoreboard) error {
	return setObject(ctx, redisClient, LatestScoreboardObjectKey, scoreboard, forever)
}

func GetScoreboard(ctx context.Context, redisClient *redis.Client, round int) (*model.Scoreboard, bool) {
	scoreboard := &model.Scoreboard{}
	return scoreboard, getObject(ctx, redisClient, getScoreboardObjectKey(round), scoreboard)
}

func SetScoreboard(ctx context.Context, redisClient *redis.Client, scoreboard *model.Scoreboard) error {
	return setObject(ctx, redisClient, getScoreboardObjectKey(scoreboard.Round.Number), scoreboard, forever)
}

func GetKothScoreboard(ctx context.Context, redisClient *redis.Client, round int) (*model.KothScoreboard, bool) {
	kothScoreboard := &model.KothScoreboard{}
	return kothScoreboard, getObject(ctx, redisClient, getKothScoreboardObjectKey(round), kothScoreboard)
}

func SetKothScoreboard(ctx context.Context, redisClient *redis.Client, kothScoreboard *model.KothScoreboard) error {
	return setObject(ctx, redisClient, getKothScoreboardObjectKey(kothScoreboard.Round.Number), kothScoreboard, forever)
}

func SetLatestKothScoreboard(ctx context.Context, redisClient *redis.Client, kothScoreboard *model.KothScoreboard) error {
	return setObject(ctx, redisClient, LatestKothScoreboardObjectKey, kothScoreboard, forever)
}

func GetLatestKothScoreboard(ctx context.Context, redisClient *redis.Client) (*model.KothScoreboard, bool) {
	kothScoreboard := &model.KothScoreboard{}
	return kothScoreboard, getObject(ctx, redisClient, LatestKothScoreboardObjectKey, kothScoreboard)
}

func GetCheck(ctx context.Context, redisClient *redis.Client, entClient *ent.Client, checkID uuid.UUID) (*ent.Check, error) {
	entCheck := &ent.Check{}
	if getObject(ctx, redisClient, getCheckObjectKey(checkID), entCheck) {
		return entCheck, nil
	}

	entCheck, err := entClient.Check.Get(ctx, checkID)
	if err != nil {
		return nil, err
	}

	return entCheck, setObject(ctx, redisClient, getCheckObjectKey(checkID), entCheck, medium)
}

func SetCheck(ctx context.Context, redisClient *redis.Client, entCheck *ent.Check) error {
	return setObject(ctx, redisClient, getCheckObjectKey(entCheck.ID), entCheck, medium)
}

func SetAuth(ctx context.Context, redisClient *redis.Client, token string, expiration int) error {
	token_digest := sha256.Sum256([]byte(token))
	token_hash := fmt.Sprintf("%x", token_digest)

	return setObject(ctx, redisClient, getJWTObjectKey(token_hash), true, time.Until(time.Unix(int64(expiration), 0)))
}

func GetAuth(ctx context.Context, redisClient *redis.Client, token string) bool {
	token_digest := sha256.Sum256([]byte(token))
	token_hash := fmt.Sprintf("%x", token_digest)

	unused := false
	return getObject(ctx, redisClient, getJWTObjectKey(token_hash), &unused)
}

func DeleteAuth(ctx context.Context, redisClient *redis.Client, token string) error {
	token_digest := sha256.Sum256([]byte(token))
	token_hash := fmt.Sprintf("%x", token_digest)

	return deleteObject(ctx, redisClient, getJWTObjectKey(token_hash))
}

func GetKothCheck(ctx context.Context, redisClient *redis.Client, entClient *ent.Client, checkID uuid.UUID) (*ent.KothCheck, error) {
	entKothCheck := &ent.KothCheck{}
	if getObject(ctx, redisClient, getKothCheckObjectKey(checkID), entKothCheck) {
		return entKothCheck, nil
	}

	entKothCheck, err := entClient.KothCheck.Get(ctx, checkID)
	if err != nil {
		return nil, err
	}

	return entKothCheck, setObject(ctx, redisClient, getCheckObjectKey(checkID), entKothCheck, medium)
}

func SetKothCheck(ctx context.Context, redisClient *redis.Client, entKothCheck *ent.KothCheck) error {
	return setObject(ctx, redisClient, getKothCheckObjectKey(entKothCheck.ID), entKothCheck, medium)
}

func GetMinion(ctx context.Context, redisClient *redis.Client, entClient *ent.Client, minionID uuid.UUID) (*ent.Minion, error) {
	entMinion := &ent.Minion{}
	if getObject(ctx, redisClient, getMinionObjectKey(minionID), entMinion) {
		return entMinion, nil
	}

	entMinion, err := entClient.Minion.Get(ctx, minionID)
	if err != nil {
		return nil, err
	}

	return entMinion, setObject(ctx, redisClient, getMinionObjectKey(minionID), entMinion, medium)
}

func SetMinion(ctx context.Context, redisClient *redis.Client, entMinion *ent.Minion) error {
	return setObject(ctx, redisClient, getMinionObjectKey(entMinion.ID), entMinion, medium)
}
