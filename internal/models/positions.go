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
