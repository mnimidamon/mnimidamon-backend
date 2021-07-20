package domain

import (
	"errors"
	"mnimidamonbackend/domain/repository"
)

type ErrDomain error

func NewError(msg string) ErrDomain {
	return errors.New(msg)
}

var (
	ErrNotFound                      = NewError("ErrNotFound")
	ErrUserNotFound                  = NewError("ErrUserNotFound")
	ErrGroupNotFound                 = NewError("ErrGroupNotFound")
	ErrComputerNotFound              = NewError("ErrComputerNotFound")
	ErrBackupNotFound                = NewError("ErrBackupNotFound")
	ErrBackupAlreadyOnServer         = NewError("ErrBackupAlreadyOnServer")
	ErrBackupWaitingForDeletion      = NewError("ErrBackupWaitingForDeletion")
	ErrAlreadyExists                 = NewError("ErrAlreadyExists")
	ErrUserWithUsernameAlreadyExists = NewError("ErrUserWithUsernameAlreadyExists")
	ErrGroupWithNameAlreadyExists    = NewError("ErrGroupWithNameAlreadyExists")
	ErrInternalDomain                = NewError("ErrInternalDomain")
	ErrInvalidCredentials            = NewError("ErrInvalidCredentials")
	ErrUserNotGroupMember            = NewError("ErrUserNotGroupMember")
	ErrUserNotInvited                = NewError("ErrUserNotInvited")
	ErrUserNotOwner                  = NewError("ErrUserNotOwner")
	ErrUserAlreadyGroupMember        = NewError("ErrUserAlreadyGroupMember")
	ErrComputerAlreadyMember         = NewError("ErrComputerAlreadyMember")
	ErrInvalidBackupSize             = NewError("ErrInvalidBackupSize")
	ErrInvalidFileName               = NewError("ErrInvalidFileName")
	ErrNameNotUnique                 = NewError("ErrNameNotUnique")
)

func ToDomainError(err error) ErrDomain {
	if errors.Is(repository.ErrNotFound, err) {
		return ErrNotFound
	}

	// Unexpected error.
	// TODO logging.
	return ErrInternalDomain
}
