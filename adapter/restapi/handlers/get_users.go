package handlers

import (
	"errors"
	"github.com/go-openapi/runtime/middleware"
	"mnimidamonbackend/adapter/restapi/endpoints/operations/user"
	"mnimidamonbackend/adapter/restapi/modelapi"
	"mnimidamonbackend/domain"
	"mnimidamonbackend/domain/usecase"
)

type getUsersImpl struct {
	LUCase usecase.ListUserInterface
}

func (impl *getUsersImpl) Handle(p user.GetUsersParams) middleware.Responder {
	username := p.Filter

	// If username filter is set.
	if username != nil {
		um, err := impl.LUCase.FindByUsername(*username)

		if errors.Is(err, domain.ErrNotFound) {

			return user.NewGetUsersOK().
				WithPayload([]*modelapi.User{})
		} else if err != nil {

			return user.NewGetUsersInternalServerError().
				WithPayload(ErrInternalServer)
		}

		return user.NewGetUsersOK().
				WithPayload([]*modelapi.User{
					MapToUser(um),
				})
	}

	ums, err := impl.LUCase.FindAll()
	if err != nil {
		return user.NewGetUsersInternalServerError().
			WithPayload(ErrInternalServer)
	}

	return user.NewGetUsersOK().
		WithPayload(MapToUsers(ums))
}

func NewGetUsersHandler(luuc usecase.ListUserInterface) user.GetUsersHandler {
	return &getUsersImpl{
		LUCase: luuc,
	}
}
