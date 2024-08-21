package models

// Defines a user type for an individual user
type User struct {
	Username       string
	Email          string
	HashedPassword []byte
	FirstName      string
	LastName       string
}
