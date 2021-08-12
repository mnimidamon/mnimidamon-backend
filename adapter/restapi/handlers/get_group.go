package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"mnimidamonbackend/adapter/restapi/authentication"
	"mnimidamonbackend/adapter/restapi/endpoints/operations/group"
	"mnimidamonbackend/domain/model"
)

type getGroupImpl struct {
	JAuth authentication.JwtAuthentication
}

func (impl *getGroupImpl) Handle(p group.GetGroupParams, _ interface{}) middleware.Responder {
	return impl.JAuth.ExtractUserFromApiKey(p.HTTPRequest, func(um *model.User) middleware.Responder {
		return impl.JAuth.WithGroup(um, uint(p.GroupID), func(gm *model.Group) middleware.Responder {
			return group.NewGetGroupOK().
				WithPayload(MapToGroup(gm))
		})
	})
}

func NewGetGroupHandler(ja authentication.JwtAuthentication) group.GetGroupHandler {
	return &getGroupImpl{
		JAuth: ja,
	}
}
