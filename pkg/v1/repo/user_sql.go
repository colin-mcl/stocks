package repo

import (
	"database/sql"
	"errors"

	"github.com/colin-mcl/stocks/internal/models"
)

// user_sql implements the user methods on the repo

// CreateUser
//
// Creates a new user which was supplied as the argument and returns the ID
// if successfully inserted into the db
func (repo *Repo) CreateUser(u *models.User) (int, error) {

	stmt := `INSERT INTO users (first_name, last_name, username, email,
	hashed_password, created_at)
	VALUES (?, ?, ?, ?, ?, NOW())`

	// Attempts to insert the user into the table
	res, err := repo.db.Exec(stmt, u.FirstName, u.LastName, u.Username, u.Email,
		string(u.HashedPassword))

	if err != nil {
		return -1, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return -1, err
	}
	return int(id), nil
}

// GetUser
//
// Gets the user with the provided id if it exists in the db
func (repo *Repo) GetUser(id int) (*models.User, error) {
	stmt := `SELECT username, email, hashed_password, first_name,
	last_name FROM users
	WHERE id = ?`

	row := repo.db.QueryRow(stmt, id)

	u := &models.User{}

	err := row.Scan(
		&u.Username,
		&u.Email,
		&u.HashedPassword,
		&u.FirstName,
		&u.LastName)

	if err != nil {
		return nil, err
	}

	return u, nil
}

// GetUserByEmail
//
// Gets the user with the provided email if it exists in the db
func (repo *Repo) GetUserByEmail(email string) (*models.User, error) {
	stmt := `SELECT username, email, hashed_password, first_name,
	last_name FROM users
	WHERE email = ?`

	row := repo.db.QueryRow(stmt, email)

	u := &models.User{}

	err := row.Scan(
		&u.Username,
		&u.Email,
		&u.HashedPassword,
		&u.FirstName,
		&u.LastName)

	if err != nil {
		return nil, err
	}

	return u, nil
}

// UserExists
//
// Checks whether there exists a user with the given ID in the db
func (repo *Repo) UserExists(id int) (bool, error) {
	var queriedID int

	stmt := `SELECT id FROM users WHERE id = ?`

	err := repo.db.QueryRow(stmt, id).Scan(&queriedID)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

// DeleteUser
//
// Attempts to delete a user by their id
func (repo *Repo) DeleteUser(id int) error {
	stmt := `DELETE FROM users WHERE id = ?`

	_, err := repo.db.Exec(stmt, id)

	return err
}

// TODO: is this used anywhere? which layer should it go on?
// Authenticate
//
// Verifies that a user with the provided email and password exists and returns
// the ID if so
// func (repo *Repo) Authenticate(email, password string) (int, error) {
// 	var id int
// 	var hashedPassword []byte

// 	stmt := `SELECT id, hashed_password FROM users WHERE email = ?`

// 	// selects ID and hashed password if they exist
// 	err := repo.db.QueryRow(stmt, email).Scan(&id, &hashedPassword)
// 	if err != nil {
// 		return -1, err
// 	}

// 	// compares plaintext password to hashed password
// 	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
// 	if err != nil {
// 		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
// 			return -1, ErrInvalidCredentials
// 		}

// 		return -1, err
// 	}

// 	return id, nil
// }
