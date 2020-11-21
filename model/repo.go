package model

import "errors"

var (
	// ErrEntityNotFound is an error when an entity is not found.
	ErrEntityNotFound = errors.New("entity not found")
)
