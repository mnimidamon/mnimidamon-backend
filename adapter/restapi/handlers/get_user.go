package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"mnimidamonbackend/adapter/restapi/endpoints/operations/user"
	"mnimidamonbackend/domain/usecase"
)

type getUserImpl struct {
	LUCase usecase.ListUserInterface
}

func (impl *getUserImpl) Handle(p user.GetUserParams) middleware.Responder {
	userID := uint(p.UserID)

	u, err := impl.LUCase.FindById(userID)

	if IsNotFoundError(err) {
		return user.NewGetUserNotFound().
			WithPayload(nil)
	}

	return user.NewGetUserOK().
		WithPayload(MapToUser(u))
}

func NewGetUserHandler(luuc usecase.ListUserInterface) user.GetUserHandler {
	return &getUserImpl{
		LUCase: luuc,
	}
}

