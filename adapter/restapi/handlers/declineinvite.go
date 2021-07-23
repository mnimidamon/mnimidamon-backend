package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"mnimidamonbackend/adapter/restapi"
	"mnimidamonbackend/adapter/restapi/endpoints/operations/invite"
	"mnimidamonbackend/domain/model"
	"mnimidamonbackend/domain/usecase"
)

type declineCurrentUserInviteImpl struct {
	GICase usecase.GroupInviteInterface
	JAuth  restapi.JwtAuthentication
}

func (impl *declineCurrentUserInviteImpl) Handle(p invite.DeclineCurrentUserInviteParams, _ interface{}) middleware.Responder {
	return impl.JAuth.ExtractUserFromApiKey(p.HTTPRequest, func(um *model.User) middleware.Responder {
		groupID := uint(p.GroupID)
		return impl.JAuth.WithInvite(um, groupID, func(im *model.Invite) middleware.Responder {
			err := impl.GICase.DeclineInvite(um.ID, im.GroupID)

			if err != nil {
				return invite.NewDeclineCurrentUserInviteInternalServerError().
						WithPayload(ToRestError(err))
			}

			return invite.NewDeclineCurrentUserInviteNoContent()
		})
	})
}

func NewDeclineCurrentUserInviteHandler(giuc usecase.GroupInviteInterface, ja restapi.JwtAuthentication) invite.DeclineCurrentUserInviteHandler {
	return &declineCurrentUserInviteImpl{
		GICase: giuc,
		JAuth:  ja,
	}
}
