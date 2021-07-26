package usecase

import (
	"io"
	"mnimidamonbackend/domain/model"
	"mnimidamonbackend/domain/usecase/payload"
)

type UserRegistrationInterface interface {
	RegisterUser(p payload.UserCredentialsPayload) (*model.User, error)
	ValidateUserCredentials(p payload.UserCredentialsPayload) (*model.User, error)
	// UnregisterUser()
}

type ManageGroupInterface interface {
	CreateGroup(p payload.CreateGroupPayload) (*model.Group, error)
}

type ListUserInterface interface {
	FindAll() ([]*model.User, error)
	FindById(userID uint) (*model.User, error)
	FindByUsername(username string) (*model.User, error)
}

type ListGroupMemberInterface interface {
	FindAllMembersOfGroup(groupID uint) ([]*model.User, error)
	IsMemberOf(userID uint, groupID uint) (bool, error)
}

type ComputerRegistrationInterface interface {
	RegisterComputer(p payload.ComputerCredentialsPayload) (*model.Computer, error)
}

type ManageGroupComputerInterface interface {
	Update(groupComputerID uint, size uint) (*model.GroupComputer, error)
	JoinGroup(computerID uint, size uint, groupID uint) (*model.GroupComputer, error)
	LeaveGroup(computerID uint, size uint) error
}

type ListComputerInterface interface {
	FindById(computerID uint) (*model.Computer, error)
	FindAllOfUser(ownerID uint) ([]*model.Computer, error)
	FindByName(ownerID uint, name string) (*model.Computer, error)
}

type ListGroupComputerInterface interface {
	FindAllOfGroup(groupID uint) ([]*model.GroupComputer, error)
	FindById(groupID uint, computerID uint) (*model.GroupComputer, error)
}

type ListComputerBackupInterface interface {
	FindAllOfBackup(backupID uint) ([]*model.ComputerBackup, error)
}

type ManageGroupBackupInterface interface {
	LogDownload(backupID uint, computerID uint, prefix string, hash string) (*model.ComputerBackup, error)
}

type GroupInviteInterface interface {
	AcceptInvite(userID uint, groupID uint) (*model.Group, error)
	DeclineInvite(userID uint, groupID uint) error
	InviteUser(userID uint, groupID uint) (*model.Invite, error)
}

type ListInviteInterface interface {
	FindById(userID uint, groupID uint) (*model.Invite, error)
	FindAllOfUser(userID uint) ([]*model.Invite, error)
	FindAllOfGroup(groupID uint) ([]*model.Invite, error)
}

type ListGroupInterface interface {
	FindById(groupID uint) (*model.Group, error)
	FindAllOfUser(userID uint) ([]*model.Group, error)
	FindByIdOfUser(groupID uint, userID uint) (*model.Group, error)
}

type ListBackupInterface interface {
	FindGroupBackups(groupID uint) ([]*model.Backup, error)
	FindById(backupID uint) (*model.Backup, error)
}

type ManageBackupInterface interface {
	InitializeBackup(p payload.InitializeBackupPayload) (*model.Backup, error)
	UploadRequest(ownerID uint, backupID uint) (*model.Backup, error)
	DeleteRequest(userID uint, backupID uint) (*model.Backup, error)
}

type ManageFileInterface interface {
	UploadBackup(backupID uint, rc io.ReadCloser) (*model.Backup, error)
	DownloadBackup(backupID uint) (io.ReadCloser, error)
}