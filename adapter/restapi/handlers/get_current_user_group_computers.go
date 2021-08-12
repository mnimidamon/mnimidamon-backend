package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"mnimidamonbackend/adapter/restapi/authentication"
	"mnimidamonbackend/adapter/restapi/endpoints/operations/computer"
	"mnimidamonbackend/domain/model"
	"mnimidamonbackend/domain/usecase"
)

type getCurrentUserGroupComputers struct {
	JAuth authentication.JwtAuthentication
	LGCCase usecase.ListGroupComputerInterface
}

func (impl *getCurrentUserGroupComputers) Handle(p computer.GetCurrentUserGroupComputersParams, _ interface{}) middleware.Responder {
	return impl.JAuth.ExtractUserFromApiKey(p.HTTPRequest, func(um *model.User) middleware.Responder {
		return impl.JAuth.ExtractComputerFromApiKey(p.HTTPRequest, um.ID, func(cm *model.Computer) middleware.Responder {
			groupID := uint(p.GroupID)
			return impl.JAuth.WithGroup(um, groupID , func(gm *model.Group) middleware.Responder {
				gcs, err := impl.LGCCase.FindAllOfGroup(groupID)

				if err != nil {
					return computer.NewGetCurrentUserGroupComputersInternalServerError().
						WithPayload(ToRestError(err))
				}

				return computer.NewGetCurrentUserGroupComputersOK().
					WithPayload(MapToGroupComputers(gcs))
			})

		})
	})
}

func NewGetCurrentUserGroupComputersHandler(jwtAuthentication authentication.JwtAuthentication, computerInterface usecase.ListGroupComputerInterface) computer.GetCurrentUserGroupComputersHandler {
	return &getCurrentUserGroupComputers{
		JAuth:  jwtAuthentication,
		LGCCase: computerInterface,
	}
}
