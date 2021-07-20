package listgroup

import (
	"mnimidamonbackend/domain"
	"mnimidamonbackend/domain/model"
	"mnimidamonbackend/domain/repository"
	"mnimidamonbackend/domain/usecase"
)

type listGroupUseCase struct {
	GRepo repository.GroupRepository
}

func NewUseCase(gr repository.GroupRepository) usecase.ListGroupInterface {
	return listGroupUseCase{
		GRepo: gr,
	}
}

func (ur listGroupUseCase) FindAllOfUser(userID uint) ([]*model.Group, error) {
	gl, err := ur.GRepo.FindAllOfUser(userID)

	if err != nil {
		return nil, domain.ToDomainError(err)
	}

	return gl, nil
}

func (ur listGroupUseCase) FindById(groupID uint) (*model.Group, error) {
	g, err := ur.GRepo.FindById(groupID)

	if err != nil {
		return nil, domain.ToDomainError(err)
	}

	return g, nil
}

func (ur listGroupUseCase) FindByIdOfUser(groupID uint, userID uint) (*model.Group, error) {
	isMember, err := ur.GRepo.IsMemberOf(userID, groupID)

	if err != nil {
		return nil, domain.ToDomainError(err)
	}

	if !isMember {
		return nil, domain.ErrUserNotGroupMember
	}

	g, err := ur.GRepo.FindById(groupID)

	if err != nil {
		return nil, domain.ToDomainError(err)
	}

	return g, nil
}
