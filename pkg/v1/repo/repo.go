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

	// get retrieves the user instance by id if it exists
	GetUser(id int) (*models.User, error)

	// GetUserByEmail gets the user by email if it exist
	GetUserByEmail(email string) (*models.User, error)

	// exists checks whether there exists a user with the given ID
	UserExists(id int) (bool, error)

	// deletes a user by their id, returning an error if unsuccessful
	DeleteUser(id int) error

	// creates a position in the table from the supplied position object,
	// returning the id if successful
	CreatePosition(p *models.Position) (int, error)

	// gets the position with the matching ID if it exists
	GetPosition(id int) (*models.Position, error)
}

type Repo struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) RepoInterface {
	return &Repo{db: db}
}
