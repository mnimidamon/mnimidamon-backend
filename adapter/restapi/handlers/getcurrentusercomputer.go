package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"mnimidamonbackend/adapter/restapi"
	"mnimidamonbackend/adapter/restapi/endpoints/operations/computer"
	"mnimidamonbackend/domain/model"
	"mnimidamonbackend/domain/usecase"
)

type getCurrentUserComputer struct {
	LUCase usecase.ListUserInterface
	LCCase usecase.ListComputerInterface
	JAuth  restapi.JwtAuthentication
}

func (impl *getCurrentUserComputer) Handle(p computer.GetCurrentComputerParams, _ interface{}) middleware.Responder {
	return impl.JAuth.ExtractUserFromApiKey(p.HTTPRequest, func(um *model.User) middleware.Responder {
		return impl.JAuth.ExtractComputerFromApiKey(p.HTTPRequest, um.ID, func(cm *model.Computer) middleware.Responder {
			return computer.NewGetCurrentUserComputerOK().
				WithPayload(MapToComputer(cm))
		})
	})
}

func NewGetCurrentUserComputer(luuc usecase.ListUserInterface, lcuc usecase.ListComputerInterface, ja restapi.JwtAuthentication) computer.GetCurrentComputerHandler {
	return &getCurrentUserComputer{
		LUCase: luuc,
		LCCase: lcuc,
		JAuth:  ja,
	}
}
