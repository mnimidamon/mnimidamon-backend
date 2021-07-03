package sqliterepo

import (
	. "gorm.io/gorm"
	"mnimidamonbackend/domain/model"
	"mnimidamonbackend/domain/repository"
)

func NewComputerRepository(db *DB) repository.ComputerRepository {
	return computerData{db}
}

// computerData store for SQLite implementation.
type computerData struct {
	*DB
}

func (cd computerData) FindAll() ([]*model.Computer, error) {
	panic("implement me")
}

func (cd computerData) FindById(computerID int) (*model.Computer, error) {
	panic("implement me")
}

func (cd computerData) FindByName(name string) (*model.Computer, error) {
	panic("implement me")
}

func (cd computerData) Create(cm *model.Computer) error {
	panic("implement me")
}

// Transaction support.
type computerDataTx struct {
	computerData

}

func (cdtx computerDataTx) Rollback() error {
	return cdtx.computerData.Rollback().Error
}

func (cdtx computerDataTx) Commit() error {
	return cdtx.computerData.Commit().Error
}

func (cd computerData) BeginTx() repository.ComputerServiceTx {
	return &computerDataTx{
		computerData{
			DB: cd.Begin(),
		},
	}
}
