package sqliterepo

import (
	. "gorm.io/gorm"
	"mnimidamonbackend/domain/repository"
)

func NewUserRepository(db *DB) repository.UserRepository {
	return userData{db}
}

// userData store for SQLite database.
type userData struct {
	 *DB
}

// Transaction support.
type userDataTx struct {
	userData
}

func (u userDataTx) Rollback() error {
	return u.userData.Rollback().Error
}

func (u userDataTx) Commit() error {
	return u.userData.Commit().Error
}

func (ud userData) BeginTx() repository.UserRepositoryTx {
	return userDataTx{
		userData{
			DB: ud.Begin(),
		},
	}
}
