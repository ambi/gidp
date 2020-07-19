package model

import "github.com/google/uuid"

// NewUUID generates a new UUID (string).
func NewUUID() string {
	return uuid.New().String()
}
