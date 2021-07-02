package model

import "time"

// GroupComputerBackup denotes that the computer has this backup locally downloaded.
type GroupComputerBackup struct {
	BackupID        uint
	GroupComputerID uint

	Backup        Backup
	GroupComputer GroupComputer

	CreatedAt time.Time
}
