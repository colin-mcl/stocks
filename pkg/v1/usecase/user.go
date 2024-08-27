package usecase

import (
	"database/sql"
	"errors"

	"github.com/colin-mcl/stocks/internal/models"
)

// the user file implements all of the user operations in the usecase interface

// CreateUser
//
// Creates a new user from the supplied argument if the email does not
// already exist
func (uc *UseCase) CreateUser(u *models.User) (int, error) {
	// checks that user does not already exist
	if _, err := uc.repo.GetUserByEmail(u.Email); !errors.Is(err, sql.ErrNoRows) {
		return -1, ErrAlreadyExists
	}

	// checks that all fields are non-empty
	// TODO: MOVE this to GRPC handler and maybe create validation package
	if u.Email == "" || u.Username == "" || len(u.HashedPassword) == 0 ||
		u.FirstName == "" || u.LastName == "" {
		return -1, ErrEmptyField
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

// GetUserByEmail
//
// Gets the user by their email and returns the instance
func (uc *UseCase) GetUserByEmail(email string) (*models.User, error) {
	user, err := uc.repo.GetUserByEmail(email)

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
