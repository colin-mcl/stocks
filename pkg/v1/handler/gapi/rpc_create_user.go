package gapi

import (
	"context"

	"github.com/colin-mcl/stocks/internal/models"
	"github.com/colin-mcl/stocks/pb"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateUser(ctx context.Context, r *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	server.infoLog.Printf("create user request received: %s\n", r.GetEmail())

	pswd, err := bcrypt.GenerateFromPassword([]byte(r.GetPassword()), 12)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "hashing failed: %s", err)
	}

	id, err := server.uc.CreateUser(&models.User{
		Username:       r.GetUsername(),
		Email:          r.GetEmail(),
		HashedPassword: pswd,
		FirstName:      r.GetFirstName(),
		LastName:       r.GetLastName()})

	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create new user %s", err)
	}

	return &pb.CreateUserResponse{Id: int32(id)}, nil
}
