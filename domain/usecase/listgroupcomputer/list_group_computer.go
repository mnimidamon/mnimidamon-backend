package listgroupcomputer

import (
	"mnimidamonbackend/domain"
	"mnimidamonbackend/domain/model"
	"mnimidamonbackend/domain/repository"
	"mnimidamonbackend/domain/usecase"
)

type listGroupComputerUseCase struct {
	GCRepo repository.GroupComputerRepository
}

func (lg listGroupComputerUseCase) FindAllOfComputer(computerID uint) ([]*model.GroupComputer, error) {
	gcs, err := lg.GCRepo.FindAllOfComputer(computerID)
	if err != nil {
		return nil, domain.ToDomainError(err)
	}
	return gcs, nil
}

func (lg listGroupComputerUseCase) FindAllOfGroup(groupID uint) ([]*model.GroupComputer, error) {
	gcs, err := lg.GCRepo.FindAllOfGroup(groupID)
	if err != nil {
		return nil, domain.ToDomainError(err)
	}
	return gcs, nil
}

func (lg listGroupComputerUseCase) FindById(groupID uint, computerID uint) (*model.GroupComputer, error) {
	gc, err := lg.GCRepo.FindById(groupID, computerID)
	if err != nil {
		return nil, domain.ToDomainError(err)
	}
	return gc, nil
}

func NewUseCase(gcr repository.GroupComputerRepository) usecase.ListGroupComputerInterface {
	return listGroupComputerUseCase{
		GCRepo: gcr,
	}
}
