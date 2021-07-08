package usecase

import (
	"mnimidamonbackend/domain/model"
	"mnimidamonbackend/domain/payload"
)

type UserRegistrationInterface interface {
	RegisterUser(p payload.UserCredentialsPayload) (*model.User, error)
	ValidateUserCredentials(p payload.UserCredentialsPayload) (*model.User, error)
}