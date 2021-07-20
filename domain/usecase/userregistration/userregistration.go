// UserRegistrationInterface use case implementation.
package userregistration

import (
	"errors"
	"mnimidamonbackend/domain"
	"mnimidamonbackend/domain/model"
	"mnimidamonbackend/domain/repository"
	"mnimidamonbackend/domain/usecase"
	"mnimidamonbackend/domain/usecase/payload"
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
	_, err := ur.URepo.FindByUsername(p.Username)
	if err != nil {
		if !errors.Is(err, repository.ErrNotFound) {
			return nil, domain.ToDomainError(err)
		}
	} else {
		return nil, domain.ErrUserWithUsernameAlreadyExists
	}

	// New User creation.
	user, err := model.NewUser(p.Username, p.Password)
	if err != nil {
		return nil, domain.ToDomainError(err)
	}

	err = ur.URepo.Create(user)
	if err != nil {
		return nil, domain.ToDomainError(err)
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

