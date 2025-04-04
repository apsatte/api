package domain

import "errors"

// common
var (
	ErrDatabase = errors.New("DATABASE_ERROR")
	ErrNotFound = errors.New("NOT_FOUND")

	ErrUnauthorized = errors.New("UNAUTHORIZED")
)

// customers
var (
	ErrIncorrectPassword = errors.New("INCORRECT_PASSWORD")
	ErrNameTooLong       = errors.New("NAME_TOO_LONG")
)
