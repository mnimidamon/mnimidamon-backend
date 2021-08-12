package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"mnimidamonbackend/adapter/restapi/authentication"
	"mnimidamonbackend/adapter/restapi/endpoints/operations/group"
	"mnimidamonbackend/domain/model"
	"mnimidamonbackend/domain/usecase"
)

type getGroupInvitesImpl struct {
	LICase usecase.ListInviteInterface
	JAuth  authentication.JwtAuthentication
}

func (impl *getGroupInvitesImpl) Handle(p group.GetGroupInvitesParams, _ interface{}) middleware.Responder {
	return impl.JAuth.ExtractUserFromApiKey(p.HTTPRequest, func(um *model.User) middleware.Responder {
		groupID := uint(p.GroupID)
		return impl.JAuth.WithGroup(um, groupID, func(gm *model.Group) middleware.Responder {
			is, err := impl.LICase.FindAllOfGroup(groupID)

			if err != nil {
				return group.NewGetGroupInvitesInternalServerError().
					WithPayload(ToRestError(err))
			}

			return group.NewGetGroupInvitesOK().
				WithPayload(MapToInvites(is))
		})
	})
}

func NewGetGroupInvitesHandler(liuc usecase.ListInviteInterface, ja authentication.JwtAuthentication) group.GetGroupInvitesHandler {
	return &getGroupInvitesImpl{
		LICase: liuc,
		JAuth:  ja,
	}
}
