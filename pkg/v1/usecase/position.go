package usecase

import (
	"database/sql"
	"errors"

	"github.com/colin-mcl/stocks/internal/models"
)

// usecase implements the business logic needed when adding instances
// to the mysql database, this file contains the implementation for the
// CRUD operations on the position model

// CreatePosition
//
// Attempts to create the position provided in the positions table, returning
// the id if successful and an error if unsuccessful
// Precondition: p represents a valid position
func (uc *UseCase) CreatePosition(p *models.Position) (int, error) {
	id, err := uc.repo.CreatePosition(p)
	if err != nil {
		return -1, err
	}

	return id, nil
}

// GetPosition
//
// Retrieves the position instance based on ID if it exists, otherwise returns
// an error
func (uc *UseCase) GetPosition(id int) (*models.Position, error) {
	p, err := uc.repo.GetPosition(id)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrDoesNotExist
		}

		return nil, err
	}

	return p, nil
}

// GetPositions
//
// Gets all position instances with matching symbol and heldBy ID and returns
// the result as a slice
func (uc *UseCase) GetPositions(symbol string, owner int) ([]*models.Position,
	error) {

	return uc.repo.GetPositions(symbol, owner)
}

// GetPortfolio
//
// Gets all positions held by owner and retunrs the result as a slice
func (uc *UseCase) GetPortfolio(owner int) ([]*models.Position, error) {
	return uc.repo.GetPortfolio(owner)
}
