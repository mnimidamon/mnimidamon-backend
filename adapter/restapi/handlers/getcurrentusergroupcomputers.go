package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"mnimidamonbackend/adapter/restapi/authentication"
	"mnimidamonbackend/adapter/restapi/endpoints/operations/computer"
	"mnimidamonbackend/domain/model"
	"mnimidamonbackend/domain/repository"
)

type getCurrentUserGroupComputers struct {
	JAuth authentication.JwtAuthentication
	GCRepo repository.GroupComputerRepository
}

func (impl *getCurrentUserGroupComputers) Handle(p computer.GetCurrentUserGroupComputersParams, _ interface{}) middleware.Responder {
	return impl.JAuth.ExtractUserFromApiKey(p.HTTPRequest, func(um *model.User) middleware.Responder {
		return impl.JAuth.ExtractComputerFromApiKey(p.HTTPRequest, um.ID, func(cm *model.Computer) middleware.Responder {
			gcs, err := impl.GCRepo.FindAllOfComputer(cm.ID)

			if err != nil {
				return computer.NewGetCurrentUserGroupComputersInternalServerError().
					WithPayload(ToRestError(err))
			}

			return computer.NewGetCurrentUserGroupComputersOK().
				WithPayload(MapToGroupComputers(gcs))
		})
	})
}

func NewGetCurrentUserGroupComputersHandler(jwtAuthentication authentication.JwtAuthentication, groupComputerRepository repository.GroupComputerRepository) computer.GetCurrentUserGroupComputersHandler {
	return &getCurrentUserGroupComputers{
		JAuth:  jwtAuthentication,
		GCRepo: groupComputerRepository,
	}
}
