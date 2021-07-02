package model

type Group struct {
	Entity

	Name string

	GroupMembers []User
	Invites      []Invite
}