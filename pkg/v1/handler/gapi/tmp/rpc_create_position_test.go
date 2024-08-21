package gapi

import (
	"context"
	"testing"

	"github.com/colin-mcl/stocks/pb"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestCreatePosition(t *testing.T) {
	server := makeDefaultServer()
	req := &pb.CreatePositionRequest{
		Symbol:        "AAPL",
		HeldBy:        11,
		PurchasedAt:   timestamppb.Now(),
		PurchasePrice: 300.1,
		Qty:           11.1,
	}

	// Fails because there is no authentication header
	resp, err := server.CreatePosition(context.Background(), req)

	assert.Error(t, err)
	assert.Nil(t, resp)

	// Bad password and practices, only for testing purposes
	r, err := server.LoginUser(context.Background(), &pb.LoginUserRequest{
		Email:    "colin.mclaughlin02@gmail.com",
		Password: "password",
	})

	assert.NoError(t, err)
	assert.NotNil(t, r)

	md := metadata.New(map[string](string){"authentication": r.GetAccessToken()})
	ctx := metadata.NewIncomingContext(context.Background(), md)
	resp, err = server.CreatePosition(ctx, req)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotZero(t, resp.GetId())
}
