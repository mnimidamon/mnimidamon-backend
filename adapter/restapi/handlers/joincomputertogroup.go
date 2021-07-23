package handlers

import (
	"errors"
	"github.com/go-openapi/runtime/middleware"
	"mnimidamonbackend/adapter/restapi"
	"mnimidamonbackend/adapter/restapi/endpoints/operations/group_computer"
	"mnimidamonbackend/domain"
	"mnimidamonbackend/domain/model"
	"mnimidamonbackend/domain/usecase"
)

type joinComputerToGroupImpl struct {
	MGCCase usecase.ManageGroupComputerInterface
	JAuth restapi.JwtAuthentication
}

func (impl *joinComputerToGroupImpl) Handle(p group_computer.JoinComputerToGroupParams, _ interface{}) middleware.Responder {
	return impl.JAuth.ExtractUserFromApiKey(p.HTTPRequest, func(um *model.User) middleware.Responder {
		return impl.JAuth.ExtractComputerFromApiKey(p.HTTPRequest, um.ID, func(cm *model.Computer) middleware.Responder {
			size, groupID := uint(p.Body.Size), uint(p.GroupID)
			gc, err := impl.MGCCase.JoinGroup(cm.ID, size, groupID)

			if err != nil {
				if errors.Is(err, domain.ErrInternalDomain) {
					return group_computer.NewJoinComputerToGroupInternalServerError().
						WithPayload(ToRestError(err))
				}
				return group_computer.NewJoinComputerToGroupBadRequest().
					WithPayload(ToRestError(err))
			}

			return group_computer.NewJoinComputerToGroupOK().
				WithPayload(MapToGroupComputer(gc))
		})
	})
}

func NewJoinComputerToGroupHandler(mgcuc usecase.ManageGroupComputerInterface, ja restapi.JwtAuthentication) group_computer.JoinComputerToGroupHandler {
	return &joinComputerToGroupImpl{
		MGCCase: mgcuc,
		JAuth:   ja,
	}
}
