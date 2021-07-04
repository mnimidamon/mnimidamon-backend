package sqliterepo

import (
	"fmt"
	. "gorm.io/gorm"
	"mnimidamonbackend/domain/model"
	"mnimidamonbackend/domain/repository"
	. "mnimidamonbackend/domain/repository/sqliterepo/modelsql"
)

func NewGroupRepository(db *DB) repository.GroupRepository {
	return groupData{db}
}

// groupData store for SQLite database.
type groupData struct {
	*DB
}

func (gd groupData) FindAll() ([]*model.Group, error) {
	var groups []Group

	result := gd.Find(&groups)

	if err := result.Error; err != nil {
		return nil, toBusinessLogicError(err)
	}

	var mGroups []*model.Group

	for _, g := range groups {
		mg := g.NewBusinessModel()
		mGroups = append(mGroups, mg)
	}

	return mGroups, nil
}

func (gd groupData) FindById(groupID int) (*model.Group, error) {
	var group Group

	result :=
		gd.First(&group, groupID)

	if err := result.Error; err != nil {
		return nil, toBusinessLogicError(err)
	}

	gm := group.NewBusinessModel()

	return gm, nil
}

func (gd groupData) FindByName(name string) (*model.Group, error) {
	var group Group

	result :=
		gd.Where("name LIKE ?", fmt.Sprintf("%s%%", name)).
		First(&group)

	if err := result.Error; err != nil {
		return nil, toBusinessLogicError(err)
	}

	gm := group.NewBusinessModel()

	return gm, nil
}

func (gd groupData) Create(gm *model.Group) error {
	g := NewGroupFromBusinessModel(gm)

	if result := gd.Omit("id").Create(&g); result.Error != nil {
		return toBusinessLogicError(result.Error)
	}

	g.CopyToBusinessModel(gm)
	return nil
}

// Transaction support.
type groupDataTx struct {
	groupData
}

func (gdtx groupDataTx) Rollback() error{
	return gdtx.groupData.Rollback().Error
}

func (gdtx groupDataTx) Commit() error {
	return gdtx.groupData.Commit().Error
}

func (gd groupData) BeginTx() repository.GroupRepositoryTx {
	return groupDataTx{
		groupData{
			DB: gd.Begin(),
		},
	}
}