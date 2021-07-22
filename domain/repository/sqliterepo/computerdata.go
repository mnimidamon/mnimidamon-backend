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

func (cd computerData) ContinueTx(mr repository.TransactionContextReader) repository.ComputerRepositoryTx {
	meta := mr.GetContext()

	if dbtx, isDB := meta.(*gorm.DB); isDB {
		return computerDataTx{
			computerData{
				DB: dbtx,
			},
		}
	}

	return cd.BeginTx()
}

func (cd computerData) GetContext() interface{} {
	return cd.DB
}

func (cd computerData) Delete(computerID uint) error {
	result :=
		cd.DB.Delete(&Computer{}, computerID)

	if err := result.Error; err != nil {
		return toRepositoryError(err)
	}

	return nil
}

func (cd computerData) Update(cm *model.Computer) error {
	c := NewComputerFromBusinessModel(cm)

	result :=
		cd.Model(c).
			Select("name").
			Updates(c).
			Select("*").
			First(c)

	if err := result.Error; err != nil {
		return toRepositoryError(err)
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
		return nil, toRepositoryError(result.Error)
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
		return nil, toRepositoryError(err)
	}

	cm := computer.NewBusinessModel()

	return cm, nil
}

func (cd computerData) FindByName(name string, ownerID uint) (*model.Computer, error) {
	var computer Computer

	result :=
		cd.Where("name = ? AND owner_id = ?", name, ownerID).
			First(&computer)

	if result.Error != nil {
		return nil, toRepositoryError(result.Error)
	}

	cm := computer.NewBusinessModel()
	return cm, nil
}

func (cd computerData) Create(cm *model.Computer, ownerID uint) error {
	c := NewComputerFromBusinessModel(cm)
	c.OwnerID = ownerID

	// Unique names for computers based on owners.
	if _, err := cd.FindByName(cm.Name, ownerID); err == nil {
		return repository.ErrUniqueConstraintViolation
	}

	result :=
		cd.Omit("id").
			Create(c)


	if err := result.Error; err != nil {
		return toRepositoryError(err)
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
