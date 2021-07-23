package handlers

import (
	"errors"
	"github.com/go-openapi/runtime/middleware"
	"mnimidamonbackend/adapter/restapi"
	"mnimidamonbackend/adapter/restapi/endpoints/operations/computer"
	"mnimidamonbackend/domain"
	"mnimidamonbackend/domain/model"
	"mnimidamonbackend/domain/usecase"
)

type getCurrentUserComputerImpl struct {
	JAuth restapi.JwtAuthentication
	LCCase usecase.ListComputerInterface
}

func (impl *getCurrentUserComputerImpl) Handle(p computer.GetCurrentUserComputerParams, _ interface{}) middleware.Responder {
	return impl.JAuth.ExtractUserFromApiKey(p.HTTPRequest, func(um *model.User) middleware.Responder {
		computerID := uint(p.ComputerID)
		c, err := impl.LCCase.FindById(computerID)

		if errors.Is(err, domain.ErrInternalDomain) {
			return computer.NewGetCurrentUserComputerInternalServerError().
				WithPayload(ErrInternalServer)
		} else if err != nil {
			return computer.NewGetCurrentUserComputerBadRequest()
		}

		return computer.NewGetCurrentUserComputerOK().
			WithPayload(MapToComputer(c))
	})
}

func NewGetCurrentUserComputerHandler(lcuc usecase.ListComputerInterface, ja restapi.JwtAuthentication) computer.GetCurrentUserComputerHandler {
	return &getCurrentUserComputerImpl{
		JAuth:  ja,
		LCCase: lcuc,
	}
}
