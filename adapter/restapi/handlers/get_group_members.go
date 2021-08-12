package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"mnimidamonbackend/adapter/restapi/authentication"
	"mnimidamonbackend/adapter/restapi/endpoints/operations/group"
	"mnimidamonbackend/domain/model"
	"mnimidamonbackend/domain/usecase"
)

type getGroupMembersImpl struct {
	LGMCase usecase.ListGroupMemberInterface
	JAuth   authentication.JwtAuthentication
}

func (impl *getGroupMembersImpl) Handle(p group.GetGroupMembersParams, _ interface{}) middleware.Responder {
	return impl.JAuth.ExtractUserFromApiKey(p.HTTPRequest, func(um *model.User) middleware.Responder {
		groupID := uint(p.GroupID)
		return impl.JAuth.WithGroup(um,groupID , func(gm *model.Group) middleware.Responder {
			gms, err := impl.LGMCase.FindAllMembersOfGroup(gm.ID)

			if err != nil {
				return group.NewGetGroupMembersInternalServerError().
					WithPayload(ToRestError(err))
			}

			return group.NewGetGroupMembersOK().
				WithPayload(MapToUsers(gms))
		})
	})
}

func NewGetGroupMembersHandler(lgmuc usecase.ListGroupMemberInterface, ja authentication.JwtAuthentication) group.GetGroupMembersHandler {
	return &getGroupMembersImpl{
		LGMCase: lgmuc,
		JAuth:   ja,
	}
}