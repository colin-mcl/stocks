package client_test

import (
	"testing"

	"github.com/colin-mcl/stocks/util"
	"github.com/stretchr/testify/assert"
)

func TestGetUser(t *testing.T) {
	u, err := testClient.GetUser("userEmail", util.RandomString(10))
	assert.Error(t, err)
	assert.Nil(t, u)
	assert.ErrorContains(t, err, "unauthorized")

	tkn, err := testClient.LoginUser("userEmail", "password")
	assert.NoError(t, err)

	u, err = testClient.GetUser("userEmail", tkn)
	assert.NoError(t, err)
	assert.NotNil(t, u)
	assert.Equal(t, "user", u.Username)
	assert.Equal(t, "userEmail", u.Email)
	assert.Equal(t, "first", u.FirstName)
	assert.Equal(t, "last", u.LastName)
}
