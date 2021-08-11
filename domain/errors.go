package domain

import (
	"errors"
	"mnimidamonbackend/domain/constants"
	"mnimidamonbackend/domain/repository"
)

type ErrDomain error

func NewError(msg string) error {
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
	ErrUploadNotRequested            = NewError("ErrUploadNotRequested")
	ErrBackupNotOnServer             = NewError("ErrBackupNotOnServer")
	ErrInvalidFile                   = NewError("ErrInvalidFile")
	ErrInvalidPrefixedHash           = NewError("ErrInvalidPrefixedHash")
)

func ToDomainError(err error) ErrDomain {
	if errors.Is(repository.ErrNotFound, err) {
		return ErrNotFound
	}

	// Unexpected error.
	constants.Log("ErrInternalDomain: %v", err)
	return ErrInternalDomain
}
