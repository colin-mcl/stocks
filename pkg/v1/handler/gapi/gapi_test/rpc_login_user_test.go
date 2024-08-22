package gapi_test

import (
	"context"
	"testing"
	"time"

	"github.com/colin-mcl/stocks/pb"
	"github.com/colin-mcl/stocks/util"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestLoginUser(t *testing.T) {
	type cases struct {
		name        string
		input       *pb.LoginUserRequest
		expectedErr error
	}

	for _, scenario := range []cases{
		{
			name: "non-existent email",
			input: &pb.LoginUserRequest{Email: "fake email",
				Password: "password"},
			expectedErr: errors.Errorf("does not exist"),
		},
		{
			name: "bad password",
			input: &pb.LoginUserRequest{Email: "userEmail",
				Password: "password123"},
			expectedErr: errors.Errorf("invalid credentials"),
		},
	} {
		t.Run(scenario.name, func(t *testing.T) {
			resp, err := testServer.LoginUser(context.Background(), scenario.input)
			assert.Error(t, err)
			assert.Nil(t, resp)
			assert.ErrorContains(t, err, scenario.expectedErr.Error())
		})
	}

	// Create new user with random password
	pswd := util.RandomString(16)
	email := util.RandomString(16)
	testServer.CreateUser(context.Background(),
		&pb.CreateUserRequest{
			Username:  email,
			Email:     email,
			Password:  pswd,
			FirstName: "first",
			LastName:  "last",
		})

	resp, err := testServer.LoginUser(context.Background(), &pb.LoginUserRequest{
		Email:    email,
		Password: pswd})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotNil(t, resp.GetAccessToken())
	assert.WithinDuration(t, resp.GetAccessTokenExpiresAt().AsTime(), time.Now(),
		5*time.Minute)
	assert.Equal(t, email, resp.GetUser().GetUsername())
	assert.Equal(t, email, resp.GetUser().GetEmail())
	assert.Equal(t, "first", resp.GetUser().GetFirstName())
	assert.Equal(t, "last", resp.GetUser().GetLastName())

}
