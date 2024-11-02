package engine

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"github.com/scorify/scorify/pkg/cache"
	"github.com/scorify/scorify/pkg/config"
	"github.com/scorify/scorify/pkg/ent"
	"github.com/scorify/scorify/pkg/ent/checkconfig"
	"github.com/scorify/scorify/pkg/ent/minion"
	"github.com/scorify/scorify/pkg/ent/round"
	"github.com/scorify/scorify/pkg/ent/status"
	"github.com/scorify/scorify/pkg/graph/model"
	"github.com/scorify/scorify/pkg/helpers"
	"github.com/scorify/scorify/pkg/static"
	"github.com/scorify/scorify/pkg/structs"
	"github.com/sirupsen/logrus"
)

type state int

const (
	EnginePaused state = iota
	EngineWaiting
	EngineRunning
)

type Client struct {
	lock             *sync.Mutex
	state            state
	ctx              context.Context
	ent              *ent.Client
	redis            *redis.Client
	taskRequestChan  chan<- *structs.TaskRequest
	taskResponseChan <-chan *structs.TaskResponse
	workerStatusChan chan<- *structs.WorkerStatus
}

func NewEngine(
	ctx context.Context,
	entClient *ent.Client,
	redis *redis.Client,
	taskRequestChan chan<- *structs.TaskRequest,
	taskResponseChan <-chan *structs.TaskResponse,
	workerStatusChan chan<- *structs.WorkerStatus,
) *Client {
	return &Client{
		lock:             &sync.Mutex{},
		state:            EnginePaused,
		ctx:              ctx,
		ent:              entClient,
		redis:            redis,
		taskRequestChan:  taskRequestChan,
		taskResponseChan: taskResponseChan,
		workerStatusChan: workerStatusChan,
	}
}

func (e *Client) Stop() error {
	if e.state == EngineRunning {
		_, err := cache.PublishEngineState(context.Background(), e.redis, model.EngineStateStopping)
		if err != nil {
			return err
		}
	}

	e.lock.Lock()
	defer e.lock.Unlock()

	if e.state == EngineWaiting {
		e.state = EnginePaused
		_, err := cache.PublishEngineState(context.Background(), e.redis, model.EngineStatePaused)
		return err
	}

	return fmt.Errorf("cannot stop engine from state %q", e.state)
}

func (e *Client) Start() error {
	if e.state == EngineRunning {
		return fmt.Errorf("engine already running")
	}

	go e.loop()

	e.state = EngineWaiting
	_, err := cache.PublishEngineState(context.Background(), e.redis, model.EngineStateWaiting)
	return err
}

func (e *Client) State() (model.EngineState, error) {
	switch e.state {
	case EnginePaused:
		return model.EngineStatePaused, nil
	case EngineWaiting:
		return model.EngineStateWaiting, nil
	case EngineRunning:
		return model.EngineStateRunning, nil
	}

	return "", fmt.Errorf("unknown engine state %q", e.state)
}

func (e *Client) loop() {
	ticker := time.NewTicker(config.Interval)

	defer func() {
		ticker.Stop()
		e.state = EnginePaused
		cache.PublishEngineState(context.Background(), e.redis, model.EngineStatePaused)
	}()

	for {
		select {
		case <-e.ctx.Done():
			return
		case <-ticker.C:
			if e.state == EnginePaused {
				return
			}

			err := e.loopRoundRunner()
			if err != nil {
				logrus.WithError(err).Error("failed to run round")
				return
			}

			e.state = EngineWaiting
			_, err = cache.PublishEngineState(context.Background(), e.redis, model.EngineStateWaiting)
			if err != nil {
				logrus.WithError(err).Error("failed to publish engine state")
				return
			}
		}
	}
}

