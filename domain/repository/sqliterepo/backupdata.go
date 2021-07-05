package sqliterepo

import (
	"gorm.io/gorm"
	"mnimidamonbackend/domain/model"
	"mnimidamonbackend/domain/repository"
	. "mnimidamonbackend/domain/repository/sqliterepo/modelsql"
)

func NewBackupRepository(db *gorm.DB, gr repository.GroupRepository) repository.BackupRepository {
	return backupData{DB: db, GR: gr}
}

type backupData struct {
	*gorm.DB
	GR repository.GroupRepository
}

func (bd backupData) Delete(bm *model.Backup) error {
	result :=
		bd.DB.Delete(&Backup{}, bm.ID)

	if err := result.Error; err != nil {
		return toBusinessLogicError(err)
	}

	return nil
}

func (bd backupData) Update(bm *model.Backup) error {
	b := NewBackupFromBusinessModel(bm)

	result :=
		bd.Model(b).
			Select("upload_request", "delete_request", "on_server", "file_name", "size", "hash").
			Omit("id", "owner_id", "group_id").
			Updates(b)

	if err := result.Error; err != nil {
		return toBusinessLogicError(err)
	}

	b.CopyToBusinessModel(bm)
	return nil
}

func (bd backupData) FindAll(groupID uint) ([]*model.Backup, error) {
	var backups []Backup

	result :=
		bd.Where("group_id = ?", groupID).
			Find(&backups)

	if result.Error != nil {
		return nil, toBusinessLogicError(result.Error)
	}

	var mBackups []*model.Backup
	for _, b := range backups {
		mb := b.NewBusinessModel()
		mBackups = append(mBackups, mb)
	}

	return mBackups, nil
}

func (bd backupData) FindById(backupID uint) (*model.Backup, error) {
	var backup Backup

	result :=
		bd.First(&backup, backupID)


	if err := result.Error; err != nil {
		return nil, toBusinessLogicError(err)
	}

	bm := backup.NewBusinessModel()

	return bm, nil
}

func (bd backupData) Create(bm *model.Backup) error {
	b := NewBackupFromBusinessModel(bm)

	result :=
		bd.Omit("id").
			Create(b)

	if err := result.Error; err != nil {
		return toBusinessLogicError(err)
	}

	b.CopyToBusinessModel(bm)
	return nil
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
