package model

type Group struct {
	Entity

	Name string

	GroupMembers []User   `gorm:"many2many:group_members;"`
	Invites      []Invite `gorm:"many2many:group_invites;"`
}
