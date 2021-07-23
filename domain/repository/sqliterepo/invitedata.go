package sqliterepo

import (
	"errors"
	"gorm.io/gorm"
	"mnimidamonbackend/domain/model"
	"mnimidamonbackend/domain/repository"
	. "mnimidamonbackend/domain/repository/sqliterepo/modelsql"
)

func NewInviteRepository(db *gorm.DB) repository.InviteRepository {
	return inviteData {
		DB: db,
	}
}

type inviteData struct {
	*gorm.DB
}

func (id inviteData) GetContext() interface{} {
	return id.DB
}

func (id inviteData) ContinueTx(mr repository.TransactionContextReader) repository.InviteRepositoryTx {
	meta := mr.GetContext()
	// If the meta is a database then set it here.
	// This is because of the database locking for sqlite.
	if dbtx, isDB := meta.(*gorm.DB); isDB {
		return inviteDataTx{
			inviteData{
				DB: dbtx,
			},
		}
	}

	// Else return a normal transaction.
	return id.BeginTx()
}

func (id inviteData) Exists(userID uint, groupID uint) (bool, error) {
	_, err := id.FindById(userID, groupID)

	if err != nil  {
		if  errors.Is(repository.ErrNotFound, err) {
			return false, nil
		}
		return false, toRepositoryError(err)
	}

	return true, nil
}

func (id inviteData) Create(im *model.Invite) error {
	i := NewInviteFromBusinessModel(im)

	if exists, _ := id.Exists(i.UserID, i.GroupID); exists {
		return repository.ErrAlreadyExists
	}

	result :=
		id.Omit("id").
			Create(i)

	if result.Error != nil {
		return toRepositoryError(result.Error)
	}

	i.CopyToBusinessModel(im)
	return nil
}

func (id inviteData) Delete(userID uint, groupID uint) error {
	result :=
		id.DB.
			Where("user_id = ? AND group_id = ?", userID, groupID).
			Delete(&Invite{})

	if result.Error != nil {
		return toRepositoryError(result.Error)
	}

	return nil
}

func (id inviteData) FindAllOfGroup(groupID uint) ([]*model.Invite, error) {
	var invites []Invite

	result :=
		id.Where("group_id = ?", groupID).
			Preload("User").
			Find(&invites)

	if result.Error != nil {
		return nil, toRepositoryError(result.Error)
	}

	var mInvites []*model.Invite
	for _, i := range invites {
		mi := i.NewBusinessModel()
		mInvites = append(mInvites, mi)
	}

	return mInvites, nil
}

func (id inviteData) FindAllOfUser(userID uint) ([]*model.Invite, error) {
	var invites []Invite

	result :=
		id.Where("user_id = ?", userID).
			Preload("Group").
			Find(&invites)

	if result.Error != nil {
		return nil, toRepositoryError(result.Error)
	}

	var mInvites []*model.Invite
	for _, i := range invites {
		mi := i.NewBusinessModel()
		mInvites = append(mInvites, mi)
	}

	return mInvites, nil
}

func (id inviteData) FindById(userID uint, groupID uint) (*model.Invite, error) {
	var invite Invite

	result :=
		id.Model(&Invite{}).
			Where("user_id = ? AND group_id = ?", userID, groupID).
			Preload("Group").
			First(&invite)

	if err := result.Error; err != nil {
		return nil, toRepositoryError(err)
	}

	im := invite.NewBusinessModel()
	return im, nil
}

type inviteDataTx struct {
	inviteData
}

func (idtx inviteDataTx) Rollback() error {
	return idtx.inviteData.DB.Rollback().Error
}

func (idtx inviteDataTx) Commit() error {
	return idtx.inviteData.DB.Commit().Error
}

func (id inviteData) BeginTx() repository.InviteRepositoryTx {
	return inviteDataTx{
		inviteData{
			DB: id.Begin(),
		},
	}
}






