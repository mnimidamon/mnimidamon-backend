package sqliterepo

import (
	"errors"
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

func (gd groupData) ContinueTx(mr repository.TransactionContextReader) repository.GroupRepositoryTx {
	meta := mr.GetContext()

	if dbtx, isDB := meta.(*gorm.DB); isDB {
		return groupDataTx{
			groupData{
				DB: dbtx,
			},
		}
	}

	return gd.BeginTx()
}

func (gd groupData) GetContext() interface{} {
	return gd.DB
}

func (gd groupData) FindAllOfUser(userID uint) ([]*model.Group, error) {
	var groups []Group

	result :=
		gd.Where("id IN (?)",
			gd.Table("group_members").
			Where("user_id = ?", userID).
			Select("group_id"),
		).Find(&groups)

	if result.Error != nil {
		return nil, toRepositoryError(result.Error)
	}

	var mGroups []*model.Group

	for _, g := range groups {
		mg := g.NewBusinessModel()
		mGroups = append(mGroups, mg)
	}

	return mGroups, nil
}

func (gd groupData) FindAllMembers(groupID uint) ([]*model.User, error) {
	var group Group

	result :=
		gd.Where("id = ?", groupID).
			Preload("GroupMembers").
			Find(&group)

	if result.Error != nil {
		return nil, toRepositoryError(result.Error)
	}

	var mUsers []*model.User
	for _, u := range group.GroupMembers {
		mb := u.NewBusinessModel()
		mUsers = append(mUsers, mb)
	}

	return mUsers, nil
}

func (gd groupData) Exists(groupID uint) (bool, error) {
	_, err := gd.FindById(groupID)

	if err != nil  {
		if  errors.Is(repository.ErrNotFound, err) {
			return false, nil
		}
		return false, toRepositoryError(err)
	}

	return true, nil
}

func (gd groupData) IsMemberOf(userID uint, groupID uint) (bool, error) {
	count := new(int64)

	result := gd.Table("group_members").
		Where("user_id = ? AND group_ID = ?", userID, groupID).
		Count(count)

	if result.Error != nil {
		return false, toRepositoryError(result.Error)
	}

	if *count > 0 {
		return true, nil
	}

	return false, nil
}

func (gd groupData) AddMember(userID uint, groupID uint) (*model.Group, error) {

	if isMember, _ := gd.IsMemberOf(userID, groupID); isMember {
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
		return nil, toRepositoryError(err)
	}

	gm := group.NewBusinessModel()
	return gm, nil
}

func (gd groupData) Delete(groupID uint) error {

	g, err := gd.FindById(groupID)

	if err != nil {
		return repository.ErrNotFound
	}

	result := gd.DB.
		Select("GroupMembers", "Invites").
		Delete(g)

	if err := result.Error; err != nil {
		return toRepositoryError(err)
	}

	return nil
}

func (gd groupData) Update(gm *model.Group) error {
	g := NewGroupFromBusinessModel(gm)

	result :=
		gd.Model(g).
			Omit("id", clause.Associations).
			Updates(g).
			Select("*").
			First(g)

	if err := result.Error; err != nil {
		return toRepositoryError(err)
	}

	g.CopyToBusinessModel(gm)
	return nil
}

func (gd groupData) FindAll() ([]*model.Group, error) {
	var groups []Group

	result := gd.Find(&groups)

	if err := result.Error; err != nil {
		return nil, toRepositoryError(err)
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
		return nil, toRepositoryError(err)
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
		return nil, toRepositoryError(err)
	}

	gm := group.NewBusinessModel()

	return gm, nil
}

func (gd groupData) Create(gm *model.Group) error {
	g := NewGroupFromBusinessModel(gm)

	if result := gd.Omit("id").Create(g); result.Error != nil {
		return toRepositoryError(result.Error)
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
