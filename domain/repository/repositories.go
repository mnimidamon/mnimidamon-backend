package repository

import (
	"mnimidamonbackend/domain/model"
)

// Repositories are used for reading and writing models.
// Repositories Tx objects are used for when a transaction is initiated.

type UserRepository interface {
	FindAll() ([]*model.User, error)
	FindById(userID uint) (*model.User, error)
	FindByUsername(username string) (*model.User, error)

	Create(um *model.User) error
	Delete(userID uint) error
	Update(um *model.User) error

	Exists(userID uint) (bool, error)

	BeginTx() UserRepositoryTx
	ContinueTx(mr TransactionContextReader) UserRepositoryTx
	TransactionContextReader
}

type GroupRepository interface {
	FindAll() ([]*model.Group, error)
	FindAllMembers(groupID uint) ([]*model.User, error)
	FindAllOfUser(userID uint) ([]*model.Group, error)
	FindById(groupID uint) (*model.Group, error)
	FindByName(name string) (*model.Group, error)

	Create(gm *model.Group) error
	Delete(groupID uint) error
	Update(gm *model.Group) error

	AddMember(userID uint, groupID uint) (*model.Group, error)
	IsMemberOf(userID uint, groupID uint) (bool, error)
	Exists(groupID uint) (bool, error)

	BeginTx() GroupRepositoryTx
	ContinueTx(mr TransactionContextReader) GroupRepositoryTx
	TransactionContextReader
}

type BackupRepository interface {
	FindAll(groupID uint) ([]*model.Backup, error)
	FindById(backupID uint) (*model.Backup, error)

	Create(bm *model.Backup) error
	Delete(backupID uint) error
	Update(bm *model.Backup) error

	BeginTx() BackupRepositoryTx
	ContinueTx(mr TransactionContextReader) BackupRepositoryTx
	TransactionContextReader
}

type InviteRepository interface {
	Create(im *model.Invite) error
	Delete(userID uint, groupID uint) error

	FindAllOfGroup(groupID uint) ([]*model.Invite, error)
	FindAllOfUser(userID uint) ([]*model.Invite, error)
	FindById(userID uint, groupID uint) (*model.Invite, error)

	Exists(userID uint, groupID uint) (bool, error)

	BeginTx() InviteRepositoryTx
	ContinueTx(mr TransactionContextReader) InviteRepositoryTx
	TransactionContextReader
}

type ComputerRepository interface {
	FindAll(ownerID uint) ([]*model.Computer, error)
	FindById(computerID uint) (*model.Computer, error)
	FindByName(name string, ownerID uint) (*model.Computer, error)

	Create(cm *model.Computer, ownerID uint) error
	Delete(computerID uint) error
	Update(cm *model.Computer) error

	BeginTx() ComputerRepositoryTx
	ContinueTx(mr TransactionContextReader) ComputerRepositoryTx
	TransactionContextReader
}

type GroupComputerRepository interface {
	FindById(groupID uint, computerID uint) (*model.GroupComputer, error)
	FindAllOfGroup(groupID uint) ([]*model.GroupComputer, error)
	FindAllOfComputer(computerID uint) ([]*model.GroupComputer, error)
	FindAllOfGroupAndComputers(groupID uint, computerIDS ...uint) ([]*model.GroupComputer, error) // TODO Testing

	Create(cm *model.GroupComputer) error
	Delete(groupID uint, computerID uint) error
	Update(cm *model.GroupComputer) error

	Exists(groupID uint, computerID uint) (bool, error)

	BeginTx() GroupComputerRepositoryTx
	ContinueTx(mr TransactionContextReader) GroupComputerRepositoryTx
	TransactionContextReader
}

type ComputerBackupRepository interface {

	FindById(groupComputerID uint, backupID uint) (*model.ComputerBackup, error)
	FindAllOfGroupComputer(groupComputerID uint) ([]*model.ComputerBackup, error)
	FindAllOfBackup(backupID uint) ([]*model.ComputerBackup, error)

	Create(cbm *model.ComputerBackup) error
	Delete(groupComputerID uint, backupID uint) error

	Exists(groupComputerID uint, backupID uint) (bool, error)

	BeginTx() ComputerBackupRepositoryTx
	ContinueTx(mr TransactionContextReader) ComputerBackupRepositoryTx
	TransactionContextReader
}

type ComputerBackupRepositoryTx interface {
	ComputerBackupRepository
	Transaction
}

type GroupComputerRepositoryTx interface {
	GroupComputerRepository
	Transaction
}

type BackupRepositoryTx interface {
	BackupRepository
	Transaction
}

type ComputerRepositoryTx interface {
	ComputerRepository
	Transaction
}

type UserRepositoryTx interface {
	UserRepository
	Transaction
}

type GroupRepositoryTx interface {
	GroupRepository
	Transaction
}

type InviteRepositoryTx interface {
	InviteRepository
	Transaction
}
