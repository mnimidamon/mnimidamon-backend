package modelsql

import (
	"mnimidamonbackend/domain/model"
	"time"
)

// User denoted with UserID is invited to the group denoted with GroupID.
type Invite struct {
	UserID  uint `gorm:"primaryKey"`
	GroupID uint `gorm:"primaryKey"`

	User  *User  `gorm:"foreignKey:UserID"`
	Group *Group `gorm:"foreignKey:GroupID"`

	CreatedAt time.Time
}


func NewInviteFromBusinessModel(im *model.Invite) *Invite {
	if im == nil {
		return nil
	}

	return &Invite{
		UserID:    im.UserID,
		GroupID:   im.GroupID,
		User:      NewUserFromBusinessModel(im.User),
		Group:     NewGroupFromBusinessModel(im.Group),
		CreatedAt: time.Time{},
	}
}

func (i *Invite) NewBusinessModel() *model.Invite {
	if i == nil {
		return nil
	}

	im := new(model.Invite)
	i.CopyToBusinessModel(im)
	return im
}


func (i *Invite) CopyToBusinessModel(im  *model.Invite) {
	if i == nil {
		im = nil
		return
	}

	im.GroupID = i.GroupID
	im.UserID = i.UserID

	im.Group = i.Group.NewBusinessModel()
	im.User = i.User.NewBusinessModel()
}