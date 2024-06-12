package gapi

import (
	"context"
	"errors"

	"github.com/colin-mcl/stocks/internal/models"
	"github.com/colin-mcl/stocks/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateUser(ctx context.Context, r *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	server.infoLog.Printf("create user request received: %s\n", r.GetEmail())

	id, err := server.users.Insert(
		r.GetFirstName(),
		r.GetLastName(),
		r.GetUsername(),
		r.GetEmail(),
		r.GetPassword(),
	)

	if err != nil {
		if errors.Is(err, models.ErrAlreadyExists) {
			return nil, status.Errorf(codes.AlreadyExists, "user already exists: %s", err)
		}
		return nil, status.Errorf(codes.Internal, "failed to create new user %s", err)
	}

	return &pb.CreateUserResponse{Id: int32(id)}, nil
}
