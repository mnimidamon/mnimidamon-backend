package modelsql

import "time"

// ComputerBackup denotes that the computer has this backup locally downloaded.
type ComputerBackup struct {
	BackupID        uint `gorm:"primaryKey"`
	GroupComputerID uint `gorm:"primaryKey"`

	Backup        Backup        `gorm:"foreignKey:BackupID"`
	GroupComputer GroupComputer `gorm:"foreignKey:GroupComputerID"`

	CreatedAt time.Time
}
