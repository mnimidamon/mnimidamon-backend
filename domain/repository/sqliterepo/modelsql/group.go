package modelsql

// Represents a Group which has GroupMembers and a list of active Invites.
type Group struct {
	Entity

	Name string `gorm:"unique; index; size:15"`

	GroupMembers []User   `gorm:"many2many:group_members;"`
	Invites      []Invite `gorm:"many2many:group_invites;"`
}
