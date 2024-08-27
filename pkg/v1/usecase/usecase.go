package usecase

import (
	"github.com/colin-mcl/stocks/internal/models"
	"github.com/colin-mcl/stocks/pkg/v1/repo"
)

// the usecase package contains the business logic of application and
// utilizes the repository interface for CRUD operations

type UseCaseInterface interface {
	// creates a user with the data supplied
	CreateUser(u *models.User) (int, error)

	// Retrieves the user instance
	GetUser(id int) (*models.User, error)

	// Retrieve the user instance by their email
	GetUserByEmail(email string) (*models.User, error)

	// Deletes the user instance
	DeleteUser(id int) error

	// Creates the position instance in the table
	CreatePosition(p *models.Position) (int, error)

	// Gets the position instance given its ID
	GetPosition(id int) (*models.Position, error)

	// Gets all positions matching the symbol and owner ID
	GetPositions(symbol string, owner int) ([]*models.Position, error)
}

type UseCase struct {
	repo repo.RepoInterface
}

func NewUC(repo repo.RepoInterface) UseCaseInterface {
	return &UseCase{repo}
}
