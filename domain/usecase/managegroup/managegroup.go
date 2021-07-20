// This package has the usecase.ManageGroupInterface implementation.
package managegroup

import (
	"errors"
	"mnimidamonbackend/domain"
	"mnimidamonbackend/domain/model"
	"mnimidamonbackend/domain/repository"
	"mnimidamonbackend/domain/usecase"
	"mnimidamonbackend/domain/usecase/payload"
)

type manageGroupUseCase struct {
	GRepo repository.GroupRepository
	URepo repository.UserRepository
}

func (mg manageGroupUseCase) CreateGroup(p payload.CreateGroupPayload) (*model.Group, error) {
	memberID, name := p.MemberID, p.Name

	// Unique name check.
	if	_, err := mg.GRepo.FindByName(name); err != nil {
		if !errors.Is(err, repository.ErrNotFound) {
			return nil, domain.ToDomainError(err)
		}
	} else {
		return nil, domain.ErrGroupWithNameAlreadyExists
	}

	// Does the user exist?
	if 	_, err := mg.URepo.FindById(memberID); err != nil {
		return nil, domain.ToDomainError(err)
	}

	gm := &model.Group{
		Entity:       model.Entity{},
		Name:         name,
		GroupMembers: nil,
		Invites:      nil,
	}

	grtx := mg.GRepo.BeginTx()

	if err := grtx.Create(gm); err != nil {
		grtx.Rollback()
		return nil, domain.ToDomainError(err)
	}

	if _,  err := grtx.AddMember(memberID, gm.ID); err != nil {
		return nil, domain.ToDomainError(err)
	}

	grtx.Commit()

	return gm, nil
}

func NewUseCase(ur repository.UserRepository, gr repository.GroupRepository) usecase.ManageGroupInterface {
	return manageGroupUseCase{
		GRepo: gr,
		URepo: ur,
	}
}