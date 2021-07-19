package handlers

import (
	errs "errors"
	"mnimidamonbackend/adapter/restapi/modelapi"
	"mnimidamonbackend/domain"
)

var (
	ErrInternalServer  = &modelapi.Error{Code: "ErrInternalServer"}
	ErrSigningToken  = &modelapi.Error{Code: "ErrSigningToken"}
)

// Common error checker for more readable error.
func IsInternalError(err error) bool {
	if errs.Is(domain.ErrInternalDomain, err) {
		return true
	}
	return false
}

func IsNotFoundError(err error) bool {
	if errs.Is(domain.ErrNotFound, err) {
		return true
	}
	return false
}

func ToRestError(err error) *modelapi.Error {
	return &modelapi.Error{
		Code: err.Error(),
	}
}