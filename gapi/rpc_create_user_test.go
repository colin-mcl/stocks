package gapi

import (
	"context"
	"testing"

	"github.com/colin-mcl/stocks/pb"
	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	server := makeDefaultServer()

	resp, err := server.CreateUser(context.Background(),
		&pb.CreateUserRequest{
			FirstName: "testing",
			LastName:  "functionality"})

	require.NotNil(t, resp)
	require.Nil(t, err)
	require.NotZero(t, resp.GetId())
	require.Positive(t, resp.GetId())

}
