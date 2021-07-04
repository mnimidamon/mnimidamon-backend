package modelsql

// Backup denotes a file that has been uploaded for backup by the Owner inside a Group.
// Flags:
//		- UploadRequest: file is requested to be uploaded from GroupComputerBackup owners (new member, more space, ...)
//		- DeleteRequest: file will no longer be backed up, all local storages GroupComputerBackup are up for deletion.
//		- OnServer: file is locally stored on server.
type Backup struct {
	Entity

	FileName string
	Size     uint

	UploadRequest bool `gorm:"default:true"`
	DeleteRequest bool `gorm:"default:false"`
	OnServer      bool `gorm:"default:false"`

	OwnerID uint
	GroupID uint

	Owner User  `gorm:"foreignKey:OwnerID"`
	Group Group `gorm:"foreignKey:GroupID"`
}
