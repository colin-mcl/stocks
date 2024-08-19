package repo

import (
	"database/sql"
	"errors"
	"strings"

	"github.com/colin-mcl/stocks/internal/models"
	"github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

// user_sql implements the user methods on the repo

var (
	ErrAlreadyExists      error = errors.New("error: user already exists")
	ErrInvalidCredentials error = errors.New("error: invalid credentials")
)

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

	// checks if user with this email already exists
	if err != nil {
		var mySQLError *mysql.MySQLError
		if errors.As(err, &mySQLError) {
			if mySQLError.Number == 1062 &&
				strings.Contains(mySQLError.Message, "users_uc_email") {
				return -1, ErrAlreadyExists
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

// GetUser
//
// Gets the user with the provided email if it exists in the db
func (repo *Repo) GetUser(email string) (*models.User, error) {
	stmt := `first_name, last_name, created_at FROM users
	WHERE email = ?`

	row := repo.db.QueryRow(stmt, email)

	u := &models.User{}

	err := row.Scan(
		&u.ID,
		&u.Username,
		&u.Email,
		&u.HashedPassword,
		&u.FirstName,
		&u.LastName,
		&u.CreatedAt)

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

// Authenticate
//
// Verifies that a user with the provided email and password exists and returns
// the ID if so
func (repo *Repo) Authenticate(email, password string) (int, error) {
	var id int
	var hashedPassword []byte

	stmt := `SELECT id, hashed_password FROM users WHERE email = ?`

	// selects ID and hashed password if they exist
	err := repo.db.QueryRow(stmt, email).Scan(&id, &hashedPassword)
	if err != nil {
		return -1, err
	}

	// compares plaintext password to hashed password
	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return -1, ErrInvalidCredentials
		}

		return -1, err
	}

	return id, nil
}
