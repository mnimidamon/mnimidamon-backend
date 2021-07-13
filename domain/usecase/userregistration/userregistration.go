// UserRegistrationInterface use case implementation.
package userregistration

import (
	"errors"
	"mnimidamonbackend/domain"
	"mnimidamonbackend/domain/model"
	"mnimidamonbackend/domain/payload"
	"mnimidamonbackend/domain/repository"
	"mnimidamonbackend/domain/usecase"
)


type userRegistrationUseCase struct {
	URepo repository.UserRepository
}

func NewUseCase(ur repository.UserRepository) usecase.UserRegistrationInterface {
	return userRegistrationUseCase{
		URepo: ur,
	}
}

func (ur userRegistrationUseCase) RegisterUser(p payload.UserCredentialsPayload) (*model.User, error) {
	// Unique username checkpoint.
	userExists, _ := ur.URepo.FindByUsername(p.Username)
	if userExists != nil {
		return nil, domain.ErrUserWithUsernameAlreadyExists
	}

	// New User creation.
	user, err := model.NewUser(p.Username, p.Password)
	if err != nil {
		return nil, domain.ErrInternalDomain
	}

	err = ur.URepo.Create(user)
	if err != nil {
		return nil, domain.ErrInternalDomain
	}

	return user, nil
}

func (ur userRegistrationUseCase) ValidateUserCredentials(p payload.UserCredentialsPayload) (*model.User, error) {
	user, err := ur.URepo.FindByUsername(p.Username)
	// If the user is not found, then the credentials are invalid.
	if errors.Is(repository.ErrNotFound, err) {
		return nil, domain.ErrInvalidCredentials
	}

	// Any other error produces a domain error.
	// TODO: Logging?
	if err != nil {
		return nil, domain.ErrInternalDomain
	}

	// Verify the password.
	if err = user.VerifyPassword(p.Password); err != nil{
		return nil, domain.ErrInvalidCredentials
	}

	return user, nil
}

