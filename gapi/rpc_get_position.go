package gapi

import (
	"context"
	"database/sql"
	"errors"

	"github.com/colin-mcl/stocks/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) GetPosition(ctx context.Context,
	r *pb.GetPositionRequest) (*pb.GetPositionResponse, error) {
	server.infoLog.Printf("get position request received: %d", r.GetId())

	_, err := server.authenticateUser(ctx)

	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthorized: %s", err.Error())
	}

	p, err := server.positions.Get(int(r.GetId()))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Errorf(codes.NotFound, "no position with id=%d found", r.GetId())
		} else {
			return nil, status.Errorf(codes.Internal, "%s", err.Error())
		}
	}

	return &pb.GetPositionResponse{Pos: convertPosition(p)}, nil
}
