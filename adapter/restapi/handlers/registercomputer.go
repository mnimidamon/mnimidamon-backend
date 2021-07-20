package handlers

import (
	"errors"
	"github.com/go-openapi/runtime/middleware"
	"mnimidamonbackend/adapter/restapi"
	"mnimidamonbackend/adapter/restapi/endpoints/operations/authorization"
	"mnimidamonbackend/adapter/restapi/modelapi"
	"mnimidamonbackend/domain"
	"mnimidamonbackend/domain/model"
	"mnimidamonbackend/domain/usecase"
	"mnimidamonbackend/domain/usecase/payload"
)

type registerComputerImpl struct {
	JAuth  restapi.JwtAuthentication
	CRCase usecase.ComputerRegistrationInterface
}

func (impl *registerComputerImpl) Handle(p authorization.RegisterComputerParams, _ interface{}) middleware.Responder {
	return impl.JAuth.ExtractUserFromApiKey(p.HTTPRequest, func(um *model.User) middleware.Responder {
		name := *p.Body.Name

		c, err := impl.CRCase.RegisterComputer(payload.ComputerCredentialsPayload{
			OwnerID: um.ID,
			Name:    name,
		})

		if err != nil {
			if errors.Is(err, domain.ErrNameNotUnique) {
				return authorization.NewRegisterComputerBadRequest().
					WithPayload(ToRestError(err))
			}
			return authorization.NewRegisterComputerBadRequest().
				WithPayload(ToRestError(err))
		}

		token, err := impl.JAuth.GenerateComputerToken(c.ID)

		if err != nil {
			return authorization.NewRegisterComputerInternalServerError().
				WithPayload(ToRestError(err))
		}

		return authorization.NewRegisterComputerOK().
			WithPayload(&modelapi.CreateComputerResponse{
				CompKey:  *token,
				Computer: MapToComputer(c),
			})
	})
}

func NewRegisterComputerHandler(crcuc usecase.ComputerRegistrationInterface, ja restapi.JwtAuthentication) authorization.RegisterComputerHandler {
	return &registerComputerImpl{
		CRCase: crcuc,
		JAuth: ja,
	}
}
