package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"mnimidamonbackend/adapter/restapi/authentication"
	"mnimidamonbackend/adapter/restapi/endpoints/operations/current_user"
	"mnimidamonbackend/domain/model"
)

type getCurrentUserImpl struct {
	JAuth authentication.JwtAuthentication
}

func (impl *getCurrentUserImpl) Handle(p current_user.GetCurrentUserParams, _ interface{}) middleware.Responder {
	return impl.JAuth.ExtractUserFromApiKey(p.HTTPRequest, func(um *model.User) middleware.Responder {
		return current_user.NewGetCurrentUserOK().
			WithPayload(MapToUser(um))
	})
}

func NewGetCurrentUserHandler(ja authentication.JwtAuthentication) current_user.GetCurrentUserHandler {
	return &getCurrentUserImpl{
		JAuth: ja,
	}
}