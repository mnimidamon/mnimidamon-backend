package domain

import (
	"errors"
	"mnimidamonbackend/domain/repository"
)

type ErrDomain error

var (
	ErrNotFound                      ErrDomain = errors.New("ErrNotFound")
	ErrAlreadyExists                           = errors.New("ErrAlreadyExists")
	ErrUserNotFound                            = errors.New("ErrUserNotFound")
	ErrGroupNotFound                           = errors.New("ErrGroupNotFound")
	ErrUserWithUsernameAlreadyExists           = errors.New("ErrUserWithUsernameAlreadyExists")
	ErrInternalDomain                          = errors.New("ErrInternalDomain")
	ErrInvalidCredentials                      = errors.New("ErrInvalidCredentials")
	ErrUserNotGroupMember                      = errors.New("ErrUserNotGroupMember")
	ErrUserNotInvited                          = errors.New("ErrUserNotInvited")
	ErrUserAlreadyGroupMember                  = errors.New("ErrUserAlreadyGroupMember")
)

func ToDomainError(err error) error {
	if errors.Is(repository.ErrNotFound, err) {
		return ErrNotFound
	}

	// Unexpected error.
	// TODO logging.
	return ErrInternalDomain
}
