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

	// Deletes the user instance
	DeleteUser(id int) error
}

type UseCase struct {
	repo repo.RepoInterface
}

func NewUC(repo repo.RepoInterface) UseCaseInterface {
	return &UseCase{repo}
}
