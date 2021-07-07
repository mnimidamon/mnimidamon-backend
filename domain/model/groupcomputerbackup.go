package model

import "time"

// ComputerBackup denotes that the computer has this backup locally downloaded.
type ComputerBackup struct {
	BackupID        uint
	GroupComputerID uint

	Backup        Backup
	GroupComputer GroupComputer

	CreatedAt time.Time
}
