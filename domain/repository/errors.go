package repository

import "errors"

type ErrRepo error

func NewError(msg string) ErrRepo {
	return errors.New(msg)
}

var (
	ErrNotFound                            = NewError("ErrNotFound")
	UnknownRepositoryError                 = NewError("UnknownRepositoryError")
	ErrTxAlreadyRolledBack                 = NewError("ErrTxAlreadyRolledBack")
	ErrUniqueConstraintViolation           = NewError("ErrUniqueConstraintViolation")
	ErrUniquePrimaryKeyConstraintViolation = NewError("ErrUniquePrimaryKeyConstraintViolation")
	ErrForeignKeyConstraintViolation       = NewError("ErrForeignKeyConstraintViolation")
	ErrUserAlreadyInGroupViolation         = NewError("ErrUserAlreadyInGroupViolation")
	ErrInvalidUpdateViolation              = NewError("ErrInvalidUpdateViolation")
	ErrAlreadyExists                       = NewError("ErrAlreadyExists")

	ErrSaveFile     = NewError("ErrSaveFile")
	ErrCreateFile   = NewError("ErrCreateFile")
	ErrOpenFile     = NewError("ErrOpenFile")
	ErrFileDeletion = NewError("ErrFileDeletion")
)
