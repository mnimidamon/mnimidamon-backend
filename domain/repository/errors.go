package repository

import "errors"

var (
	ErrNotFound                            = errors.New("ErrNotFound")
	UnknownRepositoryError                 = errors.New("UnknownRepositoryError")
	ErrTxAlreadyRolledBack                 = errors.New("ErrTxAlreadyRolledBack")
	ErrUniqueConstraintViolation           = errors.New("ErrUniqueConstraintViolation")
	ErrUniquePrimaryKeyConstraintViolation = errors.New("ErrUniquePrimaryKeyConstraintViolation")
	ErrForeignKeyConstraintViolation       = errors.New("ErrForeignKeyConstraintViolation")
	ErrUserNotGroupMemberViolation         = errors.New("ErrUserNotGroupMemberViolation")
	ErrUserAlreadyInGroupViolation         = errors.New("ErrUserAlreadyInGroupViolation")
	ErrInvalidUpdateViolation              = errors.New("ErrInvalidUpdateViolation")
)
