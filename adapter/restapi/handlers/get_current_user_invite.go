package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"mnimidamonbackend/adapter/restapi/authentication"
	"mnimidamonbackend/adapter/restapi/endpoints/operations/invite"
	"mnimidamonbackend/domain/model"
)

type getCurrentUserInviteImpl struct {
	JAuth authentication.JwtAuthentication
}

func (impl *getCurrentUserInviteImpl) Handle(p invite.GetCurrentUserInviteParams, _ interface{}) middleware.Responder {
	return impl.JAuth.ExtractUserFromApiKey(p.HTTPRequest, func(um *model.User) middleware.Responder {
		groupID := uint(p.GroupID)
		return impl.JAuth.WithInvite(um, groupID, func(im *model.Invite) middleware.Responder {
			return invite.NewGetCurrentUserInviteOK().
				WithPayload(MapToInvite(im))
		})
	})
}

func NewGetCurrentUserInviteHandler(ja authentication.JwtAuthentication) invite.GetCurrentUserInviteHandler {
	return &getCurrentUserInviteImpl{
		JAuth:  ja,
	}
}
