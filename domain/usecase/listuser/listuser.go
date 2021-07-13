package listuser

import (
	"errors"
	"mnimidamonbackend/domain"
	"mnimidamonbackend/domain/model"
	"mnimidamonbackend/domain/repository"
	"mnimidamonbackend/domain/usecase"
)

type listUserUseCase struct {
	URepo repository.UserRepository
}


func NewUseCase(ur repository.UserRepository) usecase.ListUserInterface {
	return listUserUseCase{
		URepo: ur,
	}
}

func (lu listUserUseCase) FindAll() ([]*model.User, error) {
	ul, err := lu.URepo.FindAll()

	if err != nil {
		return nil, domain.ErrInternalDomain
	}

	return ul, nil
}

func (lu listUserUseCase) FindById(userID uint) (*model.User, error) {
	u, err := lu.URepo.FindById(userID)

	if errors.Is(err, repository.ErrNotFound) {
		return nil, domain.ErrNotFound
	}

	if err != nil {
		return nil, domain.ErrInternalDomain
	}

	return u, nil
}

func (lu listUserUseCase) FindByUsername(username string) (*model.User, error) {
	u, err := lu.URepo.FindByUsername(username)

	if errors.Is(err, repository.ErrNotFound) {
		return nil, domain.ErrNotFound
	}

	if err != nil {
		return nil, domain.ErrInternalDomain
	}

	return u, nil
}
