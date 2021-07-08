package repository

import "errors"

type ErrRepo error

var (
	ErrNotFound                            ErrRepo = errors.New("ErrNotFound")
	UnknownRepositoryError                         = errors.New("UnknownRepositoryError")
	ErrTxAlreadyRolledBack                         = errors.New("ErrTxAlreadyRolledBack")
	ErrUniqueConstraintViolation                   = errors.New("ErrUniqueConstraintViolation")
	ErrUniquePrimaryKeyConstraintViolation         = errors.New("ErrUniquePrimaryKeyConstraintViolation")
	ErrForeignKeyConstraintViolation               = errors.New("ErrForeignKeyConstraintViolation")
	ErrUserAlreadyInGroupViolation                 = errors.New("ErrUserAlreadyInGroupViolation")
	ErrInvalidUpdateViolation                      = errors.New("ErrInvalidUpdateViolation")
	ErrAlreadyExists                               = errors.New("ErrAlreadyExists")
)
