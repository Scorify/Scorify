package server

import (
	"context"
	"fmt"
	"net"

	"github.com/google/uuid"
	"github.com/scorify/scorify/pkg/ent"
	"github.com/scorify/scorify/pkg/ent/minion"
	"github.com/scorify/scorify/pkg/grpc/proto"
	"google.golang.org/grpc/peer"
)

func (m *minionServer_s) Enroll(ctx context.Context, req *proto.EnrollRequest) (*proto.EnrollResponse, error) {
	minion_id := req.GetMinionId()
	minion_name := req.GetName()

	uuid, err := uuid.Parse(minion_id)
	if err != nil {
		return nil, err
	}

	p, ok := peer.FromContext(ctx)
	if !ok {
		return nil, fmt.Errorf("failed to get peer from context")
	}

	client_ip, _, err := net.SplitHostPort(p.Addr.String())
	if err != nil {
		return nil, err
	}

	exists, err := m.entClient.Minion.Query().
		Where(
			minion.IDEQ(uuid),
		).Exist(ctx)
	if err != nil {
		return nil, err
	}

	var entMinion *ent.Minion
	if exists {
		entMinion, err = m.entClient.Minion.Get(ctx, uuid)
	} else {
		entMinion, err = m.entClient.Minion.Create().
			SetID(uuid).
			SetIP(client_ip).
			SetName(minion_name).
			Save(ctx)
	}
	if err != nil {
		return nil, err
	}

	if exists && entMinion.IP != client_ip {
		_, err = entMinion.Update().
			SetIP(client_ip).
			Save(ctx)
		if err != nil {
			return nil, err
		}
	}

	return &proto.EnrollResponse{}, nil
}
