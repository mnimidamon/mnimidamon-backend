package sqliterepo

import (
	"gorm.io/gorm"
	"mnimidamonbackend/domain/model"
	"mnimidamonbackend/domain/repository"
)

func NewComputerBackupRepository(db *gorm.DB) repository.ComputerBackupRepository {
	return computerBackupData {
		DB: db,
	}
}

type computerBackupData struct {
	*gorm.DB
}

func (cbd computerBackupData) FindById(computerID uint, backupID uint) (*model.ComputerBackup, error) {
	panic("implement me")
}

func (cbd computerBackupData) FindAllOfComputer(computerID uint) ([]*model.ComputerBackup, error) {
	panic("implement me")
}

func (cbd computerBackupData) FindAllOfBackup(backupID uint) ([]*model.ComputerBackup, error) {
	panic("implement me")
}

func (cbd computerBackupData) Create(cbm *model.ComputerBackup) error {
	panic("implement me")
}

func (cbd computerBackupData) Delete(computerID uint, backupID uint) error {
	panic("implement me")
}

func (cbd computerBackupData) Exists(computerID uint, backupID uint) (bool, error) {
	panic("implement me")
}

type computerBackupDataTx struct {
	computerBackupData
}

func (cbdtx computerBackupDataTx) Rollback() error {
	return cbdtx.computerBackupData.DB.Rollback().Error
}

func (cbdtx computerBackupDataTx) Commit() error {
	return cbdtx.computerBackupData.DB.Commit().Error
}

func (cbd computerBackupData) BeginTx() repository.ComputerBackupRepositoryTx {
	return computerBackupDataTx{
		computerBackupData{
			DB: cbd.Begin(),
		},
	}
}

