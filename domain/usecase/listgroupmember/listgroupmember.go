package listgroupmember

import (
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

func (lgm listGroupMemberUseCase) FindAllMembersOfGroup(groupID uint) ([]*model.User, error) {
	ul, err := lgm.GRepo.FindAllMembers(groupID)

	if err != nil {
		return nil, err
	}

	return ul, nil
}
