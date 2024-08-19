package models

import (
	"errors"
)

// Defines a user type for an individual user
type User struct {
	Username       string
	Email          string
	HashedPassword []byte
	FirstName      string
	LastName       string
}

var ErrEmptyField error = errors.New("error: user cannot have empty field")

// creates and returns a new user with the provided fields, returns an error
// if any field is empty is empty
func NewUser(username string, email string, hashedPassword []byte,
	firstName string, lastName string) (*User, error) {

	if username == "" || email == "" || len(hashedPassword) == 0 ||
		firstName == "" || lastName == "" {
		return nil, ErrEmptyField
	}

	return &User{
		Username:       username,
		Email:          email,
		HashedPassword: hashedPassword,
		FirstName:      firstName,
		LastName:       lastName,
	}, nil
}
