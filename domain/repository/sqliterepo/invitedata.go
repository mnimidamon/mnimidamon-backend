package sqliterepo

import (
	"gorm.io/gorm"
	"mnimidamonbackend/domain/repository"
)

func NewInviteRepository(db *gorm.DB, grc repository.GroupRepositoryChecker) repository.InviteRepository {
	return inviteData {
		DB: db,
		GRC: grc,
	}
}

type inviteData struct {
	*gorm.DB
	GRC repository.GroupRepositoryChecker
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
			GRC: id.GRC,
		},
	}
}





