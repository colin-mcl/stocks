package gapi_test

import (
	"context"
	"testing"

	"github.com/colin-mcl/stocks/pb"
	"github.com/colin-mcl/stocks/util"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/metadata"
)

func TestGetUser(t *testing.T) {

	type cases struct {
		name         string
		input        *pb.GetUserRequest
		inputContext context.Context
		expectedErr  error
	}

	// TODO: come back when login user is implemented
	testServer.CreateUser(context.Background(),
		&pb.CreateUserRequest{
			Username:  "user",
			Email:     "userEmail",
			Password:  "password",
			FirstName: "first",
			LastName:  "last",
		})
	r, err := testServer.LoginUser(context.Background(), &pb.LoginUserRequest{
		Email:    "userEmail",
		Password: "password",
	})
	assert.NoError(t, err)

	for _, scenario := range []cases{
		{
			name:         "no authentication",
			input:        &pb.GetUserRequest{Email: "userEmail"},
			inputContext: context.Background(),
			expectedErr:  errors.Errorf("unauthorized"),
		},
		{
			name:  "bad authentication",
			input: &pb.GetUserRequest{Email: "userEmail"},
			inputContext: metadata.NewIncomingContext(context.Background(),
				metadata.New(map[string](string){
					"authentication": util.RandomString(16)})),
			expectedErr: errors.Errorf("unauthorized"),
		},
		{
			name:  "bad id",
			input: &pb.GetUserRequest{Email: "fake email"},
			inputContext: metadata.NewIncomingContext(context.Background(),
				metadata.New(map[string](string){
					"authentication": r.GetAccessToken()})),
			expectedErr: errors.Errorf("failed to get user with email fake email"),
		},
	} {
		t.Run(scenario.name, func(t *testing.T) {
			resp, err := testServer.GetUser(scenario.inputContext, scenario.input)
			assert.Error(t, err)
			assert.Nil(t, resp)
			assert.ErrorContains(t, err, scenario.expectedErr.Error())
		})
	}

	resp, err := testServer.GetUser(
		metadata.NewIncomingContext(context.Background(), metadata.New(
			map[string](string){"authentication": r.GetAccessToken()},
		)), &pb.GetUserRequest{Email: "userEmail"})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "userEmail", resp.GetUser().GetEmail())
	assert.Equal(t, "user", resp.GetUser().GetUsername())
	assert.Equal(t, "first", resp.GetUser().GetFirstName())
	assert.Equal(t, "last", resp.GetUser().GetLastName())
}
