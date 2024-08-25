package client_test

import (
	"testing"

	"github.com/colin-mcl/stocks/util"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	id, err := testClient.CreateUser(
		"first", "last", "username", "userEmail", "password")

	assert.Equal(t, -1, id)
	assert.Error(t, err)
	assert.ErrorContains(t, err, "already exists")

	id, err = testClient.CreateUser(
		"first", "last", "username", util.RandomString(16), "password")

	assert.NoError(t, err)
	assert.NotEqual(t, -1, id)
	assert.Positive(t, id)
}
