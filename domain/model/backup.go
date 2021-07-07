package model

// Backup denotes a file that has been uploaded for backup by the Owner inside a Group.
// Flags:
//		- UploadRequest: file is requested to be uploaded from ComputerBackup owners (new member, more space, ...)
//		- DeleteRequest: file will no longer be backed up, all local storages ComputerBackup are up for deletion.
//		- OnServer: file is locally stored on server.
type Backup struct {
	Entity

	FileName string
	Size     uint
	Hash     string

	UploadRequest bool
	DeleteRequest bool
	OnServer      bool

	OwnerID uint
	GroupID uint

	Owner User
	Group Group
}
