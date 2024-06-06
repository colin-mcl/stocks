package gapi

import (
	"context"
	"testing"

	"github.com/colin-mcl/stocks/pb"
	"github.com/stretchr/testify/require"
)

func TestGetUser(t *testing.T) {
	server := makeDefaultServer()

	resp, err := server.GetUser(context.Background(), &pb.GetUserRequest{Id: 0})

	require.Nil(t, resp)
	require.Error(t, err)

	resp, err = server.GetUser(context.Background(), &pb.GetUserRequest{Id: 1})

	require.NotNil(t, resp)
	require.Nil(t, err)
	require.Equal(t, resp.User.GetFirstName(), "Colin")
	require.Equal(t, resp.User.GetLastName(), "Mclaughlin")
	require.NotNil(t, resp.User.GetCreatedAt())
}
