package models

import (
	"time"
)

// Defines one position held by a user
type Position struct {
	ID            int
	Symbol        string
	HeldBy        int
	PurchasedAt   time.Time
	PurchasePrice float64
	Qty           float64
}

// // Given a valid stock symbol and user in the users database, returns a list
// // of all positions of that stock held by owner
// func (m *PositionModel) GetStock(symbol string, owner int) ([]*Position, error) {
// 	stmt := `SELECT * FROM positions WHERE symbol = ? AND heldBy = ?`

// 	rows, err := m.DB.Query(stmt, strings.ToUpper(symbol), owner)

// 	if err != nil {
// 		return nil, err
// 	}

// 	var stocks []*Position

// 	for rows.Next() {
// 		p := &Position{}

// 		err := rows.Scan(
// 			&p.ID,
// 			&p.Symbol,
// 			&p.HeldBy,
// 			&p.PurchasedAt,
// 			&p.PurchasePrice,
// 			&p.Qty)

// 		if err != nil {
// 			return nil, err
// 		}

// 		stocks = append(stocks, p)
// 	}

// 	if err := rows.Err(); err != nil {
// 		return nil, err
// 	}

// 	return stocks, nil

// }
