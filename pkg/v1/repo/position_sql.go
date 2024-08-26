package repo

import (
	"github.com/colin-mcl/stocks/internal/models"
)

// Creates the provided position in the table and returns its ID if successful
// otherwise returns an error
// Preconditions: `symbol` must be a valid stock symbol, `heldBy` must be the ID
//
//	of a user in the users model, all other fields must be non null
func (repo *Repo) CreatePosition(p *models.Position) (int, error) {

	stmt := `INSERT INTO positions (symbol, heldBy, purchasedAt, purchasePrice,
	qty)
	VALUES (?, ?, ?, ?, ?)`

	res, err := repo.db.Exec(stmt, p.Symbol, p.HeldBy, p.PurchasedAt,
		p.PurchasePrice, p.Qty)

	// Currently no custom errors for positions
	if err != nil {
		return -1, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return -1, err
	}
	return int(id), nil

}
