package gapi

import (
	"context"
	"errors"
	"time"

	"github.com/colin-mcl/stocks/internal/models"
	"github.com/colin-mcl/stocks/pb"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) LoginUser(ctx context.Context, r *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
	// Get the user and check if it exists
	user, err := server.users.Get(r.GetEmail())

	if err != nil {
		if errors.Is(err, models.ErrInvalidCredentials) {
			return nil, status.Errorf(codes.Unauthenticated, "%s", err)
		}
		return nil, status.Errorf(codes.Internal, "failed to login: %s", err)
	}

	err = bcrypt.CompareHashAndPassword(user.HashedPassword, []byte(r.GetPassword()))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return nil, status.Errorf(codes.Unauthenticated, "%s", models.ErrInvalidCredentials)
		}
		return nil, status.Errorf(codes.Internal, "failed to login: %s", err)
	}

	accessToken, accessPayload, err := server.tokenMaker.CreateToken(
		user.Email,
		time.Minute*5,
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create access token: %s", err)
	}

	rsp := &pb.LoginUserResponse{
		User:                 convertUser(user),
		AccessToken:          accessToken,
		AccessTokenExpiresAt: timestamppb.New(accessPayload.ExpiredAt),
	}

	return rsp, nil
}
