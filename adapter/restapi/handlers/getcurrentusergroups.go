package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"mnimidamonbackend/adapter/restapi"
	"mnimidamonbackend/adapter/restapi/endpoints/operations/current_user"
	"mnimidamonbackend/domain/model"
	"mnimidamonbackend/domain/usecase"
)

type getCurrentUserGroupsImpl struct {
	LGCase usecase.ListGroupInterface
	JAuth  restapi.JwtAuthentication
}

func (impl *getCurrentUserGroupsImpl) Handle(p current_user.GetCurrentUserGroupsParams, _ interface{}) middleware.Responder {
	return impl.JAuth.ExtractUserFromApiKey(p.HTTPRequest, func(um *model.User) middleware.Responder {
		gms, err := impl.LGCase.FindAllOfUser(um.ID)

		if err != nil {
			return current_user.NewGetCurrentUserGroupsInternalServerError().
				WithPayload(ErrInternalServer)
		}

		return current_user.NewGetCurrentUserGroupsOK().
			WithPayload(MapToGroups(gms))
	})
}

func NewGetCurrentUserGroupsHandler(lguc usecase.ListGroupInterface, ja restapi.JwtAuthentication) current_user.GetCurrentUserGroupsHandler {
	return &getCurrentUserGroupsImpl{
		LGCase: lguc,
		JAuth:  ja,
	}
}
