package model

type Computer struct {
	Entity

	OwnerID uint
	Name    string

	Owner User
}
