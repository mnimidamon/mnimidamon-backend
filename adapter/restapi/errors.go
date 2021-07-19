package restapi

import "errors"

var (
	ErrInvalidComputerAuthToken = errors.New("ErrInvalidComputerAuthToken")
	ErrInvalidUserAuthToken     = errors.New("ErrInvalidUserAuthToken")
)
