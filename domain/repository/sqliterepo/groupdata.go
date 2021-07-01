package sqliterepo

import (
	. "gorm.io/gorm"
	"mnimidamonbackend/domain/repository"
)

func NewGroupRepository(db *DB) repository.GroupRepository {
	return groupData{db}
}

// groupData store for SQLite database.
type groupData struct {
	*DB
}

type groupDataTx struct {
	groupData
}

func (g groupDataTx) Rollback() error{
	return g.groupData.Rollback().Error
}

func (g groupDataTx) Commit() error {
	return g.groupData.Commit().Error
}

func (gd groupData) BeginTx() repository.GroupRepositoryTx {
	return groupDataTx{
		groupData{
			DB: gd.Begin(),
		},
	}
}