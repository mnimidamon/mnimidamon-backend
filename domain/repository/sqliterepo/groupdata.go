package sqliterepo

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"mnimidamonbackend/domain/model"
	"mnimidamonbackend/domain/repository"
	. "mnimidamonbackend/domain/repository/sqliterepo/modelsql"
)

func NewGroupRepository(db *gorm.DB) repository.GroupRepository {
	return groupData{db}
}

// groupData store for SQLite database.
type groupData struct {
	*gorm.DB
}

func (gd groupData) IsMemberOf(userID uint, groupID uint) bool {
	var user User
	var group Group

	group.ID = groupID
	user.ID = userID

	count :=
		gd.Find(&group, groupID).
			Where("user_id = ?", userID).
			Association("GroupMembers").
			Count()

	if count > 0 {
		return true
	}

	return false
}

func (gd groupData) AddMember(userID uint, groupID uint) (*model.User, error) {

	if gd.IsMemberOf(userID, groupID) {
		return nil, repository.ErrUserAlreadyInGroupViolation
	}

	var user User
	var group Group

	group.ID = groupID
	user.ID = userID



	err :=
		gd.Find(&group, groupID).
			Omit("GroupMembers.*").
			Association("GroupMembers").
			Append(&user)

	if err != nil {
		return nil, toBusinessLogicError(err)
	}

	um := user.NewBusinessModel()
	return um, nil
}

func (gd groupData) Delete(gm *model.Group) error {
	result := gd.DB.Delete(&Group{}, gm.ID)

	if err := result.Error; err != nil {
		return toBusinessLogicError(err)
	}

	return nil
}

func (gd groupData) Update(gm *model.Group) error {
	g := NewGroupFromBusinessModel(gm)

	result :=
		gd.Model(g).
			Omit("id", clause.Associations).
			Updates(g)

	if err := result.Error; err != nil {
		return toBusinessLogicError(err)
	}

	g.CopyToBusinessModel(gm)
	return nil
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

func (gd groupData) FindById(groupID uint) (*model.Group, error) {
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

	if result := gd.Omit("id").Create(g); result.Error != nil {
		return toBusinessLogicError(result.Error)
	}

	g.CopyToBusinessModel(gm)
	return nil
}

// Transaction support.
type groupDataTx struct {
	groupData
}

func (gdtx groupDataTx) Rollback() error {
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
