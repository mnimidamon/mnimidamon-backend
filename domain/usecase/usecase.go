package usecase

import (
	"mnimidamonbackend/domain/model"
	"mnimidamonbackend/domain/payload"
	"os"
)

type UserRegistrationInterface interface {
	RegisterUser(p payload.UserCredentialsPayload) (*model.User, error)
	ValidateUserCredentials(p payload.UserCredentialsPayload) (*model.User, error)
	// UnregisterUser()
}

type UserInterface interface {
	FindAll() ([]*model.User, error)
	FindById(userID uint) (*model.User, error)
	FindByUsername(username uint) (*model.User, error)
}

type ComputerRegistrationInterface interface {
	RegisterComputer(p payload.ComputerCredentialsPayload) (*model.Computer, error)
}

type ComputerInterface interface {
	FindById(computerID uint) (*model.Computer, error)

	FindAllOfGroup(groupID uint) (*model.Computer, error)
	FindAllOfUser(userID uint) (*model.Computer, error)
}

type GroupComputerInterface interface {
	FindAllOfGroup(groupID uint) ([]*model.GroupComputer, error)
}

type ComputerBackupInterface interface {
	FindAllOfBackup(backupID uint) ([]*model.ComputerBackup, error)
}

type GroupInviteInterface interface {
	AcceptInvite(userID uint, groupID uint) ([]*model.Group, error)
	DeclineInvite(userID uint, groupID uint) ([]*model.Group, error)
}

type InviteInterface interface {
	FindById(userID uint, groupID uint) ([]*model.Invite, error)
	FindAllOfUser(userID uint) ([]*model.Invite, error)
	FindAllOfGroup(groupID uint) ([]*model.Invite, error)
	InviteUser(userID uint) ([]*model.Invite, error)
}

type GroupInterface interface {
	FindAllOfUser(userID uint) ([]*model.Group, error)
	FindByIdOfUser(groupID uint, userID uint) (*model.Group, error)
}

type BackupInterface interface {
	GetGroupBackups(groupID uint) ([]*model.Backup, error)
	InitializeBackup(p payload.InitializeBackupPayload) (*model.Backup, error)
	UploadRequest(backupID uint) (*model.Backup, error)
	DeleteRequest(backupID uint) (*model.Backup, error)
}

type FileInterface interface {
	UploadBackup(backupID uint, data []byte) (*model.Backup, error)
	DownloadBackup(backupID uint) (os.File, error)
}