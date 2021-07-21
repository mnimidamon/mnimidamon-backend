package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"mnimidamonbackend/adapter/restapi"
	"mnimidamonbackend/adapter/restapi/endpoints/operations/group"
	"mnimidamonbackend/domain/model"
	"mnimidamonbackend/domain/usecase"
)

type inviteUserToGroupImpl struct {
	GICase usecase.GroupInviteInterface
	LUCase usecase.ListUserInterface
	JAuth  restapi.JwtAuthentication
}

func (impl *inviteUserToGroupImpl) Handle(p group.InviteUserToGroupParams, _ interface{}) middleware.Responder {
	return impl.JAuth.ExtractUserFromApiKey(p.HTTPRequest, func(um *model.User) middleware.Responder {
		return impl.JAuth.WithGroup(um, uint(p.GroupID), func(gm *model.Group) middleware.Responder {
			username := *p.Body.Username

			u, err := impl.LUCase.FindByUsername(username)

			if IsInternalError(err) {
				return group.NewInviteUserToGroupInternalServerError().
					WithPayload(ErrInternalServer)
			} else if IsNotFoundError(err) {
				return group.NewInviteUserToGroupBadRequest()
			} else if err != nil {
				return group.NewInviteUserToGroupInternalServerError().
					WithPayload(ToRestError(err))
			}

			// Does not match the name fully.
			if u.Username != username {
				return group.NewInviteUserToGroupBadRequest()
			}

			i, err := impl.GICase.InviteUser(u.ID, gm.ID)
			if IsInternalError(err) {
				return group.NewInviteUserToGroupInternalServerError().
					WithPayload(ErrInternalServer)
			}

			if err != nil {
				return group.NewInviteUserToGroupBadRequest().
					WithPayload(ToRestError(err))
			}

			return group.NewInviteUserToGroupOK().
				WithPayload(MapToInvite(i))
		})
	})
}

func NewInviteUserToGroupHandler(giuc usecase.GroupInviteInterface, luuc usecase.ListUserInterface, ja restapi.JwtAuthentication) group.InviteUserToGroupHandler {
	return &inviteUserToGroupImpl{
		GICase: giuc,
		LUCase: luuc,
		JAuth:  ja,
	}
}
