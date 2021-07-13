package domain

import (
	"errors"
	"mnimidamonbackend/domain/repository"
)

type ErrDomain error

var (
	ErrNotFound                      ErrDomain = errors.New("ErrNotFound")
	ErrUserWithUsernameAlreadyExists           = errors.New("ErrUserWithUsernameAlreadyExists")
	ErrInternalDomain                          = errors.New("ErrInternalDomain")
	ErrInvalidCredentials                      = errors.New("ErrInvalidCredentials")
	ErrUserNotGroupMember                      = errors.New("ErrUserNotGroupMember")
)

func ToDomainError(err error) error {
	if errors.Is(repository.ErrNotFound, err) {
		return ErrNotFound
	}

	// Unexpected error.
	// TODO logging.
	return ErrInternalDomain
}
