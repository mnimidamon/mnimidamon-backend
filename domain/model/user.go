package model

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Entity

	Username     string
	PasswordHash string
}

func NewUser(username string, password string) (*User, error) {
	hash, err := hashPassword(password)
	if err != nil {
		return nil, ErrPasswordHash
	}

	return &User{
		PasswordHash: *hash,
		Username: username,
	}, nil
}

func (u *User) VerifyPassword(password string) error {
	bytePassword, byteHashedPassword := []byte(password), []byte(u.PasswordHash)
	return bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
}

func hashPassword(password string) (*string, error) {
	bytePassword := []byte(password)
	passwordHash, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	hashed := string(passwordHash)
	return &hashed, nil
}
