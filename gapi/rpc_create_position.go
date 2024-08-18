package gapi

import (
	"context"

	"github.com/colin-mcl/stocks/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreatePosition(ctx context.Context,
	r *pb.CreatePositionRequest) (*pb.CreatePositionResponse, error) {

	server.infoLog.Printf("create position request received: %s", r.GetSymbol())

	_, err := server.authenticateUser(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthorized: %s", err.Error())
	}

	id, err := server.positions.Insert(
		r.GetSymbol(),
		int(r.GetHeldBy()),
		r.GetPurchasedAt().AsTime(),
		r.GetPurchasePrice(),
		r.GetQty(),
	)

	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &pb.CreatePositionResponse{Id: int32(id)}, nil
}
