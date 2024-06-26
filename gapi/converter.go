package gapi

import (
	"github.com/colin-mcl/stocks/internal/models"
	"github.com/colin-mcl/stocks/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Provides a simple function to convert user model to user protobuf struct
func convertUser(user *models.User) *pb.User {
	return &pb.User{
		Id:        int32(user.ID),
		Username:  user.Username,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		CreatedAt: timestamppb.New(user.CreatedAt),
	}
}
