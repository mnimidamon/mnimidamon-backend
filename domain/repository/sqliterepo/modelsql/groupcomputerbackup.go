package modelsql

import "time"

// GroupComputerBackup denotes that the computer has this backup locally downloaded.
type GroupComputerBackup struct {
	BackupID        uint `gorm:"primaryKey"`
	GroupComputerID uint `gorm:"primaryKey"`

	Backup        Backup        `gorm:"foreignKey:BackupID"`
	GroupComputer GroupComputer `gorm:"foreignKey:GroupComputerID"`

	CreatedAt time.Time
}
