package handlers

import (
	"mnimidamonbackend/adapter/restapi/modelapi"
	"mnimidamonbackend/domain/model"
)

func MapToUser(um *model.User) *modelapi.User {
	return &modelapi.User{
		UserID:   int64(um.ID),
		Username: um.Username,
	}
}

func MapToUsers(ums []*model.User) []*modelapi.User {
	var us []*modelapi.User

	for _, um := range ums {
		u := MapToUser(um)
		us = append(us, u)
	}

	return us
}