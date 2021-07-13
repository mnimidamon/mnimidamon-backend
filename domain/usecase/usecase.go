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

type ListUserInterface interface {
	FindAll() ([]*model.User, error)
	FindById(userID uint) (*model.User, error)
	FindByUsername(username string) (*model.User, error)
}

type ComputerRegistrationInterface interface {
	RegisterComputer(p payload.ComputerCredentialsPayload) (*model.Computer, error)
}

type ListComputerInterface interface {
	FindById(computerID uint) (*model.Computer, error)

	FindAllOfGroup(groupID uint) (*model.Computer, error)
	FindAllOfUser(userID uint) (*model.Computer, error)
}

type ListGroupComputerInterface interface {
	FindAllOfGroup(groupID uint) ([]*model.GroupComputer, error)
}

type ListComputerBackupInterface interface {
	FindAllOfBackup(backupID uint) ([]*model.ComputerBackup, error)
}

type GroupInviteInterface interface {
	AcceptInvite(userID uint, groupID uint) ([]*model.Group, error)
	DeclineInvite(userID uint, groupID uint) ([]*model.Group, error)
	InviteUser(userID uint, groupID uint) ([]*model.Invite, error)
}

type ListInviteInterface interface {
	FindById(userID uint, groupID uint) ([]*model.Invite, error)
	FindAllOfUser(userID uint) ([]*model.Invite, error)
	FindAllOfGroup(groupID uint) ([]*model.Invite, error)
}

type ListGroupInterface interface {
	FindAllOfUser(userID uint) ([]*model.Group, error)
	FindByIdOfUser(groupID uint, userID uint) (*model.Group, error)
}

type ListBackupInterface interface {
	GetGroupBackups(groupID uint) ([]*model.Backup, error)
}

type ManageBackupInterface interface {
	InitializeBackup(p payload.InitializeBackupPayload) (*model.Backup, error)
	UploadRequest(backupID uint) (*model.Backup, error)
	DeleteRequest(backupID uint) (*model.Backup, error)
}

type ManageFileInterface interface {
	UploadBackup(backupID uint, data []byte) (*model.Backup, error)
	DownloadBackup(backupID uint) (os.File, error)
}