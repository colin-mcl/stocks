package repo

import (
	"database/sql"
	"testing"

	"github.com/colin-mcl/stocks/internal/models"
	"github.com/colin-mcl/stocks/util"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestCreateUser(t *testing.T) {
	pswd, _ := bcrypt.GenerateFromPassword([]byte("password"), 12)
	u, _ := models.NewUser(
		"user",
		util.RandomString(16),
		pswd,
		"first",
		"last")

	id, err := testRepo.CreateUser(u)

	assert.Nil(t, err)
	assert.Positive(t, id)
}

func TestGetUser(t *testing.T) {
	// non-existent user id
	u, err := testRepo.GetUser(0)
	assert.Error(t, err)
	assert.EqualError(t, err, sql.ErrNoRows.Error())
	assert.Nil(t, u)

	u, err = testRepo.GetUser(11)
	assert.Nil(t, err)
	assert.NotNil(t, u)
	assert.Equal(t, "colin", u.Username)
	assert.Equal(t, "colin.mclaughlin02@gmail.com", u.Email)
	assert.Equal(t, "Colin", u.FirstName)
	assert.Equal(t, "Mclaughlin", u.LastName)
}

func TestGetUserByEmail(t *testing.T) {
	// non-existent email
	u, err := testRepo.GetUserByEmail("fake")
	assert.Nil(t, u)
	assert.Error(t, err)
	assert.EqualError(t, err, sql.ErrNoRows.Error())

	u, err = testRepo.GetUserByEmail("colin.mclaughlin02@gmail.com")
	assert.Nil(t, err)
	assert.NotNil(t, u)
	assert.Equal(t, "colin", u.Username)
	assert.Equal(t, "colin.mclaughlin02@gmail.com", u.Email)
	assert.Equal(t, "Colin", u.FirstName)
	assert.Equal(t, "Mclaughlin", u.LastName)
}

func TestUserExists(t *testing.T) {
	exists, err := testRepo.UserExists(-1)
	assert.False(t, exists)
	assert.NoError(t, err)

	exists, err = testRepo.UserExists(11)
	assert.True(t, exists)
	assert.NoError(t, err)
}
