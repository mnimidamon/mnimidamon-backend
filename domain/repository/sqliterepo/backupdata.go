package sqliterepo

import (
	. "gorm.io/gorm"
	"mnimidamonbackend/domain/repository"
)

func NewBackupRepository(db *DB) repository.BackupRepository {
	return backupData{db}
}

type backupData struct {
	*DB
}

type BackupDataTx struct {
	backupData
}

func (b BackupDataTx) Rollback() error {
	return b.backupData.Rollback().Error
}

func (b BackupDataTx) Commit() error {
	return b.backupData.Commit().Error
}

func (bd backupData) BeginTx() repository.BackupRepositoryTx {
	return &BackupDataTx{
		backupData{
			DB: bd.Begin(),
		},
	}
}
