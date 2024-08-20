package usecase

import (
	"testing"

	"github.com/colin-mcl/stocks/internal/models"
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
	assert.EqualError(t, err, ErrAlreadyExists.Error())

	u.Email = util.RandomString(22)

	id, err = testUC.CreateUser(u)
	assert.NoError(t, err)
	assert.Positive(t, id)

}

func TestGetUser(t *testing.T) {
	user, err := testUC.GetUser(0)
	assert.Nil(t, user)
	assert.Error(t, err)
	assert.EqualError(t, err, ErrDoesNotExist.Error())

	user, err = testUC.GetUser(11)
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
	assert.EqualError(t, ErrDoesNotExist, err.Error())

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
	assert.EqualError(t, err, ErrDoesNotExist.Error())
}
