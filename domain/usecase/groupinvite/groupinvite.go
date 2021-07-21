package groupinvite

import (
	"errors"
	"mnimidamonbackend/domain"
	"mnimidamonbackend/domain/model"
	"mnimidamonbackend/domain/repository"
	"mnimidamonbackend/domain/usecase"
	"time"
)

type groupInviteUseCase struct {
	GRepo repository.GroupRepository
	IRepo repository.InviteRepository
	URepo repository.UserRepository
}

func (gi groupInviteUseCase) findUserAndGroup(userID uint, groupID uint) (*model.User, *model.Group, error) {
	// Check if user exists
	u, err := gi.URepo.FindById(userID)
	if errors.Is(err, repository.ErrNotFound) {
		return nil, nil, domain.ErrUserNotFound
	}

	if err != nil {
		return nil, nil, domain.ToDomainError(err)
	}

	// Check if group exists
	g, err := gi.GRepo.FindById(groupID)
	if errors.Is(err, repository.ErrNotFound) {
		return nil, nil, domain.ErrGroupNotFound
	}

	if err != nil {
		return nil, nil, domain.ToDomainError(err)
	}

	return u, g, nil
}

func (gi groupInviteUseCase) AcceptInvite(userID uint, groupID uint) (*model.Group, error) {
	_, g, err := gi.findUserAndGroup(userID, groupID)
	if err != nil {
		return nil, err
	}

	// Check if invitation exists.
	exists, err := gi.IRepo.Exists(userID, groupID)
	if err != nil {
		return nil, domain.ToDomainError(err)
	}

	if !exists {
		return nil, domain.ErrUserNotInvited
	}

	// Transaction on group
	// Transaction on invite
	gtx := gi.GRepo.BeginTx()
	itx := gi.IRepo.BeginTx()

	// Commit group
	_, err = gtx.AddMember(userID, groupID)
	if err != nil {
		gtx.Rollback()
		itx.Rollback()
		return nil, domain.ToDomainError(err)
	}

	// Commit invite
	err = itx.Delete(userID, groupID)
	if err != nil {
		gtx.Rollback()
		itx.Rollback()
		return nil, domain.ToDomainError(err)
	}

	gtx.Commit()
	itx.Commit()

	return g, nil
}

func (gi groupInviteUseCase) DeclineInvite(userID uint, groupID uint) error {
	_, _, err := gi.findUserAndGroup(userID, groupID)
	if err != nil {
		return err
	}

	// Check if invitation exists.
	exists, err := gi.IRepo.Exists(userID, groupID)
	if err != nil {
		return domain.ToDomainError(err)
	}

	if !exists {
		return domain.ErrNotFound
	}

	// Delete the invitation.
	if err := gi.IRepo.Delete(userID, groupID); err != nil {
		return domain.ToDomainError(err)
	}

	return err
}

func (gi groupInviteUseCase) InviteUser(userID uint, groupID uint) (*model.Invite, error) {
	u, g, err := gi.findUserAndGroup(userID, groupID)
	if err != nil {
		return nil, err
	}

	// Check if invitation exists.
	exists, err := gi.IRepo.Exists(userID, groupID)
	if err != nil {
		return nil, domain.ToDomainError(err)
	}

	if exists {
		return nil, domain.ErrAlreadyExists
	}

	isMember, err := gi.GRepo.IsMemberOf(userID, groupID)
	if err != nil {
		return nil, domain.ToDomainError(err)
	}

	if isMember {
		return nil, domain.ErrUserAlreadyGroupMember
	}

	i := &model.Invite{
		UserID:    userID,
		GroupID:   groupID,
		User:      u,
		Group:     g,
		CreatedAt: time.Time{},
	}

	if 	err := gi.IRepo.Create(i); err != nil {
		return nil, domain.ToDomainError(err)
	}

	i.Group = g
	return i, nil
}

func NewUseCase(gr repository.GroupRepository, ir repository.InviteRepository, ur repository.UserRepository) usecase.GroupInviteInterface {
	return groupInviteUseCase{
		GRepo: gr,
		IRepo: ir,
		URepo: ur,
	}
}