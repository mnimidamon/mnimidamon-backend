package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"mnimidamonbackend/adapter/restapi/authentication"
	"mnimidamonbackend/adapter/restapi/endpoints/operations/backup"
	"mnimidamonbackend/domain/model"
	"mnimidamonbackend/domain/usecase"
)

type getGroupBackupsImpl struct {
	LBCase usecase.ListBackupInterface
	JAuth  authentication.JwtAuthentication
}

func (impl *getGroupBackupsImpl) Handle(p backup.GetGroupBackupsParams, _ interface{}) middleware.Responder {
	return impl.JAuth.ExtractUserFromApiKey(p.HTTPRequest, func(um *model.User) middleware.Responder {
		return impl.JAuth.ExtractComputerFromApiKey(p.HTTPRequest, um.ID, func(cm *model.Computer) middleware.Responder {
			groupID := uint(p.GroupID)
			return impl.JAuth.WithGroup(um, groupID, func(gm *model.Group) middleware.Responder {
				bs, err := impl.LBCase.FindGroupBackups(groupID)

				if err != nil {
					return backup.NewGetGroupBackupsInternalServerError().
						WithPayload(ToRestError(err))
				}

				return backup.NewGetGroupBackupsOK().
					WithPayload(MapToBackups(bs))
			})
		})
	})
}

func NewGetGroupBackupsHandler(lbuc usecase.ListBackupInterface, ja authentication.JwtAuthentication) backup.GetGroupBackupsHandler {
	return &getGroupBackupsImpl{
		LBCase: lbuc,
		JAuth:  ja,
	}
}
