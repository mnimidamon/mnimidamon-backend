package modelsql

import (
	"mnimidamonbackend/domain/model"
	"time"
)

type User struct {
	Entity

	Username     string `gorm:"unique; index; size:15"`
	PasswordHash string

	Computers []Computer `gorm:"foreignKey:OwnerID"`
}

func NewUserFromBusinessModel(um *model.User) *User {
	if um == nil {
		return nil
	}

	return &User{
		Entity:       Entity{
			ID:        um.ID,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
		},
		Username:     um.Username,
		PasswordHash: um.PasswordHash,
		Computers:    nil,
	}
}

func (u *User) NewBusinessModel()  *model.User {
	if u == nil {
		return nil
	}

	um := new(model.User)
	u.CopyToBusinessModel(um)
	return um
}

func (u *User) CopyToBusinessModel(um *model.User)  {
	if u == nil {
		um = nil
		return
	}

	um.PasswordHash = u.PasswordHash
	um.Username = u.Username
	um.ID = u.ID
}
