package models

import (
	"database/sql"
	"strings"
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

// Given a valid stock symbol and user in the users database, returns a list
// of all positions of that stock held by owner
func (m *PositionModel) GetStock(symbol string, owner int) ([]*Position, error) {
	stmt := `SELECT * FROM positions WHERE symbol = ? AND heldBy = ?`

	rows, err := m.DB.Query(stmt, strings.ToUpper(symbol), owner)

	if err != nil {
		return nil, err
	}

	var stocks []*Position

	for rows.Next() {
		p := &Position{}

		err := rows.Scan(
			&p.ID,
			&p.symbol,
			&p.heldBy,
			&p.purchasedAt,
			&p.purchasePrice,
			&p.qty)

		if err != nil {
			return nil, err
		}

		stocks = append(stocks, p)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return stocks, nil

}
