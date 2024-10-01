package helpers

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/scorify/scorify/pkg/ent"
	"github.com/scorify/scorify/pkg/ent/check"
	"github.com/scorify/scorify/pkg/ent/round"
	"github.com/scorify/scorify/pkg/ent/scorecache"
	"github.com/scorify/scorify/pkg/ent/status"
	"github.com/scorify/scorify/pkg/ent/user"
	"github.com/scorify/scorify/pkg/graph/model"
)

func Scoreboard(ctx context.Context, entClient *ent.Client) (*model.Scoreboard, error) {
	latestRound, err := entClient.Round.Query().
		Order(
			ent.Desc(round.FieldNumber),
		).
		First(ctx)
	if err != nil {
		rounds, countErr := entClient.Round.Query().Count(ctx)
		if countErr != nil {
			return nil, countErr
		}

		if rounds == 0 {
			return EmptyScoreboard(ctx, entClient)
		}

		return nil, err
	}

	return ScoreboardByRound(ctx, entClient, latestRound.Number)
}

func EmptyScoreboard(ctx context.Context, entClient *ent.Client) (*model.Scoreboard, error) {
	scoreboard := &model.Scoreboard{}

	entUsers, err := entClient.User.Query().
		Where(
			user.RoleEQ(user.RoleUser),
		).
		Order(
			ent.Asc(user.FieldNumber),
		).
		All(ctx)
	if err != nil {
		return nil, err
	}

	scoreboard.Teams = entUsers

	entChecks, err := entClient.Check.Query().
		Order(
			ent.Asc(check.FieldName),
		).
		All(ctx)
	if err != nil {
		return nil, err
	}

	scoreboard.Checks = entChecks

	scoreboard.Round = &ent.Round{
		Number: 0,
	}

	scoreboard.Statuses = make([][]*ent.Status, len(entChecks))
	for i := range scoreboard.Statuses {
		scoreboard.Statuses[i] = make([]*ent.Status, len(entUsers))
	}

	scoreboard.Scores = make([]*model.Score, len(entUsers))
	for i, entUser := range entUsers {
		scoreboard.Scores[i] = &model.Score{
			User:  entUser,
			Score: 0,
		}
	}

	return scoreboard, nil
}

func ScoreboardByRound(ctx context.Context, entClient *ent.Client, roundNumber int) (*model.Scoreboard, error) {
	scoreboard := &model.Scoreboard{}

	latestRound, err := entClient.Round.Query().
		Order(
			ent.Desc(round.FieldNumber),
		).
		First(ctx)
	if err != nil {
		return nil, err
	}

	if roundNumber > latestRound.Number {
		return nil, fmt.Errorf("round %d has not started yet", roundNumber)
	}

	entUsers, err := entClient.User.Query().
		Where(
			user.RoleEQ(user.RoleUser),
		).
		Order(
			ent.Asc(user.FieldNumber),
		).
		All(ctx)
	if err != nil {
		return nil, err
	}

	scoreboard.Teams = entUsers

	entChecks, err := entClient.Check.Query().
		Order(
			ent.Asc(check.FieldName),
		).
		All(ctx)
	if err != nil {
		return nil, err
	}

	scoreboard.Checks = entChecks

	entRound, err := entClient.Round.Query().
		Where(
			round.NumberEQ(roundNumber),
		).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	scoreboard.Round = entRound

	entStatuses, err := entClient.Status.Query().
		Where(
			status.HasRoundWith(
				round.IDEQ(entRound.ID),
			),
		).All(ctx)
	if err != nil {
		return nil, err
	}

	lookup := make(map[uuid.UUID]int, len(entUsers)+len(entChecks))
	for i, entUser := range entUsers {
		lookup[entUser.ID] = i
	}
	for i, entCheck := range entChecks {
		lookup[entCheck.ID] = i
	}

	scoreboard.Statuses = make([][]*ent.Status, len(entChecks))
	for i := range scoreboard.Statuses {
		scoreboard.Statuses[i] = make([]*ent.Status, len(entUsers))
	}

	for _, entStatus := range entStatuses {
		check_index, ok := lookup[entStatus.CheckID]
		if !ok {
			continue
		}

		user_index, ok := lookup[entStatus.UserID]
		if !ok {
			continue
		}

		scoreboard.Statuses[check_index][user_index] = entStatus
	}

	var TeamScore []struct {
		TeamID uuid.UUID `json:"user_id"`
		Sum    int       `json:"sum"`
	}

	err = entClient.ScoreCache.Query().
		Where(
			scorecache.HasRoundWith(
				round.NumberLTE(roundNumber),
			),
		).
		GroupBy(
			scorecache.FieldUserID,
		).
		Aggregate(
			ent.Sum(
				scorecache.FieldPoints,
			),
		).
		Scan(ctx, &TeamScore)
	if err != nil {
		return nil, err
	}

	scoreboard.Scores = make([]*model.Score, len(entUsers))
	for _, entUser := range entUsers {
		user_index, ok := lookup[entUser.ID]
		if !ok {
			continue
		}

		scoreboard.Scores[user_index] = &model.Score{
			User:  entUser,
			Score: 0,
		}
	}

	for _, teamScore := range TeamScore {
		user_index, ok := lookup[teamScore.TeamID]
		if !ok {
			continue
		}

		scoreboard.Scores[user_index].Score = teamScore.Sum
	}

	return scoreboard, nil
}
