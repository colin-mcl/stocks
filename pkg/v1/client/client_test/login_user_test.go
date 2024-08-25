package client_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoginUser(t *testing.T) {
	tkn, err := testClient.LoginUser("userEmail", "badPass")
	assert.Error(t, err)
	assert.Equal(t, "", tkn)

	tkn, err = testClient.LoginUser("userEmail", "password")
	assert.NoError(t, err)
	assert.NotEqual(t, "", tkn)
}
