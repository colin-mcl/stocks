package token

import "time"

// Maker is an interfface for managing authentication tokens
type Maker interface {
	// Creates a new token for a specific username and duration
	CreateToken(username string, duration time.Duration) (string, error)

	// Check if the input token is valid or not and return the payload if so
	VerifyToken(token string) (*Payload, error)
}
