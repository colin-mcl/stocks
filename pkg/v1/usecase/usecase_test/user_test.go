package usecase_test

import (
	"testing"

	"github.com/colin-mcl/stocks/internal/models"
	"github.com/colin-mcl/stocks/pkg/v1/usecase"
	"github.com/colin-mcl/stocks/util"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {

	u := &models.User{
		Username:       util.RandomString(20),
		Email:          "colin.mclaughlin02@gmail.com",
		HashedPassword: []byte("password"),
		FirstName:      util.RandomString(7),
		LastName:       util.RandomString(10),
	}
	// should fail because of email already in use

	id, err := testUC.CreateUser(u)

	assert.Equal(t, -1, id)
	assert.Error(t, err)
	assert.EqualError(t, err, usecase.ErrAlreadyExists.Error())

	u.Email = ""
	id, err = testUC.CreateUser(u)
	assert.Equal(t, -1, id)
	assert.Error(t, err)
	assert.EqualError(t, err, usecase.ErrEmptyField.Error())

	u.Email = util.RandomString(10)
	u.Username = ""
	id, err = testUC.CreateUser(u)
	assert.Equal(t, -1, id)
	assert.Error(t, err)
	assert.EqualError(t, err, usecase.ErrEmptyField.Error())

	u.Username = util.RandomString(10)
	u.HashedPassword = []byte("")
	id, err = testUC.CreateUser(u)
	assert.Equal(t, -1, id)
	assert.Error(t, err)
	assert.EqualError(t, err, usecase.ErrEmptyField.Error())

	u.HashedPassword = []byte("password")
	u.FirstName = ""
	id, err = testUC.CreateUser(u)
	assert.Equal(t, -1, id)
	assert.Error(t, err)
	assert.EqualError(t, err, usecase.ErrEmptyField.Error())

	u.FirstName = util.RandomString(10)
	u.LastName = ""
	id, err = testUC.CreateUser(u)
	assert.Equal(t, -1, id)
	assert.Error(t, err)
	assert.EqualError(t, err, usecase.ErrEmptyField.Error())

	u.LastName = util.RandomString(10)

	id, err = testUC.CreateUser(u)
	assert.NoError(t, err)
	assert.Positive(t, id)

}

func TestGetUser(t *testing.T) {
	user, err := testUC.GetUser(0)
	assert.Nil(t, user)
	assert.Error(t, err)
	assert.EqualError(t, err, usecase.ErrDoesNotExist.Error())

	user, err = testUC.GetUser(11)
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, "colin.mclaughlin02@gmail.com", user.Email)
	assert.Equal(t, "colin", user.Username)
	assert.Equal(t, "Colin", user.FirstName)
	assert.Equal(t, "Mclaughlin", user.LastName)
}

func TestGetUserByEmail(t *testing.T) {
	user, err := testUC.GetUserByEmail("fake email")
	assert.Nil(t, user)
	assert.Error(t, err)
	assert.EqualError(t, err, usecase.ErrDoesNotExist.Error())

	user, err = testUC.GetUserByEmail("colin.mclaughlin02@gmail.com")
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, "colin.mclaughlin02@gmail.com", user.Email)
	assert.Equal(t, "colin", user.Username)
	assert.Equal(t, "Colin", user.FirstName)
	assert.Equal(t, "Mclaughlin", user.LastName)
}

func TestDeleteUser(t *testing.T) {
	err := testUC.DeleteUser(0)
	assert.Error(t, err)
	assert.EqualError(t, usecase.ErrDoesNotExist, err.Error())

	u := &models.User{
		Username:       util.RandomString(20),
		Email:          util.RandomString(20),
		HashedPassword: []byte("password"),
		FirstName:      util.RandomString(7),
		LastName:       util.RandomString(10),
	}

	id, err := testUC.CreateUser(u)
	assert.NoError(t, err)

	err = testUC.DeleteUser(id)
	assert.NoError(t, err)

	err = testUC.DeleteUser(id)
	assert.Error(t, err)
	assert.EqualError(t, err, usecase.ErrDoesNotExist.Error())
}
