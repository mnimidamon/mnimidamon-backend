package repository

import "errors"

var (
	ErrNotFound = errors.New("ErrNotFound")
	UnknownRepositoryError = errors.New("UnknownRepositoryError")
	ErrTxAlreadyRolledBack = errors.New("ErrTxAlreadyRolledBack")
	ErrUniqueConstraintViolation = errors.New("ErrUniqueConstraintViolation")
	ErrUniquePrimaryKeyConstraintViolation = errors.New("ErrUniquePrimaryKeyConstraintViolation")
)
