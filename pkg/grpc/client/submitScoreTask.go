package client

import (
	"context"

	"github.com/google/uuid"
	"github.com/scorify/scorify/pkg/ent/status"
	"github.com/scorify/scorify/pkg/grpc/proto"
)

func (c *MinionClient) SubmitScoreTask(ctx context.Context, statusID uuid.UUID, err string, checkStatus status.Status) (*proto.SubmitScoreTaskResponse, error) {
	return c.client.SubmitScoreTask(ctx, &proto.SubmitScoreTaskRequest{
		MinionId: c.MinionID.String(),
		StatusId: statusID.String(),
		Error:    err,
		Status: func() proto.Status {
			switch checkStatus {
			case status.StatusUp:
				return proto.Status_up
			case status.StatusDown:
				return proto.Status_down
			default:
				return proto.Status_unknown
			}
		}(),
	})
}
