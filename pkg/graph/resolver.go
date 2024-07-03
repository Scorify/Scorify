package graph

import (
	"github.com/redis/go-redis/v9"
	"github.com/scorify/scorify/pkg/engine"
	"github.com/scorify/scorify/pkg/ent"
	"github.com/scorify/scorify/pkg/grpc/proto"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

//go:generate go run github.com/99designs/gqlgen generate

type Resolver struct {
	Ent                  *ent.Client
	Redis                *redis.Client
	Engine               *engine.Client
	ScoreTaskChan        <-chan *proto.GetScoreTaskResponse
	ScoreTaskReponseChan chan<- *proto.SubmitScoreTaskRequest
}
