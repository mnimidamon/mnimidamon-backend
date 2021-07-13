package listcomputer

import (
	"mnimidamonbackend/domain"
	"mnimidamonbackend/domain/model"
	"mnimidamonbackend/domain/repository"
	"mnimidamonbackend/domain/usecase"
)

type listComputerUseCase struct {
	CRepo repository.ComputerRepository
}

func (lc listComputerUseCase) FindByName(ownerID uint, name string) (*model.Computer, error) {
	c, err := lc.CRepo.FindByName(name, ownerID)

	if err != nil {
		return nil, domain.ToDomainError(err)
	}

	return c, nil
}

func NewUseCase(cr repository.ComputerRepository) usecase.ListComputerInterface {
	return listComputerUseCase{
		CRepo: cr,
	}
}


func (lc listComputerUseCase) FindById(computerID uint) (*model.Computer, error) {
	c, err := lc.CRepo.FindById(computerID)

	if err != nil {
		return nil, domain.ToDomainError(err)
	}

	return c, nil
}

func (lc listComputerUseCase) FindAllOfUser(ownerID uint) ([]*model.Computer, error) {
	cl, err := lc.CRepo.FindAll(ownerID)

	if err != nil {
		return nil, domain.ToDomainError(err)
	}

	return cl, nil
}
