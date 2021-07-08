package domain

import "errors"

type ErrDomain error

var (
	ErrNotFound                      ErrDomain = errors.New("ErrNotFound")
	ErrUserWithUsernameAlreadyExists           = errors.New("ErrUserWithUsernameAlreadyExists")
	ErrPasswordHash                            = errors.New("ErrPasswordHash")
	ErrInternalDomain                          = errors.New("ErrInternalDomain")
	ErrInvalidCredentials                      = errors.New("ErrInvalidCredentials")
)
