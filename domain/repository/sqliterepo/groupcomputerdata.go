package sqliterepo

import (
	"errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"mnimidamonbackend/domain/model"
	"mnimidamonbackend/domain/repository"
	. "mnimidamonbackend/domain/repository/sqliterepo/modelsql"
)

func NewComputerDataRepository(db *gorm.DB) repository.GroupComputerRepository {
	return groupComputerData {
		DB: db,
	}
}

type groupComputerData struct {
	*gorm.DB
}

func (gcd groupComputerData) FindById(groupID uint, computerID uint) (*model.GroupComputer, error) {
	var c GroupComputer

	result :=
		gcd.Model(&GroupComputer{}).
			Where("group_id = ? AND computer_id = ?", groupID, computerID).
			First(&c)

	if err := result.Error; err != nil {
		return nil, toBusinessLogicError(err)
	}

	cm := c.NewBusinessModel()
	return cm, nil
}

func (gcd groupComputerData) FindAllOfGroup(groupID uint) ([]*model.GroupComputer, error) {
	var computers []GroupComputer

	result :=
		gcd.Where("group_id = ?", groupID).
			Find(&computers)

	if result.Error != nil {
		return nil, toBusinessLogicError(result.Error)
	}

	var mGComputers []*model.GroupComputer
	for _, c := range computers {
		cm := c.NewBusinessModel()
		mGComputers = append(mGComputers, cm)
	}

	return mGComputers, nil
}

func (gcd groupComputerData) FindAllOfComputer(computerID uint) ([]*model.GroupComputer, error) {
	var computers []GroupComputer

	result :=
		gcd.Where("computer_id = ?", computerID).
			Find(&computers)

	if result.Error != nil {
		return nil, toBusinessLogicError(result.Error)
	}

	var mGComputers []*model.GroupComputer
	for _, c := range computers {
		cm := c.NewBusinessModel()
		mGComputers = append(mGComputers, cm)
	}

	return mGComputers, nil
}

func (gcd groupComputerData) Create(cm *model.GroupComputer) error {
	c := NewGroupComputerFromBusinessModel(cm)

	if exists, _ := gcd.Exists(c.GroupID, c.ComputerID); exists {
		return repository.ErrAlreadyExists
	}

	result :=
		gcd.Omit("id").
			Create(c)

	if result.Error != nil {
		return toBusinessLogicError(result.Error)
	}

	c.CopyToBusinessModel(cm)
	return nil
}

func (gcd groupComputerData) Delete(groupID uint, computerID uint) error {
	result :=
		gcd.DB.
			Where("group_id = ? AND computer_id = ?", groupID, computerID).
			Delete(&GroupComputer{})

	if result.Error != nil {
		return toBusinessLogicError(result.Error)
	}

	return nil
}

func (gcd groupComputerData) Update(cm *model.GroupComputer) error {
	c := NewGroupComputerFromBusinessModel(cm)

	if cm.StorageSize < 0 {
		return repository.ErrInvalidUpdateViolation
	}

	result :=
		gcd.Model(c).
			Omit("id", "group_id", "computer_id", clause.Associations).
			Updates(c).
			Select("*").
			First(c)

	if err := result.Error; err != nil {
		return toBusinessLogicError(err)
	}

	c.CopyToBusinessModel(cm)
	return nil
}

func (gcd groupComputerData) Exists(groupID uint, computerID uint) (bool, error) {
	_, err := gcd.FindById(computerID, groupID)

	if err != nil {
		if  errors.Is(repository.ErrNotFound, err) {
			return false, nil
		}
		return false, toBusinessLogicError(err)
	}

	return true, nil
}

type groupComputerDataTx struct {
	groupComputerData
}

func (gcdtx groupComputerDataTx) Rollback() error {
	return gcdtx.groupComputerData.DB.Rollback().Error
}

func (gcdtx groupComputerDataTx) Commit() error {
	return gcdtx.groupComputerData.DB.Commit().Error
}

func (gcd groupComputerData) BeginTx() repository.GroupComputerRepositoryTx {
	return groupComputerDataTx{
		groupComputerData{
			DB: gcd.Begin(),
		},
	}
}

