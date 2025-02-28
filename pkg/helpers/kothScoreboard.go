package helpers

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/scorify/scorify/pkg/ent"
	"github.com/scorify/scorify/pkg/ent/kothcheck"
	"github.com/scorify/scorify/pkg/ent/kothstatus"
	"github.com/scorify/scorify/pkg/ent/round"
	"github.com/scorify/scorify/pkg/ent/user"
	"github.com/scorify/scorify/pkg/graph/model"
)

func KothScoreboard(ctx context.Context, entClient *ent.Client) (*model.KothScoreboard, error) {
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
			return EmptyKothScoreboard(ctx, entClient)
		}

		return nil, err
	}

	return KothScoreboardByRound(ctx, entClient, latestRound.Number)
}

func KothScoreboardByRound(ctx context.Context, entClient *ent.Client, roundNumber int) (*model.KothScoreboard, error) {
	kothScoreboard := &model.KothScoreboard{}

	// Check if the round has started
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

	// Get the teams
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

	// Get the round
	entRound, err := entClient.Round.Query().
		Where(
			round.NumberEQ(roundNumber),
		).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	kothScoreboard.Round = entRound

	// Get koth checks
	entKothChecks, err := entClient.KothCheck.Query().
		WithStatuses(
			// Get koth status for check if it has been scored in this round
			func(q *ent.KothStatusQuery) {
				q.Where(
					kothstatus.RoundIDEQ(entRound.ID),
				).WithUser()
			},
		).
		Order(
			ent.Asc(kothcheck.FieldName),
		).
		All(ctx)
	if err != nil {
		return nil, err
	}

	lookup := make(map[uuid.UUID]int, len(entUsers)+len(entKothChecks))
	for i, entUser := range entUsers {
		lookup[entUser.ID] = i
	}
	for i, entKothCheck := range entKothChecks {
		lookup[entKothCheck.ID] = i
	}

	// Get the team scores
	var teamScores []struct {
		TeamID uuid.UUID `json:"user_id"`
		Sum    int       `json:"sum"`
	}

	// Get team scores
	err = entClient.KothStatus.Query().
		Where(
			kothstatus.HasRoundWith(
				round.NumberLTE(roundNumber),
			),
		).
		GroupBy(
			kothstatus.FieldUserID,
		).
		Aggregate(
			ent.Sum(
				kothstatus.FieldPoints,
			),
		).
		Scan(ctx, &teamScores)
	if err != nil {
		return nil, err
	}

	kothScoreboard.Scores = make([]*model.Score, len(entUsers))
	for _, entUser := range entUsers {
		userIndex, ok := lookup[entUser.ID]
		if !ok {
			continue
		}

		kothScoreboard.Scores[userIndex] = &model.Score{
			User:  entUser,
			Score: 0,
		}
	}

	for _, teamScore := range teamScores {
		userIndex, ok := lookup[teamScore.TeamID]
		if !ok {
			continue
		}

		kothScoreboard.Scores[userIndex].Score = teamScore.Sum
	}

	// Get pwnd koth checks
	entKothPwndChecks, err := entClient.KothCheck.Query().
		WithStatuses(
			func(ksq *ent.KothStatusQuery) {
				ksq.WithRound().
					WithUser().
					Order(
						ent.Desc(kothstatus.FieldCreateTime),
					).
					Limit(1)
			},
		).
		Where(
			kothcheck.HasStatusesWith(
				// check has been claimed in this round
				kothstatus.HasUser(),
				kothstatus.HasRoundWith(
					round.NumberLTE(roundNumber),
				),
			),
		).
		All(ctx)

	if err != nil {
		return nil, err
	}

	// Create a slice of booleans to track if a check has been pwnd
	isPwnd := make([]bool, len(entKothChecks))
	for _, entKothCheck := range entKothPwndChecks {
		kothCheckIndex, ok := lookup[entKothCheck.ID]
		if !ok {
			continue
		}

		isPwnd[kothCheckIndex] = true
	}

	kothScoreboard.Checks = make([]*model.KothCheckScore, len(entKothChecks))
	for i, entKothCheck := range entKothChecks {
		var (
			host        *string   = nil
			user        *ent.User = nil
			statusError *string   = nil
		)

		if isPwnd[i] {
			host = &entKothCheck.Host
		}

		if len(entKothCheck.Edges.Statuses) > 0 {
			if entKothCheck.Edges.Statuses[0].Edges.User != nil {
				user = entKothCheck.Edges.Statuses[0].Edges.User
			}

			if entKothCheck.Edges.Statuses[0].Error != "" {
				statusError = &entKothCheck.Edges.Statuses[0].Error
			}
		}

		kothScoreboard.Checks[i] = &model.KothCheckScore{
			ID:         entKothCheck.ID,
			Name:       entKothCheck.Name,
			Host:       host,
			User:       user,
			Error:      statusError,
			CreateTime: entKothCheck.CreateTime,
			UpdateTime: entKothCheck.UpdateTime,
		}

	}

	return kothScoreboard, nil
}

func EmptyKothScoreboard(ctx context.Context, entClient *ent.Client) (*model.KothScoreboard, error) {
	kothScoreboard := &model.KothScoreboard{}

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

	entChecks, err := entClient.KothCheck.Query().
		Order(
			ent.Asc(kothcheck.FieldName),
		).
		All(ctx)
	if err != nil {
		return nil, err
	}

	kothScoreboard.Checks = make([]*model.KothCheckScore, len(entChecks))
	for i, entKothCheck := range entChecks {
		kothScoreboard.Checks[i] = &model.KothCheckScore{
			ID:         entKothCheck.ID,
			Name:       entKothCheck.Name,
			User:       nil,
			Host:       nil,
			Error:      nil,
			CreateTime: time.Now(),
			UpdateTime: time.Now(),
		}
	}

	kothScoreboard.Round = &ent.Round{
		Number: 0,
	}

	kothScoreboard.Scores = make([]*model.Score, len(entUsers))
	for i, entUser := range entUsers {
		kothScoreboard.Scores[i] = &model.Score{
			User:  entUser,
			Score: 0,
		}
	}

	return kothScoreboard, nil
}
