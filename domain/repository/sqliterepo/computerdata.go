package sqliterepo

import (
	"gorm.io/gorm"
	"mnimidamonbackend/domain/model"
	"mnimidamonbackend/domain/repository"
	. "mnimidamonbackend/domain/repository/sqliterepo/modelsql"
)

func NewComputerRepository(db *gorm.DB) repository.ComputerRepository {
	return computerData{db}
}

// computerData store for SQLite implementation.
type computerData struct {
	*gorm.DB
}

func (cd computerData) Delete(computerID uint) error {
	result :=
		cd.DB.Delete(&Computer{}, computerID)

	if err := result.Error; err != nil {
		return toBusinessLogicError(err)
	}

	return nil
}

func (cd computerData) Update(cm *model.Computer) error {
	c := NewComputerFromBusinessModel(cm)

	result :=
		cd.Model(c).
			Select("name").
			Updates(c)

	if err := result.Error; err != nil {
		return toBusinessLogicError(err)
	}

	c.CopyToBusinessModel(cm)
	return nil
}

func (cd computerData) FindAll(ownerID uint) ([]*model.Computer, error) {
	var computers []Computer

	result :=
		cd.Where("owner_id = ?", ownerID).
			Find(&computers)

	if result.Error != nil {
		return nil, toBusinessLogicError(result.Error)
	}

	var mComputers []*model.Computer
	for _, c := range computers {
		mc := c.NewBusinessModel()
		mComputers = append(mComputers, mc)
	}

	return mComputers, nil
}

func (cd computerData) FindById(computerID uint) (*model.Computer, error) {
	var computer Computer

	result :=
		cd.First(&computer, computerID)

	if err := result.Error; err != nil {
		return nil, toBusinessLogicError(err)
	}

	cm := computer.NewBusinessModel()

	return cm, nil
}

func (cd computerData) FindByName(name string, ownerID uint) (*model.Computer, error) {
	var computer Computer

	result :=
		cd.Where("name = ? AND owner_id = ?", name, ownerID).
			Find(&computer)

	if result.Error != nil {
		return nil, toBusinessLogicError(result.Error)
	}

	cm := computer.NewBusinessModel()
	return cm, nil
}

func (cd computerData) Create(cm *model.Computer) error {
	c := NewComputerFromBusinessModel(cm)

	result :=
		cd.Omit("id").
			Create(c)

	if err := result.Error; err != nil {
		return toBusinessLogicError(err)
	}

	c.CopyToBusinessModel(cm)
	return nil
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

func (cd computerData) BeginTx() repository.ComputerRepositoryTx {
	return &computerDataTx{
		computerData{
			DB: cd.Begin(),
		},
	}
}
