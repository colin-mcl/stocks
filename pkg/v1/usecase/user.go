package usecase

import (
	"database/sql"
	"errors"

	"github.com/colin-mcl/stocks/internal/models"
)

// the user file implements all of the user operations in the usecase interface

var (
	ErrAlreadyExists error = errors.New("error: user already exists")
	ErrDoesNotExist  error = errors.New("error: does not exist")
)

// CreateUser
//
// Creates a new user from the supplied argument if the email does not
// already exist
func (uc *UseCase) CreateUser(u *models.User) (int, error) {
	if _, err := uc.repo.GetUserByEmail(u.Email); !errors.Is(err, sql.ErrNoRows) {
		return -1, ErrAlreadyExists
	}

	id, err := uc.repo.CreateUser(u)

	if err != nil {
		return -1, err
	}

	return id, nil
}

// GetUser
//
// Gets the user by their ID and returns the instance
func (uc *UseCase) GetUser(id int) (*models.User, error) {
	user, err := uc.repo.GetUser(id)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrDoesNotExist
		}

		return nil, err
	}

	return user, nil
}

// DeleteUser
//
// Deletes the user instance if it exists
func (uc *UseCase) DeleteUser(id int) error {
	if exists, _ := uc.repo.UserExists(id); !exists {
		return ErrDoesNotExist
	}

	err := uc.repo.DeleteUser(id)

	return err
}
