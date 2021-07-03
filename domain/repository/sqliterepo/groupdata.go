package sqliterepo

import (
	. "gorm.io/gorm"
	"mnimidamonbackend/domain/model"
	"mnimidamonbackend/domain/repository"
)

func NewGroupRepository(db *DB) repository.GroupRepository {
	return groupData{db}
}

// groupData store for SQLite database.
type groupData struct {
	*DB
}

func (gd groupData) FindAll() ([]*model.Group, error) {
	panic("implement me")
}

func (gd groupData) FindById(groupID int) (*model.Group, error) {
	panic("implement me")
}

func (gd groupData) FindByName(username string) (*model.Group, error) {
	panic("implement me")
}

func (gd groupData) Create(gm *model.Group) error {
	panic("implement me")
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