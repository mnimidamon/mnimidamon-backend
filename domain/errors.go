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
	ErrAlreadyExists                 = NewError("ErrAlreadyExists")
	ErrUserNotFound                  = NewError("ErrUserNotFound")
	ErrGroupNotFound                 = NewError("ErrGroupNotFound")
	ErrUserWithUsernameAlreadyExists = NewError("ErrUserWithUsernameAlreadyExists")
	ErrInternalDomain                = NewError("ErrInternalDomain")
	ErrInvalidCredentials            = NewError("ErrInvalidCredentials")
	ErrUserNotGroupMember            = NewError("ErrUserNotGroupMember")
	ErrUserNotInvited                = NewError("ErrUserNotInvited")
	ErrUserAlreadyGroupMember        = NewError("ErrUserAlreadyGroupMember")
)

func ToDomainError(err error) ErrDomain {
	if errors.Is(repository.ErrNotFound, err) {
		return ErrNotFound
	}

	// Unexpected error.
	// TODO logging.
	return ErrInternalDomain
}
