package repository

import "errors"

var (
	ErrNotFound = errors.New("ErrNotFound")
	UnknownRepositoryError = errors.New("UnknownRepositoryError")
	ErrUniqueConstraintViolation = errors.New("ErrUniqueConstraintViolation")
)
