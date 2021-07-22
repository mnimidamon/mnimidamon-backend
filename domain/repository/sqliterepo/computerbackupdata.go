package sqliterepo

import (
	"errors"
	"gorm.io/gorm"
	"mnimidamonbackend/domain/model"
	"mnimidamonbackend/domain/repository"
	. "mnimidamonbackend/domain/repository/sqliterepo/modelsql"
)

func NewComputerBackupRepository(db *gorm.DB) repository.ComputerBackupRepository {
	return computerBackupData {
		DB: db,
	}
}

type computerBackupData struct {
	*gorm.DB
}

func (cbd computerBackupData) ContinueTx(mr repository.TransactionContextReader) repository.ComputerBackupRepositoryTx {
	meta := mr.GetContext()

	if dbtx, isDB := meta.(*gorm.DB); isDB {
		return computerBackupDataTx{
			computerBackupData{
				DB: dbtx,
			},
		}
	}

	return cbd.BeginTx()
}

func (cbd computerBackupData) GetContext() interface{} {
	return cbd.DB
}

func (cbd computerBackupData) FindById(groupComputerID uint, backupID uint) (*model.ComputerBackup, error) {
	var cb ComputerBackup

	err :=
		cbd.Model(&ComputerBackup{}).
			Where("group_computer_id = ? AND backup_id = ?", groupComputerID, backupID).
			First(&cb).Error

	if err != nil {
		return nil, toRepositoryError(err)
	}

	cbm := cb.NewBusinessModel()
	return cbm, nil
}

func (cbd computerBackupData) FindAllOfGroupComputer(groupComputerID uint) ([]*model.ComputerBackup, error) {
	var cbackups []ComputerBackup

	result :=
		cbd.Where("group_computer_id = ?", groupComputerID).
			Find(&cbackups)

	if result.Error != nil {
		return nil, toRepositoryError(result.Error)
	}

	var mCBackups []*model.ComputerBackup
	for _, c := range cbackups {
		cm := c.NewBusinessModel()
		mCBackups = append(mCBackups, cm)
	}

	return mCBackups, nil
}

func (cbd computerBackupData) FindAllOfBackup(backupID uint) ([]*model.ComputerBackup, error) {
	var cbackups []ComputerBackup

	result :=
		cbd.Where("backup_id = ?", backupID).
			Find(&cbackups)

	if result.Error != nil {
		return nil, toRepositoryError(result.Error)
	}

	var mCBackups []*model.ComputerBackup
	for _, c := range cbackups {
		cm := c.NewBusinessModel()
		mCBackups = append(mCBackups, cm)
	}

	return mCBackups, nil
}

func (cbd computerBackupData) Create(cbm *model.ComputerBackup) error {
	cb := NewComputerBackupFromBusinessModel(cbm)

	if exists, _ := cbd.Exists(cb.GroupComputerID, cb.BackupID); exists {
		return repository.ErrAlreadyExists
	}

	result :=
		cbd.Omit("id").
			Create(cbm)

	if result.Error != nil {
		return toRepositoryError(result.Error)
	}

	cb.CopyToBusinessModel(cbm)
	return nil

}

func (cbd computerBackupData) Delete(groupComputerID uint, backupID uint) error {
	result :=
		cbd.DB.
			Where("group_computer_id = ? AND backup_id = ?", groupComputerID, backupID).
			Delete(&ComputerBackup{})

	if result.Error != nil {
		return toRepositoryError(result.Error)
	}

	return nil
}

func (cbd computerBackupData) Exists(groupComputerID uint, backupID uint) (bool, error) {
	_, err := cbd.FindById(groupComputerID, backupID)

	if err != nil  {
		if  errors.Is(repository.ErrNotFound, err) {
			return false, nil
		}
		return false, toRepositoryError(err)
	}

	return true, nil
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

