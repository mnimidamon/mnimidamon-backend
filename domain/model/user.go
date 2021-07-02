package model

type User struct {
	Entity

	Username string
	PasswordHash string
}
