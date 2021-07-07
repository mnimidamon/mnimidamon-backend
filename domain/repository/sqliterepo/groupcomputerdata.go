package sqliterepo

import (
	"gorm.io/gorm"
	"mnimidamonbackend/domain/model"
	"mnimidamonbackend/domain/repository"
)

func NewComputerDataRepository(db *gorm.DB) repository.GroupComputerRepository {
	return groupComputerData {
		DB: db,
	}
}

type groupComputerData struct {
	*gorm.DB
}

func (g groupComputerData) BeginTx() repository.GroupComputerRepositoryTx {
	panic("implement me")
}

func (g groupComputerData) FindOwner(groupID uint, computerID uint) (*model.User, error) {
	panic("implement me")
}

func (g groupComputerData) FindById(groupID uint, computerID uint) (*model.GroupComputer, error) {
	panic("implement me")
}

func (g groupComputerData) FindAllOfGroup(groupID uint) ([]*model.GroupComputer, error) {
	panic("implement me")
}

func (g groupComputerData) FindAllOfComputer(computerID uint) ([]*model.GroupComputer, error) {
	panic("implement me")
}

func (g groupComputerData) Create(cm *model.GroupComputer) (*model.GroupComputer, error) {
	panic("implement me")
}

func (g groupComputerData) Delete(groupID uint, computerID uint) error {
	panic("implement me")
}

func (g groupComputerData) Update(cm *model.GroupComputer) (*model.GroupComputer, error) {
	panic("implement me")
}

func (g groupComputerData) Exists(userID uint, groupID uint) (bool, error) {
	panic("implement me")
}

type groupComputerDataTx struct {
	inviteData
}

func (g groupComputerDataTx) BeginTx() repository.GroupComputerRepositoryTx {
	panic("implement me")
}

func (g groupComputerDataTx) FindOwner(groupID uint, computerID uint) (*model.User, error) {
	panic("implement me")
}

func (g groupComputerDataTx) FindById(groupID uint, computerID uint) (*model.GroupComputer, error) {
	panic("implement me")
}

func (g groupComputerDataTx) FindAllOfGroup(groupID uint) ([]*model.GroupComputer, error) {
	panic("implement me")
}

func (g groupComputerDataTx) FindAllOfComputer(computerID uint) ([]*model.GroupComputer, error) {
	panic("implement me")
}

func (g groupComputerDataTx) Create(cm *model.GroupComputer) (*model.GroupComputer, error) {
	panic("implement me")
}

func (g groupComputerDataTx) Update(cm *model.GroupComputer) (*model.GroupComputer, error) {
	panic("implement me")
}

func (g groupComputerDataTx) Rollback() error {
	panic("implement me")
}

func (g groupComputerDataTx) Commit() error {
	panic("implement me")
}
