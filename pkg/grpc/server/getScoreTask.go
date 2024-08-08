package server

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/scorify/scorify/pkg/config"
	"github.com/scorify/scorify/pkg/ent/minion"
	"github.com/scorify/scorify/pkg/grpc/proto"
)

var ErrMinionDeactivated = fmt.Errorf("minion is deactivated")

func (s *minionServer_s) GetScoreTask(ctx context.Context, req *proto.GetScoreTaskRequest) (*proto.GetScoreTaskResponse, error) {
	//TODO: Optomize (redis) and figure out better solution than sleep
	uuid, err := uuid.Parse(req.GetMinionId())
	if err != nil {
		time.Sleep(config.Interval)
		return nil, err
	}

	exist, err := s.entClient.Minion.Query().
		Where(
			minion.ID(uuid),
			minion.Deactivated(false),
		).Exist(ctx)
	if err != nil {
		time.Sleep(config.Interval)
		return nil, err
	}

	if !exist {
		time.Sleep(config.Interval)
		return nil, ErrMinionDeactivated
	}

	return <-s.ScoreTasks, nil
}
