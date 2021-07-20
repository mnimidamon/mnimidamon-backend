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

func (lg listGroupComputerUseCase) FindAllOfGroup(groupID uint) ([]*model.GroupComputer, error) {
	gcs, err := lg.GCRepo.FindAllOfGroup(groupID)
	if err != nil {
		return nil, domain.ToDomainError(err)
	}
	return gcs, nil
}

func NewUseCase(gcr repository.GroupComputerRepository) usecase.ListGroupComputerInterface {
	return listGroupComputerUseCase{
		GCRepo: gcr,
	}
}
