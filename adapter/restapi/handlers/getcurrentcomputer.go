package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"mnimidamonbackend/adapter/restapi"
	"mnimidamonbackend/adapter/restapi/endpoints/operations/computer"
	"mnimidamonbackend/domain/model"
)

type getCurrentComputerImpl struct {
	JAuth  restapi.JwtAuthentication
}

func (impl *getCurrentComputerImpl) Handle(p computer.GetCurrentComputerParams, _ interface{}) middleware.Responder {
	return impl.JAuth.ExtractUserFromApiKey(p.HTTPRequest, func(um *model.User) middleware.Responder {
		return impl.JAuth.ExtractComputerFromApiKey(p.HTTPRequest, um.ID, func(cm *model.Computer) middleware.Responder {
			return computer.NewGetCurrentUserComputerOK().
				WithPayload(MapToComputer(cm))
		})
	})
}

func NewGetCurrentUserComputer(ja restapi.JwtAuthentication) computer.GetCurrentComputerHandler {
	return &getCurrentComputerImpl{
		JAuth:  ja,
	}
}
