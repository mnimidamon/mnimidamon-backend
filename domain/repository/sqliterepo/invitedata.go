package sqliterepo

import (
	"gorm.io/gorm"
	"mnimidamonbackend/domain/repository"
)

func NewInviteRepository(db *gorm.DB) repository.InviteRepository {
	return inviteData {
		DB: db,
	}
}

type inviteData struct {
	*gorm.DB
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