func (e *Client) loopRoundRunner() error {
	e.lock.Lock()
	defer e.lock.Unlock()

	e.state = EngineRunning
	_, err := cache.PublishEngineState(context.Background(), e.redis, model.EngineStateRunning)
	if err != nil {
		return err
	}

	roundCtx, cancel := context.WithTimeout(e.ctx, config.Interval-time.Millisecond*500)
	defer cancel()

	// Get the current round number
	var roundNumber int
	entLastRound, err := e.ent.Round.Query().
		Order(
			ent.Desc(round.FieldNumber),
		).
		First(e.ctx)
	if err != nil {
		roundNumber = 1
	} else {
		roundNumber = entLastRound.Number + 1
	}

	// Create new round
	entRound, err := e.ent.Round.Create().
		SetNumber(roundNumber).
		Save(e.ctx)
	if err != nil {
		logrus.WithError(err).Error("failed to create new round")
		return nil
	}
	err = cache.SetRound(roundCtx, e.redis, entRound)
	if err != nil {
		logrus.WithError(err).Error("failed to set round")
		return err
	}

	_, err = cache.PublishLatestRound(roundCtx, e.redis, entRound)
	if err != nil {
		logrus.WithError(err).Error("failed to set latest round object")
		return err
	}

	entDisabledMinions, err := e.ent.Minion.Query().
		Where(
			minion.Deactivated(true),
		).
		All(e.ctx)
	if err != nil {
		logrus.WithError(err).Error("failed to get disabled minions")
		return err
	}

	disabledMinions := structs.WorkerStatus(
		static.MapSlice(
			entDisabledMinions,
			func(i int, m *ent.Minion) uuid.UUID {
				return m.ID
			},
		),
	)

	select {
	case <-roundCtx.Done():
		logrus.WithError(roundCtx.Err()).Error("failed to send disabled minions to workers")
		return nil
	case <-time.After(time.Millisecond * 500):
		logrus.Error("failed to send disabled minions to workers")
	case e.workerStatusChan <- &disabledMinions:
	}

	logrus.WithField("time", time.Now()).Infof("Running round %d", roundNumber)

	// Run the round
	err = e.runRound(roundCtx, entRound)

	logrus.WithField("time", time.Now()).Infof("Round %d complete", roundNumber)

	return err
}

