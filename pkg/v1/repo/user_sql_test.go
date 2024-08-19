package repo

import (
	"testing"

	"github.com/colin-mcl/stocks/internal/models"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestCreateUser(t *testing.T) {
	pswd, _ := bcrypt.GenerateFromPassword([]byte("password"), 12)
	u, _ := models.NewUser(
		"user",
		"user@email.com",
		pswd,
		"first",
		"last")

	id, err := testRepo.CreateUser(u)

	assert.Nil(t, err)
	assert.Positive(t, id)
}
