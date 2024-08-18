package gapi

import (
	"context"
	"testing"

	"github.com/colin-mcl/stocks/pb"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/metadata"
)

func TestGetPosition(t *testing.T) {
	server := makeDefaultServer()

	req := &pb.GetPositionRequest{Id: 1}
	badReq := &pb.GetPositionRequest{Id: -1}

	// fails because no authentication
	resp, err := server.GetPosition(context.Background(), req)
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

	resp, err = server.GetPosition(ctx, badReq)
	assert.EqualError(t, err, "rpc error: code = NotFound desc = no position with id=-1 found")
	assert.Nil(t, resp)

	resp, err = server.GetPosition(ctx, req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "TSLA", resp.GetPos().GetSymbol())
	assert.Equal(t, int32(11), resp.GetPos().GetHeldBy())
	assert.Equal(t, 210.1, resp.GetPos().GetPurchasePrice())
	assert.Equal(t, 2.5, resp.GetPos().GetQty())
	assert.Equal(t, int32(1), resp.GetPos().GetId())
}
