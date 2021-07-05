package repository

import "mnimidamonbackend/domain/model"

type UserRepository interface {
	BeginTx() UserRepositoryTx

	FindAll() ([]*model.User, error)
	FindById(userID uint) (*model.User, error)
	FindByUsername(username string) (*model.User, error)

	Create(um *model.User) error
	Delete(um *model.User) error
	Update(um *model.User) error

	// TODO: Functionalities
	// 		- this is completed?
}

type UserRepositoryTx interface {
	UserRepository
	Transaction
}

type GroupRepository interface {
	BeginTx() GroupRepositoryTx

	FindAll() ([]*model.Group, error)
	FindById(groupID uint) (*model.Group, error)
	FindByName(name string) (*model.Group, error)

	Create(gm *model.Group) error
	Delete(gm *model.Group) error
	Update(gm *model.Group) error

	AddMember(userID uint, groupID uint) (*model.Group, error)
	IsMemberOf(userID uint, groupID uint) (bool, error)

	// TODO: Functionalities:
	//		- inviting
	//		- declining invites
	//		- accepting invites
	//		- get members
}

type GroupRepositoryTx interface {
	GroupRepository
	Transaction
}

type ComputerRepository interface {
	BeginTx() ComputerServiceTx

	FindAll() ([]*model.Computer, error)
	FindById(computerID uint) (*model.Computer, error)
	FindByName(name string) (*model.Computer, error)

	Create(cm *model.Computer) error
	Delete(cm *model.Computer) error
	Update(cm *model.Computer) error

	// TODO: Functionalities
	// 		- have to think about this little bit
}

type ComputerServiceTx interface {
	ComputerRepository
	Transaction
}

type BackupRepository interface {
	BeginTx() BackupRepositoryTx

	FindAll(groupID uint) ([]*model.Backup, error)
	FindById(backupID uint) (*model.Backup, error)

	Create(bm *model.Backup) error
	Delete(bm *model.Backup) error
	Update(bm *model.Backup) error

	// TODO: Functionalities
	//		- field updating
	//		- deleting works?
}

type BackupRepositoryTx interface {
	BackupRepository
	Transaction
}
