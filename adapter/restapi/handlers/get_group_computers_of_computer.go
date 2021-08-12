package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"mnimidamonbackend/adapter/restapi/authentication"
	"mnimidamonbackend/adapter/restapi/endpoints/operations/group_computer"
	"mnimidamonbackend/domain/model"
	"mnimidamonbackend/domain/usecase"
)

type getGroupComputersOfComputerImpl struct {
	JAuth   authentication.JwtAuthentication
	LGCCase usecase.ListGroupComputerInterface
}

func (impl *getGroupComputersOfComputerImpl) Handle(p group_computer.GetGroupComputersOfComputerParams, _ interface{}) middleware.Responder {
	return impl.JAuth.ExtractUserFromApiKey(p.HTTPRequest, func(um *model.User) middleware.Responder {
		computerID := uint(p.ComputerID)
		return impl.JAuth.WithComputer(um, computerID, func(cm *model.Computer) middleware.Responder {
			gcs, err := impl.LGCCase.FindAllOfComputer(computerID)

			if err != nil {
				return group_computer.NewGetGroupComputersOfComputerInternalServerError().
					WithPayload(ToRestError(err))
			}

			return group_computer.NewGetGroupComputersOfComputerOK().
				WithPayload(MapToGroupComputers(gcs))
		})
	})
}

func NewGetGroupComputersOfComputerHandler(jwtAuthentication authentication.JwtAuthentication, computerInterface usecase.ListGroupComputerInterface) group_computer.GetGroupComputersOfComputerHandler {
	return &getGroupComputersOfComputerImpl{
		JAuth:   jwtAuthentication,
		LGCCase: computerInterface,
	}
}