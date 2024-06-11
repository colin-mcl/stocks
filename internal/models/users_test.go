package models

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAuthenticateUser(t *testing.T) {
	// Tests a non existent user
	id, err := testModels.Authenticate("fakeEmail", "badPassword")

	require.Equal(t, id, -1)
	require.NotNil(t, err)
	require.Equal(t, err, fmt.Errorf("Error: invalid credentials"))

	// Test an existing user with the wrong password
	id, err = testModels.Authenticate("colin.mcl@gmail.com", "blah")

	require.Equal(t, id, -1)
	require.NotNil(t, err)
	require.Equal(t, err, fmt.Errorf("Error: invalid credentials"))

	// Test an existing user with the correct password
	id, err = testModels.Authenticate("colin.mcl@gmail.com", "Password123!")

	require.Equal(t, id, 1)
	require.Nil(t, err)
}
