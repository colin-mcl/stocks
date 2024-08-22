package gapi

import (
	"github.com/colin-mcl/stocks/internal/models"
	"github.com/colin-mcl/stocks/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Provides a simple function to convert user model to user protobuf struct
func convertUser(user *models.User) *pb.User {
	return &pb.User{
		Username:  user.Username,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}
}

func convertPosition(p *models.Position) *pb.Position {
	return &pb.Position{
		Id:            int32(p.ID),
		Symbol:        p.Symbol,
		HeldBy:        int32(p.HeldBy),
		PurchasedAt:   timestamppb.New(p.PurchasedAt),
		PurchasePrice: p.PurchasePrice,
		Qty:           p.Qty,
	}
}
