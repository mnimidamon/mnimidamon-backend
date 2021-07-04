package sqliterepo

import (
	"gorm.io/gorm"
	"mnimidamonbackend/domain/model"
	"mnimidamonbackend/domain/repository"
)

func NewBackupRepository(db *gorm.DB) repository.BackupRepository {
	return backupData{db}
}

type backupData struct {
	*gorm.DB
}

func (bd backupData) Delete(bm *model.Backup) error {
	panic("implement me")
}

func (bd backupData) Update(bm *model.Backup) error {
	panic("implement me")
}

func (bd backupData) FindAll() ([]*model.Backup, error) {
	panic("implement me")
}

func (bd backupData) FindById(backupID int) (*model.Backup, error) {
	panic("implement me")
}

func (bd backupData) Create(bm *model.Backup) error {
	panic("implement me")
}

// Transaction support.
type BackupDataTx struct {
	backupData
}

func (bdtx BackupDataTx) Rollback() error {
	return bdtx.backupData.Rollback().Error
}

func (bdtx BackupDataTx) Commit() error {
	return bdtx.backupData.Commit().Error
}

func (bd backupData) BeginTx() repository.BackupRepositoryTx {
	return &BackupDataTx{
		backupData{
			DB: bd.Begin(),
		},
	}
}
