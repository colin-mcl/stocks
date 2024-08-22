package gapi

import (
	"context"
	"errors"
	"time"

	"github.com/colin-mcl/stocks/pb"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) LoginUser(ctx context.Context, r *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
	server.infoLog.Printf("login request received: %s", r.GetEmail())

	// Get the user and check if it exists
	user, err := server.uc.GetUserByEmail(r.GetEmail())

	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to login: %s", err)
	}

	err = bcrypt.CompareHashAndPassword(user.HashedPassword, []byte(r.GetPassword()))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return nil, status.Errorf(codes.Unauthenticated, "invalid credentials")
		}
		return nil, status.Errorf(codes.Internal, "failed to login: %s", err)
	}

	accessToken, accessPayload, err := server.tokenMaker.CreateToken(user.Email,
		time.Minute*5)

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
