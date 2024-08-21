package gapi_test

import (
	"context"
	"testing"

	"github.com/colin-mcl/stocks/pb"
	"github.com/colin-mcl/stocks/pkg/v1/usecase"
	"github.com/colin-mcl/stocks/util"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {

	type cases struct {
		name        string
		input       *pb.CreateUserRequest
		expectedErr error
	}

	for _, scenario := range []cases{
		{
			name: "duplicate email",
			input: &pb.CreateUserRequest{
				Username: "user", Email: "colin.mclaughlin02@gmail.com",
				Password: "password", FirstName: "colin",
				LastName: "mclaughlin"},
			expectedErr: usecase.ErrAlreadyExists,
		},
		{
			name: "blank Username",
			input: &pb.CreateUserRequest{
				Username: "", Email: util.RandomString(10),
				Password: "pwd", FirstName: "colin", LastName: "mclaughlin"},
			expectedErr: usecase.ErrEmptyField,
		},
		{
			name: "blank email",
			input: &pb.CreateUserRequest{
				Username: "user", Email: "",
				Password: "pwd", FirstName: "colin", LastName: "mclaughlin"},
			expectedErr: usecase.ErrEmptyField,
		},
		{
			name: "blank FirstName",
			input: &pb.CreateUserRequest{
				Username: "user", Email: util.RandomString(15),
				Password: "pwd", FirstName: "", LastName: "mclaughlin"},
			expectedErr: usecase.ErrEmptyField,
		},
		{
			name: "blank LastName",
			input: &pb.CreateUserRequest{
				Username: "user", Email: util.RandomString(15),
				Password: "password", FirstName: "colin", LastName: ""},
			expectedErr: usecase.ErrEmptyField,
		}} {
		t.Run(scenario.name, func(t *testing.T) {
			resp, err := testServer.CreateUser(context.Background(), scenario.input)

			assert.Error(t, err)
			assert.Nil(t, resp)
			assert.ErrorContains(t, err, scenario.expectedErr.Error())
		})
	}

	for _, scenario := range []cases{
		{
			name: "long everything",
			input: &pb.CreateUserRequest{
				Username: util.RandomString(255), Email: util.RandomString(255),
				Password: util.RandomString(72), FirstName: util.RandomString(255),
				LastName: util.RandomString(255),
			},
			expectedErr: nil,
		},
		{
			name: "short everything",
			input: &pb.CreateUserRequest{
				Username: util.RandomString(1), Email: util.RandomString(1),
				Password: util.RandomString(1), FirstName: util.RandomString(1),
				LastName: util.RandomString(1),
			},
			expectedErr: nil,
		}} {
		t.Run(scenario.name, func(t *testing.T) {
			resp, err := testServer.CreateUser(context.Background(), scenario.input)

			assert.NoError(t, err)
			assert.NotNil(t, resp)
			assert.Positive(t, resp.GetId())
		})
	}
}
