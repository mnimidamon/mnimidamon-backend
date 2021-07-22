package handlers

import (
	"errors"
	"github.com/go-openapi/runtime/middleware"
	"mnimidamonbackend/adapter/restapi"
	"mnimidamonbackend/adapter/restapi/endpoints/operations/invite"
	"mnimidamonbackend/domain"
	"mnimidamonbackend/domain/model"
	"mnimidamonbackend/domain/usecase"
)

type acceptInviteImpl struct {
	GICase usecase.GroupInviteInterface
	JAuth  restapi.JwtAuthentication
}

func (impl *acceptInviteImpl) Handle(p invite.AcceptCurrentUserInviteParams, _ interface{}) middleware.Responder {
	return impl.JAuth.ExtractUserFromApiKey(p.HTTPRequest, func(um *model.User) middleware.Responder {
		groupID := uint(p.GroupID)
		return impl.JAuth.WithInvite(um, groupID, func(im *model.Invite) middleware.Responder {
			gm, err := impl.GICase.AcceptInvite(um.ID, groupID)

			if err != nil {
				if errors.Is(err, domain.ErrInternalDomain) {
					return invite.NewAcceptCurrentUserInviteInternalServerError().
						WithPayload(ErrInternalServer)
				}
				return invite.NewAcceptCurrentUserInviteUnauthorized().
					WithPayload(ToRestError(err))
			}

			return invite.NewAcceptCurrentUserInviteOK().
				WithPayload(MapToGroup(gm))
		})
	})
}

func NewAcceptInviteHandler(giuc usecase.GroupInviteInterface, ja restapi.JwtAuthentication) invite.AcceptCurrentUserInviteHandler {
	return &acceptInviteImpl{
		GICase: giuc,
		JAuth:  ja,
	}
}
