package repo

import (
	"database/sql"

	"github.com/colin-mcl/stocks/internal/models"
)

// repo defines the repo interface, struct and new function to create a repo for
// data access to sql table

// RepoInterface is the interface for the data access layer of the application
type RepoInterface interface {
	// creates a user with the data supplied and returns the id if successful
	CreateUser(u *models.User) (int, error)

	// get retrieves the user instance
	GetUser(email string) (*models.User, error)

	// exists checks whether there exists a user with the given ID
	UserExists(id int) (bool, error)

	// authenticate verifies that a user with the provided email and password
	// exists and returns the ID if so
	Authenticate(email, password string) (int, error)
}

type Repo struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) RepoInterface {
	return &Repo{db: db}
}
