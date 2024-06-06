package models

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetUser(t *testing.T) {
	u, err := testModels.Get(-1)

	require.Nil(t, u)
	require.Error(t, err)
	require.Equal(t, err, sql.ErrNoRows)

	u, err = testModels.Get(1)

	require.NotNil(t, u)
	require.Nil(t, err)
	require.Equal(t, u.FirstName, "Colin")
	require.Equal(t, u.LastName, "Mclaughlin")
	require.Equal(t, u.ID, 1)
}

func TestInsertUser(t *testing.T) {
	id, err := testModels.Insert("test", "user")

	require.NotZero(t, id)
	require.Positive(t, id)
	require.Nil(t, err)

	u, err := testModels.Get(id)
	require.NotNil(t, u)
	require.Nil(t, err)
	require.Equal(t, u.FirstName, "test")
	require.Equal(t, u.LastName, "user")
	require.Equal(t, u.ID, id)
}