func (e *Client) runRound(ctx context.Context, entRound *ent.Round) error {
	// Get all the tasks
	tasks, err := e.ent.CheckConfig.Query().
		WithUser().
		WithCheck().
		Order(
			// ID are uuids and thus check orders are shuffled
			ent.Desc(checkconfig.FieldID),
		).
		All(ctx)
	if err != nil {
		return err
	}

	entStatuses, err := e.ent.Status.CreateBulk(
		static.MapSlice(
			tasks,
			func(i int, task *ent.CheckConfig) *ent.StatusCreate {
				return e.ent.Status.Create().
					SetRound(entRound).
					SetUser(task.Edges.User).
					SetCheck(task.Edges.Check).
					SetPoints(task.Edges.Check.Weight).
					SetStatus(status.StatusUnknown)
			},
		)...,
	).Save(ctx)
	if err != nil {
		return err
	}

	// Create a map of round tasks to keep track of the tasks
	roundTasks := structs.NewSyncMap[uuid.UUID, *ent.CheckConfig]()

	for i, entStatus := range entStatuses {
		roundTasks.Set(entStatus.ID, tasks[i])
	}

	wg := &sync.WaitGroup{}

	wg.Add(roundTasks.Length())

	// Submit tasks to the workers
	go func() {
		for _, entStatus := range entStatuses {
			entConfig, ok := roundTasks.Get(entStatus.ID)
			if !ok {
				logrus.WithField("id", entStatus.ID).Error("failed to get task")
				continue
			}

			conf, err := json.Marshal(entConfig.Config)
			if err != nil {
				logrus.WithError(err).Error("failed to marshal check config")
				continue
			}

			e.taskRequestChan <- &structs.TaskRequest{
				StatusID:   entStatus.ID,
				SourceName: entConfig.Edges.Check.Source,
				Config:     string(conf),
			}
		}
	}()

	allChecksReported := make(chan struct{})
	checksRemain := true

	defer func() {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()

		for status_id := range roundTasks.Map() {
			entStatus, err := e.ent.Status.UpdateOneID(status_id).
				SetStatus(status.StatusUnknown).
				SetPoints(0).
				Save(ctx)
			if err != nil {
				logrus.WithField("id", status_id).WithError(err).Error("failed to update status")
			} else {
				logrus.WithField("status", entStatus).Debug("status not reported, set to 0")
			}
		}

		_, err = entRound.Update().
			SetComplete(true).
			Save(ctx)
		if err != nil {
			logrus.WithError(err).Error("failed to set round as complete")
		}

		type userScore struct {
			UserID uuid.UUID `json:"user_id"`
			Sum    int       `json:"sum"`
		}
		var users []userScore

		err = e.ent.Status.Query().
			Where(
				status.HasRoundWith(round.ID(entRound.ID)),
			).
			GroupBy(status.FieldUserID).
			Aggregate(ent.Sum(status.FieldPoints)).
			Scan(ctx, &users)
		if err != nil {
			logrus.WithError(err).Error("failed to aggregate points")
			return
		}

		_, err = e.ent.ScoreCache.CreateBulk(
			static.MapSlice(
				users,
				func(i int, user userScore) *ent.ScoreCacheCreate {
					return e.ent.ScoreCache.Create().
						SetRound(entRound).
						SetUserID(user.UserID).
						SetPoints(user.Sum)
				},
			)...,
		).Save(ctx)
		if err != nil {
			logrus.WithError(err).Error("failed to create score cache")
			return
		}

		scoreboard, err := helpers.Scoreboard(ctx, e.ent)
		if err != nil {
			logrus.WithError(err).Error("failed to get scoreboard")
			return
		}

		_, err = cache.PublishScoreboardUpdate(ctx, e.redis, scoreboard)
		if err != nil {
			logrus.WithError(err).Error("failed to publish scoreboard")
		}
	}()

	// Wait for the results
	for checksRemain {
		select {
		case <-allChecksReported:
			checksRemain = false
		case <-ctx.Done():
			return nil
		case result := <-e.taskResponseChan:
			if result.Status == status.StatusUp || result.Status == status.StatusDown || result.Status == status.StatusUnknown {
				go e.updateStatus(ctx, roundTasks, result.StatusID, result.Error, result.Status, result.MinionID, allChecksReported, wg)
			} else {
				go e.updateStatus(ctx, roundTasks, result.StatusID, result.Error, status.StatusUnknown, uuid.UUID{}, allChecksReported, wg)
				logrus.WithFields(logrus.Fields{
					"status":    result.Status,
					"status_id": result.StatusID,
				}).Error("unknown status")
			}
		}
	}

	wgDone := make(chan struct{})
	go func() {
		defer close(wgDone)
		wg.Wait()
	}()

	select {
	case <-wgDone:
	case <-ctx.Done():
	}

	return nil
}

func cleanStatus(s string) string {
	return strings.ReplaceAll(s, "\x00", "")
}

func (e *Client) updateStatus(ctx context.Context, roundTasks *structs.SyncMap[uuid.UUID, *ent.CheckConfig], status_id uuid.UUID, errorMessage string, _status status.Status, minionID uuid.UUID, allChecksReported chan<- struct{}, wg *sync.WaitGroup) {
	_, ok := roundTasks.Get(status_id)
	if !ok {
		logrus.WithField("status_id", status_id).Error("uuid not belong to round was submitted")
		return
	}

	roundTasks.Delete(status_id)

	defer wg.Done()

	entStatusUpdate := e.ent.Status.UpdateOneID(status_id).
		SetStatus(status.Status(_status)).
		SetMinionID(minionID)

	if errorMessage != "" {
		entStatusUpdate.SetError(cleanStatus(errorMessage))
	}

	if _status != status.StatusUp {
		entStatusUpdate.SetPoints(0)
	}

	_, err := entStatusUpdate.Save(ctx)
	if err != nil {
		logrus.WithField("id", status_id).WithError(err).Error("failed to update status")
		return
	}

	if roundTasks.Length() == 0 {
		allChecksReported <- struct{}{}
	}
}
