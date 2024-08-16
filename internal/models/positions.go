package models

import (
	"database/sql"
	"time"
)

// Defines one position held by a user
type Position struct {
	ID            int
	symbol        string
	heldBy        int
	purchasedAt   time.Time
	purchasePrice float64
	qty           float64
}

// Define a position model which wraps a db connection
type PositionModel struct {
	DB *sql.DB
}

// Inserts the provided position into the table and returns its ID if successful
// otherwise returns an error
// Preconditions: `symbol` must be a valid stock symbol, `heldBy` must be the ID
//
//	of a user in the users model, all other fields must be non null
func (m *PositionModel) Insert(symbol string, heldBy int, purchasedAt time.Time,
	purchasePrice float64, qty float64) (int, error) {

	stmt := `INSERT INTO positions (symbol, heldBy, purchasedAt, purchasePrice,
	qty)
	VALUES (?, ?, ?, ?, ?)`

	res, err := m.DB.Exec(stmt, symbol, heldBy, purchasedAt, purchasePrice, qty)

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

// Returns a position with the matching ID
func (m *PositionModel) Get(id int) (*Position, error) {
	stmt := `SELECT * FROM positions WHERE id = ?`

	row := m.DB.QueryRow(stmt, id)

	p := &Position{}

	err := row.Scan(
		&p.ID,
		&p.symbol,
		&p.heldBy,
		&p.purchasedAt,
		&p.purchasePrice,
		&p.qty)

	if err != nil {
		return nil, err
	}

	return p, nil
}
