package repository

import "mnimidamonbackend/domain/model"

type UserRepository interface {
	BeginTx() UserRepositoryTx

	FindAll() ([]*model.User, error)
	FindById(userID int) (*model.User, error)
	FindByUsername(username string) (*model.User, error)

	Create(um *model.User) error

	// TODO: Functionalities
}

type UserRepositoryTx interface {
	UserRepository
	Transaction
}

type GroupRepository interface {
	BeginTx() GroupRepositoryTx

	FindAll() ([]*model.Group, error)
	FindById(groupID int) (*model.Group, error)
	FindByName(username string) (*model.Group, error)

	Create(gm *model.Group) error

	// TODO: Functionalities
}

type GroupRepositoryTx interface {
	GroupRepository
	Transaction
}

type ComputerRepository interface {
	BeginTx() ComputerServiceTx


	FindAll() ([]*model.Computer, error)
	FindById(computerID int) (*model.Computer, error)
	FindByName(name string) (*model.Computer, error)

	Create(cm *model.Computer) error

	// TODO: Functionalities
}

type ComputerServiceTx interface {
	ComputerRepository
	Transaction
}

type BackupRepository interface {
	BeginTx() BackupRepositoryTx

	FindAll() ([]*model.Backup, error)
	FindById(backupID int) (*model.Backup, error)

	Create(bm *model.Backup) error


	// TODO: Functionalities
}

type BackupRepositoryTx interface {
	BackupRepository
	Transaction
}
