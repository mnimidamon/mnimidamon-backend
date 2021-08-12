package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"mnimidamonbackend/adapter/restapi/authentication"
	"mnimidamonbackend/adapter/restapi/endpoints/operations/current_user"
	"mnimidamonbackend/domain/model"
	"mnimidamonbackend/domain/usecase"
)

type getCurrentUserInvitesImpl struct {
	JAuth  authentication.JwtAuthentication
	LICase usecase.ListInviteInterface
}

func (impl *getCurrentUserInvitesImpl) Handle(p current_user.GetCurrentUserInvitesParams, _ interface{}) middleware.Responder {
	return impl.JAuth.ExtractUserFromApiKey(p.HTTPRequest, func(um *model.User) middleware.Responder {
		is, err := impl.LICase.FindAllOfUser(um.ID)

		if err != nil {
			return current_user.NewGetCurrentUserInvitesInternalServerError().
				WithPayload(ToRestError(err))
		}

		return current_user.NewGetCurrentUserInvitesOK().
			WithPayload(MapToInvites(is))
	})
}

func NewGetCurrentUserInvitesHandler(liuc usecase.ListInviteInterface, ja authentication.JwtAuthentication) current_user.GetCurrentUserInvitesHandler {
	return &getCurrentUserInvitesImpl{
		JAuth: ja,
		LICase: liuc,
	}
}