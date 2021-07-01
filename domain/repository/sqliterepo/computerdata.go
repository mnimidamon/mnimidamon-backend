package sqliterepo

import (
	. "gorm.io/gorm"
	"mnimidamonbackend/domain/repository"
)

func NewComputerRepository(db *DB) repository.ComputerRepository {
	return computerData{db}
}

// computerData store for SQLite implementation.
type computerData struct {
	*DB
}

// Transaction support.
type computerDataTx struct {
	computerData

}

func (c computerDataTx) Rollback() error {
	return c.computerData.Rollback().Error
}

func (c computerDataTx) Commit() error {
	return c.computerData.Commit().Error
}

func (cd computerData) BeginTx() repository.ComputerServiceTx {
	return &computerDataTx{
		computerData{
			DB: cd.Begin(),
		},
	}
}
