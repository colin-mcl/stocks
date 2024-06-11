package models

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

// Defines a user type for an individual user
type User struct {
	ID             int
	Username       string
	Email          string
	HashedPassword []byte
	FirstName      string
	LastName       string
	CreatedAt      time.Time
}

// Define a user model type which wraps a db connection pool
type UserModel struct {
	DB *sql.DB
}

// Define a function to insert a user into the table and return its ID
// For now returns nothing
func (m *UserModel) Insert(firstName, lastName, username, email, password string) (int, error) {

	// Generates a bcrypt hash from the plaintext password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return -1, err
	}

	stmt := `INSERT INTO users (first_name, last_name, username, email,
	hashed_password, created_at)
	VALUES (?, ?, ?, ?, ?, NOW())`

	// Attempts to insert the user into the table
	res, err := m.DB.Exec(stmt, firstName, lastName, username, email,
		string(hashedPassword))

	// checks if the new user has an already existing account by comparing to other emails
	if err != nil {
		var mySQLError *mysql.MySQLError
		if errors.As(err, &mySQLError) {
			if mySQLError.Number == 1062 &&
				strings.Contains(mySQLError.Message, "users_uc_email") {
				return -1, fmt.Errorf("Error: duplicate email, %w", err)
			}
		}

		return -1, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return -1, err
	}
	return int(id), nil
}

// Verifies that a user with the provided email and password exists and returns
// the user's ID if so.
func (m *UserModel) Authenticate(email, password string) (int, error) {
	var id int
	var hashedPassword []byte

	stmt := `SELECT id, hashed_password FROM users WHERE email = ?`

	err := m.DB.QueryRow(stmt, email).Scan(&id, &hashedPassword)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return -1, fmt.Errorf("Error: invalid credentials")
		}

		return -1, err
	}

	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return -1, fmt.Errorf("Error: invalid credentials")
		}

		return -1, err
	}

	return id, nil
}

// Checks whether there exists a user with the given ID
func (m *UserModel) Exists(id int) (bool, error) {
	var queriedID int

	stmt := `SELECT id FROM users WHERE id = ?`

	err := m.DB.QueryRow(stmt, id).Scan(&queriedID)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

// Returns a user with the matching ID
func (m *UserModel) Get(id int) (*User, error) {
	stmt := `SELECT id, first_name, last_name, created FROM users
	WHERE id = ?`

	row := m.DB.QueryRow(stmt, id)

	u := &User{}

	err := row.Scan(&u.ID, &u.FirstName, &u.LastName, &u.CreatedAt)
	if err != nil {
		return nil, err
	}

	return u, nil
}
