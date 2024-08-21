package token

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

// Different error types returned by VerifyToken
var (
	ErrInvalidToken = errors.New("invalid token")
	ErrExpiredToken = errors.New("expired token")
)

// Payload contains the decrypted contents of an auth token
type Payload struct {
	ID        uuid.UUID
	Email     string
	IssuedAt  time.Time
	ExpiredAt time.Time
}

// NewPayload creates a new token payload with the specified fields
func NewPayload(email string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:        tokenID,
		Email:     email,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}

	return payload, nil
}

// Valid checks if the token payload has expired
func (p *Payload) Valid() error {
	if time.Now().After(p.ExpiredAt) {
		return ErrExpiredToken
	}

	return nil
}
