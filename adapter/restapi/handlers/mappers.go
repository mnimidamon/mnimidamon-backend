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
