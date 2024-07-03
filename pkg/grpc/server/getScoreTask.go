package server

import (
	"context"

	"github.com/scorify/scorify/pkg/grpc/proto"
)

func (s *minionServer_s) GetScoreTask(ctx context.Context, req *proto.GetScoreTaskRequest) (*proto.GetScoreTaskResponse, error) {
	return <-s.ScoreTasks, nil
}
