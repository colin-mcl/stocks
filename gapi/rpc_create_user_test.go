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

	u, err := server.GetUser(context.Background(), &pb.GetUserRequest{Id: resp.GetId()})

	require.NotNil(t, u)
	require.Nil(t, err)
	require.Equal(t, u.GetUser().GetId(), resp.GetId())
	require.Equal(t, u.GetUser().GetFirstName(), "testing")
	require.Equal(t, u.GetUser().GetLastName(), "functionality")

}
