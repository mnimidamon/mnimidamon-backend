package handlers

import (
	"errors"
	"github.com/go-openapi/runtime/middleware"
	"mnimidamonbackend/adapter/restapi"
	"mnimidamonbackend/adapter/restapi/endpoints/operations/group"
	"mnimidamonbackend/domain"
	"mnimidamonbackend/domain/model"
	"mnimidamonbackend/domain/usecase"
	"mnimidamonbackend/domain/usecase/payload"
)

type createGroupImpl struct {
	MGCase usecase.ManageGroupInterface
	JAuth  restapi.JwtAuthentication
}

func (impl *createGroupImpl) Handle(p group.CreateGroupParams, _ interface{}) middleware.Responder {
	return impl.JAuth.ExtractUserFromApiKey(p.HTTPRequest, func(um *model.User) middleware.Responder {
		g, err := impl.MGCase.CreateGroup(payload.CreateGroupPayload{
			MemberID: um.ID,
			Name:     *p.Body.Name,
		})

		if errors.Is(err, domain.ErrGroupWithNameAlreadyExists) {
			return group.NewCreateGroupBadRequest().
				WithPayload(ToRestError(err))
		} else if err != nil {
			return group.NewCreateGroupInternalServerError().
				WithPayload(ErrInternalServer)
		}

		return group.NewCreateGroupOK().
			WithPayload(MapToGroup(g))
	})
}

func NewCreateGroupHandler(mguc usecase.ManageGroupInterface, ja restapi.JwtAuthentication) group.CreateGroupHandler {
	return &createGroupImpl{
		MGCase: mguc,
		JAuth:  ja,
	}
}
