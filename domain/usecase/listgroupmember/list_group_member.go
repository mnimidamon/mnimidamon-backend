package listgroupmember

import (
	"mnimidamonbackend/domain"
	"mnimidamonbackend/domain/model"
	"mnimidamonbackend/domain/repository"
	"mnimidamonbackend/domain/usecase"
)

type listGroupMemberUseCase struct {
	GRepo repository.GroupRepository
}

func NewUseCase(gr repository.GroupRepository) usecase.ListGroupMemberInterface {
	return listGroupMemberUseCase{
		GRepo: gr,
	}
}

func (lgm listGroupMemberUseCase) IsMemberOf(userID uint, groupID uint) (bool, error) {
	isMember, err := lgm.GRepo.IsMemberOf(userID, groupID)

	if err != nil {
		return false, domain.ToDomainError(err)
	}

	return isMember, nil
}

func (lgm listGroupMemberUseCase) FindAllMembersOfGroup(groupID uint) ([]*model.User, error) {
	ul, err := lgm.GRepo.FindAllMembers(groupID)

	if err != nil {
		return nil, err
	}

	return ul, nil
}
