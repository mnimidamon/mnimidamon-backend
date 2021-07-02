package repository

import "mnimidamonbackend/domain/model"

type UserRepository interface {
	BeginTx() UserRepositoryTx

	// TODO: Functionalities
	FindAll() ([]*model.User, error)
	FindByUsername(username string) (*model.User, error)
	Save(*model.User) error
}

type UserRepositoryTx interface {
	UserRepository
	Transaction
}

type GroupRepository interface {
	BeginTx() GroupRepositoryTx

	// TODO: Functionalities
}

type GroupRepositoryTx interface {
	GroupRepository
	Transaction
}

type ComputerRepository interface {
	BeginTx() ComputerServiceTx

	// TODO: Functionalities
}

type ComputerServiceTx interface {
	ComputerRepository
	Transaction
}

type BackupRepository interface {
	BeginTx() BackupRepositoryTx

	// TODO: Functionalities
}

type BackupRepositoryTx interface {
	BackupRepository
	Transaction
}
