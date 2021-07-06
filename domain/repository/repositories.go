package repository

import "mnimidamonbackend/domain/model"

// Repositories are used for reading and writing models.
// Checkers are used for checking common conditions that other repositories need for constraints checking.
// Repositories Tx objects are used for when a transaction is initiated.

type UserRepository interface {
	BeginTx() UserRepositoryTx

	FindAll() ([]*model.User, error)
	FindById(userID uint) (*model.User, error)
	FindByUsername(username string) (*model.User, error)

	Create(um *model.User) error
	Delete(userID uint) error
	Update(um *model.User) error

	UserRepositoryChecker
	// TODO: Functionalities
	// 		- this is completed?
}

type UserRepositoryChecker interface {
	Exists(userID uint) (bool, error)
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
	Delete(groupID uint) error
	Update(gm *model.Group) error

	AddMember(userID uint, groupID uint) (*model.Group, error)

	GroupRepositoryChecker
	// TODO: Functionalities:
	//		- get members
}

type GroupRepositoryChecker interface {
	IsMemberOf(userID uint, groupID uint) (bool, error)
	Exists(groupID uint) (bool, error)
}

type GroupRepositoryTx interface {
	GroupRepository
	Transaction
}

type ComputerRepository interface {
	BeginTx() ComputerRepositoryTx

	FindAll() ([]*model.Computer, error)
	FindById(computerID uint) (*model.Computer, error)
	FindByName(name string) (*model.Computer, error)

	Create(cm *model.Computer) error
	Delete(computerID uint) error
	Update(cm *model.Computer) error

	ComputerRepositoryChecker
	// TODO: Functionalities
	// 		- have to think about this little bit
}

type ComputerRepositoryChecker interface {

}

type ComputerRepositoryTx interface {
	ComputerRepository
	Transaction
}

type BackupRepository interface {
	BeginTx() BackupRepositoryTx

	FindAll(groupID uint) ([]*model.Backup, error)
	FindById(backupID uint) (*model.Backup, error)

	Create(bm *model.Backup) error
	Delete(backupID uint) error
	Update(bm *model.Backup) error

	BackupRepositoryChecker
	// TODO: Functionalities
	//		- field updating
	//		- deleting works?
}

type BackupRepositoryChecker interface {

}

type BackupRepositoryTx interface {
	BackupRepository
	Transaction
}

type InviteRepository interface {
	BeginTx() InviteRepositoryTx

	InviteRepositoryChecker
	// TODO: Functionalities:
	//		- inviting
	//		- declining invites
	//		- accepting invites
}

type InviteRepositoryChecker interface {

}

type InviteRepositoryTx interface {
	InviteRepository
	Transaction
}
