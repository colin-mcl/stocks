package gapi

import (
	"context"

	"github.com/colin-mcl/stocks/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateUser(ctx context.Context, r *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	server.infoLog.Printf("create user request received: %s %s\n",
		r.GetFirstName(), r.GetLastName(),
	)

	id, err := server.users.Insert(r.GetFirstName(), r.GetLastName())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create new user %s", err)
	}

	return &pb.CreateUserResponse{Id: int32(id)}, nil
}
