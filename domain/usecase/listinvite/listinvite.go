package listinvite

import (
	"mnimidamonbackend/domain"
	"mnimidamonbackend/domain/model"
	"mnimidamonbackend/domain/repository"
	"mnimidamonbackend/domain/usecase"
)

type listInviteUseCase struct {
	IRepo repository.InviteRepository
}

func (li listInviteUseCase) FindById(userID uint, groupID uint) (*model.Invite, error) {
	i, err := li.FindById(userID, groupID)

	if err != nil {
		return nil, err
	}

	return i, nil
}

func (li listInviteUseCase) FindAllOfUser(userID uint) ([]*model.Invite, error) {
	il, err := li.IRepo.FindAllOfUser(userID)

	if err != nil {
		return nil, domain.ToDomainError(err)
	}

	return il, nil
}

func (li listInviteUseCase) FindAllOfGroup(groupID uint) ([]*model.Invite, error) {
	il, err := li.IRepo.FindAllOfGroup(groupID)

	if err != nil {
		return nil, domain.ToDomainError(err)
	}

	return il, nil
}

func NewUseCase(ir repository.InviteRepository) usecase.ListInviteInterface {
	return listInviteUseCase{
		IRepo: ir,
	}
}
