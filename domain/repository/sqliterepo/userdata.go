package sqliterepo

import (
	"fmt"
	. "gorm.io/gorm"
	"mnimidamonbackend/domain/model"
	"mnimidamonbackend/domain/repository"
	. "mnimidamonbackend/domain/repository/sqliterepo/modelsql"
)

func NewUserRepository(db *DB) repository.UserRepository {
	return userData{db}
}

// userData store for SQLite database.
type userData struct {
	 *DB
}

func (ud userData) FindAll() ([]*model.User, error) {
	var users []User

	result := ud.Find(&users)

	if err := result.Error; err != nil {
		return nil, toBusinessLogicError(err)
	}

	var mUsers []*model.User

	for _, u := range users {
		mu := new(model.User)

		mu.ID = u.ID
		mu.Username = u.Username
		mu.PasswordHash = u.PasswordHash

		mUsers = append(mUsers, mu)
	}

	return mUsers, nil
}

func (ud userData) FindByUsername(username string) (*model.User, error) {
	var user User

	result :=
		ud.Where("username LIKE ?", fmt.Sprintf("%%%s%%", username)).
		First(&user)

	if err := result.Error; err != nil {
		return nil, toBusinessLogicError(err)
	}

	um := new(model.User)
	um.PasswordHash = user.PasswordHash
	um.Username = user.Username
	um.ID = user.ID

	return um, nil
}

func (ud userData) Save(user *model.User) error {
	var u User

	u.Username = user.Username
	u.PasswordHash = user.PasswordHash

	if result := ud.Create(&u); result.Error != nil {
		return toBusinessLogicError(result.Error)
	}

	user.ID = u.ID

	return nil
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
