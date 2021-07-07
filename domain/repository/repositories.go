package repository

import "mnimidamonbackend/domain/model"

// Repositories are used for reading and writing models.
// Repositories Tx objects are used for when a transaction is initiated.

type UserRepository interface {
	BeginTx() UserRepositoryTx

	FindAll() ([]*model.User, error)
	FindById(userID uint) (*model.User, error)
	FindByUsername(username string) (*model.User, error)

	Create(um *model.User) error
	Delete(userID uint) error
	Update(um *model.User) error

	Exists(userID uint) (bool, error)
	// TODO: Functionalities
	// 		- this is completed?
}

type GroupRepository interface {
	BeginTx() GroupRepositoryTx

	FindAll() ([]*model.Group, error)
	FindAllMembers(groupID uint) ([]*model.User, error) // TODO: testing
	FindById(groupID uint) (*model.Group, error)
	FindByName(name string) (*model.Group, error)

	Create(gm *model.Group) error
	Delete(groupID uint) error
	Update(gm *model.Group) error

	AddMember(userID uint, groupID uint) (*model.Group, error)
	IsMemberOf(userID uint, groupID uint) (bool, error)
	Exists(groupID uint) (bool, error)

	// TODO: Functionalities:
	//		- get members
}

type BackupRepository interface {
	BeginTx() BackupRepositoryTx

	FindAll(groupID uint) ([]*model.Backup, error)
	FindById(backupID uint) (*model.Backup, error)

	Create(bm *model.Backup) error
	Delete(backupID uint) error
	Update(bm *model.Backup) error

	// TODO: Functionalities
	//		- field updating
}

type InviteRepository interface {
	BeginTx() InviteRepositoryTx

	Create(im *model.Invite) error
	Delete(userID uint, groupID uint) error

	FindAllOfGroup(groupID uint) ([]*model.Invite, error)
	FindAllOfUser(userID uint) ([]*model.Invite, error)
	FindById(userID uint, groupID uint) (*model.Invite, error)

	Exists(userID uint, groupID uint) (bool, error)
}

type ComputerRepository interface {
	BeginTx() ComputerRepositoryTx

	FindAll(ownerID uint) ([]*model.Computer, error)
	FindById(computerID uint) (*model.Computer, error)
	FindByName(name string, ownerID uint) (*model.Computer, error)

	Create(cm *model.Computer, ownerID uint) error
	Delete(computerID uint) error
	Update(cm *model.Computer) error

	// TODO: Functionalities
	// 		- have to think about this little bit
}

type GroupComputerRepository interface {
	BeginTx() GroupComputerRepositoryTx

	FindById(groupID uint, computerID uint) (*model.GroupComputer, error)
	FindAllOfGroup(groupID uint) ([]*model.GroupComputer, error)
	FindAllOfComputer(computerID uint) ([]*model.GroupComputer, error)

	Create(cm *model.GroupComputer) error
	Delete(groupID uint, computerID uint) error
	Update(cm *model.GroupComputer) error

	Exists(groupID uint, computerID uint) (bool, error)
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
