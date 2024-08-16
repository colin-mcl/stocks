package models

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestAuthenticateUser(t *testing.T) {
	// Tests a non existent user
	id, err := users.Authenticate("fakeEmail", "badPassword")

	require.Equal(t, id, -1)
	require.Error(t, err)
	require.EqualError(t, err, ErrInvalidCredentials.Error())

	// Test an existing user with the wrong password
	id, err = users.Authenticate("colin.mcl@gmail.com", "blah")

	require.Equal(t, id, -1)
	require.Error(t, err)
	require.EqualError(t, err, ErrInvalidCredentials.Error())

	// Test an existing user with the correct password
	id, err = users.Authenticate("colin.mcl@gmail.com", "Password123!")

	require.Equal(t, id, 1)
	require.Nil(t, err)
}

func TestInsertUser(t *testing.T) {
	// Inserting user with email that already exists
	id, err := users.Insert("bad", "user", "badusername",
		"colin.mcl@gmail.com", "pass")

	require.Equal(t, id, -1)
	require.Error(t, err)

	// Check that non existant user doesn't exist
	exists, _ := users.Exists(0)
	require.False(t, exists)

	// Inserting real user and check it exists
	password := time.Now().String()
	username := "realuser#" + password
	email := username + "@email.com"
	id, err = users.Insert("real", "user", username, email, password)

	require.NotEqual(t, id, -1)
	require.Nil(t, err)

	exists, err = users.Exists(id)
	require.True(t, exists)
	require.Nil(t, err)
}

func TestExists(t *testing.T) {
	exists, err := users.Exists(0)
	require.Nil(t, err)
	require.False(t, exists)

	exists, err = users.Exists(1)
	require.Nil(t, err)
	require.True(t, exists)
}
