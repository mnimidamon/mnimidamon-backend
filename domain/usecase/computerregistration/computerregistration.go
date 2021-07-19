package computerregistration

import (
	"errors"
	"mnimidamonbackend/domain"
	"mnimidamonbackend/domain/model"
	"mnimidamonbackend/domain/repository"
	"mnimidamonbackend/domain/usecase"
	"mnimidamonbackend/domain/usecase/payload"
)

type computerRegistrationUseCase struct {
	CRepo repository.ComputerRepository
	URepo repository.UserRepository
}

func (cr computerRegistrationUseCase) RegisterComputer(p payload.ComputerCredentialsPayload) (*model.Computer, error) {
	// Find the owner
	u, err := cr.URepo.FindById(p.OwnerID)
	if err != nil {
		return nil, domain.ToDomainError(err)
	}

	// Find if name is unique
	_, err = cr.CRepo.FindByName(p.Name, u.ID)

	if err == nil {
		return nil, domain.ErrNameNotUnique
	} else {
		if !errors.Is(err, repository.ErrNotFound) {
			return nil, domain.ToDomainError(err)
		}
	}

	// Create
	cm := &model.Computer{
		Entity:  model.Entity{},
		OwnerID: u.ID,
		Name:    p.Name,
		Owner:   nil,
	}

	if err := cr.CRepo.Create(cm, u.ID); err != nil {
		return nil, domain.ToDomainError(err)
	}

	return cm, nil
}

func NewUseCase(cr repository.ComputerRepository, ur repository.UserRepository) usecase.ComputerRegistrationInterface {
	return computerRegistrationUseCase{
		CRepo: cr,
		URepo: ur,
	}
}


