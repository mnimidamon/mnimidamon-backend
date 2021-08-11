package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"mnimidamonbackend/adapter/restapi/authentication"
	"mnimidamonbackend/adapter/restapi/endpoints/operations/computer"
	"mnimidamonbackend/domain/model"
	"mnimidamonbackend/domain/usecase"
)

type deleteComputerImpl struct {
	CRCase usecase.ComputerRegistrationInterface
	JAuth  authentication.JwtAuthentication
}

func (impl *deleteComputerImpl) Handle(p computer.DeleteComputerParams, _ interface{}) middleware.Responder {
	return impl.JAuth.ExtractUserFromApiKey(p.HTTPRequest, func(um *model.User) middleware.Responder {
		computerID := uint(p.ComputerID)
		return impl.JAuth.WithComputer(um, computerID, func(cm *model.Computer) middleware.Responder {
			err := impl.CRCase.UnregisterComputer(computerID)

			if err != nil {
				return computer.NewDeleteComputerInternalServerError().
					WithPayload(ToRestError(err))
			}

			return computer.NewDeleteComputerNoContent()
		})
	})
}

func NewDeleteComputerHandler(ja authentication.JwtAuthentication, cruc usecase.ComputerRegistrationInterface) computer.DeleteComputerHandler {
	return &deleteComputerImpl{
		CRCase: cruc,
		JAuth:  ja,
	}
}
