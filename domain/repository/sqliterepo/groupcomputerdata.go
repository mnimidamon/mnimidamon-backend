package sqliterepo

import (
	"errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"mnimidamonbackend/domain/model"
	"mnimidamonbackend/domain/repository"
	. "mnimidamonbackend/domain/repository/sqliterepo/modelsql"
)

func NewGroupComputerRepository(db *gorm.DB) repository.GroupComputerRepository {
	return groupComputerData{
		DB: db,
	}
}

type groupComputerData struct {
	*gorm.DB
}

func (gcd groupComputerData) ContinueTx(mr repository.TransactionContextReader) repository.GroupComputerRepositoryTx {
	meta := mr.GetContext()

	if dbtx, isDB := meta.(*gorm.DB); isDB {
		return groupComputerDataTx{
			groupComputerData{
				DB: dbtx,
			},
		}
	}

	return gcd.BeginTx()
}

func (gcd groupComputerData) GetContext() interface{} {
	return gcd.DB
}

func (gcd groupComputerData) FindAllOfGroupAndComputers(groupID uint, computerIDS ...uint) ([]*model.GroupComputer, error) {
	var groupComputers []*GroupComputer

	result :=
		gcd.Model(&GroupComputer{}).
			Where("group_id = ? AND computer_id IN (?)", groupID, computerIDS).
			Find(&groupComputers)

	if err := result.Error; err != nil {
		return nil, toRepositoryError(err)
	}

	var mGComputers []*model.GroupComputer
	for _, c := range groupComputers {
		cm := c.NewBusinessModel()
		mGComputers = append(mGComputers, cm)
	}
	return mGComputers, nil
}

func (gcd groupComputerData) FindById(groupID uint, computerID uint) (*model.GroupComputer, error) {
	var c GroupComputer

	result :=
		gcd.Model(&GroupComputer{}).
			Where("group_id = ? AND computer_id = ?", groupID, computerID).
			First(&c)

	if err := result.Error; err != nil {
		return nil, toRepositoryError(err)
	}

	cm := c.NewBusinessModel()
	return cm, nil
}

func (gcd groupComputerData) FindAllOfGroup(groupID uint) ([]*model.GroupComputer, error) {
	var computers []GroupComputer

	result :=
		gcd.Where("group_id = ?", groupID).
			Preload("Computer").
			Find(&computers)

	if result.Error != nil {
		return nil, toRepositoryError(result.Error)
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
			Preload("Group").
			Find(&computers)

	if result.Error != nil {
		return nil, toRepositoryError(result.Error)
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
		return toRepositoryError(result.Error)
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
		return toRepositoryError(result.Error)
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
		return toRepositoryError(err)
	}

	c.CopyToBusinessModel(cm)
	return nil
}

func (gcd groupComputerData) Exists(groupID uint, computerID uint) (bool, error) {
	_, err := gcd.FindById(computerID, groupID)

	if err != nil {
		if errors.Is(repository.ErrNotFound, err) {
			return false, nil
		}
		return false, toRepositoryError(err)
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
