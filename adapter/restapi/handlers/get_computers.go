package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"mnimidamonbackend/adapter/restapi/authentication"
	"mnimidamonbackend/adapter/restapi/endpoints/operations/computer"
	"mnimidamonbackend/domain/model"
	"mnimidamonbackend/domain/usecase"
)

type getComputersImpl struct {
	JAuth authentication.JwtAuthentication
	LCCase usecase.ListComputerInterface
}

func (impl *getComputersImpl) Handle(p computer.GetCurrentUserComputersParams, _ interface{}) middleware.Responder {
	return impl.JAuth.ExtractUserFromApiKey(p.HTTPRequest, func(um *model.User) middleware.Responder {
		cs, err := impl.LCCase.FindAllOfUser(um.ID)
		if err != nil {
			return computer.NewGetCurrentUserComputerInternalServerError().
				WithPayload(ToRestError(err))
		}

		return computer.NewGetCurrentUserComputersOK().
			WithPayload(MapToComputers(cs))
	})
}

func NewGetComputersHandler(ja authentication.JwtAuthentication, lcuc usecase.ListComputerInterface) computer.GetCurrentUserComputersHandler {
	return &getComputersImpl{
		JAuth:  ja,
		LCCase: lcuc,
	}
}
