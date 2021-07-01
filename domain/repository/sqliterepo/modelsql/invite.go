package modelsql

import "time"

// User denoted with UserID is invited to the group denoted with GroupID.
type Invite struct {
	UserID  uint `gorm:"primaryKey"`
	GroupID uint `gorm:"primaryKey"`

	User  User  `gorm:"foreignKey:UserID"`
	Group Group `gorm:"foreignKey:GroupID"`

	CreatedAt time.Time
}
