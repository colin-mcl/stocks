package models

import (
	"database/sql"
	"time"
)

// Defines a user type for an individual user
type User struct {
	ID        int
	firstName string
	lastName  string
	createdAt time.Time
}

// Define a user model type which wraps a db connection pool
type UserModel struct {
	DB *sql.DB
}

// Define a function to insert a user into the table and return its ID
// For now returns nothing
func (m *UserModel) Insert(firstName string, lastName string) (int, error) {
	// Hard coded SQL statement for inserting a new user into the table
	stmt := `INSERT INTO users (first_name, last_name, created)
	VALUES(?, ?, NOW())`

	result, err := m.DB.Exec(stmt, firstName, lastName)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

// Returns a user with the matching ID
func (m *UserModel) Get(id int) (*User, error) {
	return nil, nil
}
