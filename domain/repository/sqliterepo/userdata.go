package sqliterepo

import (
	"fmt"
	"gorm.io/gorm"
	"mnimidamonbackend/domain/model"
	"mnimidamonbackend/domain/repository"
	. "mnimidamonbackend/domain/repository/sqliterepo/modelsql"
)

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return userData{db}
}

// userData store for SQLite database.
type userData struct {
	 *gorm.DB
}

func (ud userData) Delete(um *model.User) error {
	panic("implement me")
}

func (ud userData) Update(um *model.User) error {
	panic("implement me")
}

func (ud userData) FindById(userID int) (*model.User, error) {
	var user User

	result :=
		ud.First(&user, userID)

	if err := result.Error; err != nil {
		return nil, toBusinessLogicError(err)
	}

	um := user.NewBusinessModel()

	return um, nil
}

func (ud userData) FindAll() ([]*model.User, error) {
	var users []User

	result := ud.Find(&users)

	if err := result.Error; err != nil {
		return nil, toBusinessLogicError(err)
	}

	var mUsers []*model.User

	for _, u := range users {
		mu := u.NewBusinessModel()
		mUsers = append(mUsers, mu)
	}

	return mUsers, nil
}

func (ud userData) FindByUsername(username string) (*model.User, error) {
	var user User

	result :=
		ud.Where("username LIKE ?", fmt.Sprintf("%s%%", username)).
		First(&user)

	if err := result.Error; err != nil {
		return nil, toBusinessLogicError(err)
	}

	um := user.NewBusinessModel()

	return um, nil
}

func (ud userData) Create(um *model.User) error {
	u := NewUserFromBusinessModel(um)


	if result := ud.Omit("id").Create(&u); result.Error != nil {
		return toBusinessLogicError(result.Error)
	}

	u.CopyToBusinessModel(um)
	return nil
}

// Transaction support.
type userDataTx struct {
	userData
}

func (udtx userDataTx) Rollback() error {
	return udtx.userData.Rollback().Error
}

func (udtx userDataTx) Commit() error {
	return udtx.userData.Commit().Error
}

func (ud userData) BeginTx() repository.UserRepositoryTx {
	return userDataTx{
		userData{
			DB: ud.Begin(),
		},
	}
}
