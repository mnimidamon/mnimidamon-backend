package modelsql

import (
	"mnimidamonbackend/domain/model"
	"time"
)

// ComputerBackup denotes that the computer has this backup locally downloaded.
type ComputerBackup struct {
	BackupID        uint `gorm:"primaryKey"`
	GroupComputerID uint `gorm:"primaryKey"`

	Backup        Backup        `gorm:"foreignKey:BackupID"`
	GroupComputer GroupComputer `gorm:"foreignKey:GroupComputerID"`

	CreatedAt time.Time
}

func NewComputerBackupFromBusinessModel(cbm *model.ComputerBackup) *ComputerBackup {
	if cbm == nil {
		return nil
	}

	gc := &ComputerBackup{
		BackupID:        cbm.BackupID,
		GroupComputerID: cbm.GroupComputerID,
		Backup:          *NewBackupFromBusinessModel(&cbm.Backup),
		GroupComputer:   *NewGroupComputerFromBusinessModel(&cbm.GroupComputer),
		CreatedAt:       time.Time{},
	}

	return gc
}

func (cb *ComputerBackup) NewBusinessModel() *model.ComputerBackup {
	if cb == nil {
		return nil
	}

	cbm := new(model.ComputerBackup)
	cb.CopyToBusinessModel(cbm)
	return cbm
}

func (cb *ComputerBackup) CopyToBusinessModel(cbm *model.ComputerBackup) {
	if cb == nil {
		cbm = nil
		return
	}

	cbm.BackupID = cb.BackupID
	cbm.GroupComputerID = cb.GroupComputerID
	cbm.Backup = *cb.Backup.NewBusinessModel()
	cbm.GroupComputer = *cb.GroupComputer.NewBusinessModel()
}