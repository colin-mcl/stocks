package gapi

import (
	"context"

	"github.com/colin-mcl/stocks/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) GetUser(
	ctx context.Context, r *pb.GetUserRequest) (*pb.GetUserResponse, error) {

	server.infoLog.Printf("get user request received: %d", r.GetId())

	_, err := server.authenticateUser(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthorized: %s", err.Error())
	}

	u, err := server.uc.GetUser(int(r.GetId()))
	if err != nil {
		return nil, status.Errorf(codes.Internal,
			"failed to get user with id %d: %s\n", r.GetId(), err)
	}

	return &pb.GetUserResponse{User: convertUser(u)}, nil
}
