package handlers

import (
	"errors"
	"github.com/go-openapi/runtime/middleware"
	"mnimidamonbackend/adapter/restapi"
	"mnimidamonbackend/adapter/restapi/endpoints/operations/backup"
	"mnimidamonbackend/domain"
	"mnimidamonbackend/domain/model"
	"mnimidamonbackend/domain/usecase"
)

type groupBackupDeletionImpl struct {
	MBCase usecase.ManageBackupInterface
	JAuth  restapi.JwtAuthentication
}

func (impl *groupBackupDeletionImpl) Handle(p backup.InitializeGroupBackupDeletionParams, _ interface{}) middleware.Responder {
	return impl.JAuth.ExtractUserFromApiKey(p.HTTPRequest, func(um *model.User) middleware.Responder {
		return impl.JAuth.ExtractComputerFromApiKey(p.HTTPRequest, um.ID, func(cm *model.Computer) middleware.Responder {
			groupID := uint(p.GroupID)
			backupID := uint(p.BackupID)

			return impl.JAuth.WithGroup(um, groupID, func(gm *model.Group) middleware.Responder {
				return impl.JAuth.WithGroupComputer(cm, gm, func(gcm *model.GroupComputer) middleware.Responder {
					return impl.JAuth.WithBackup(um, gm, backupID, func(bm *model.Backup) middleware.Responder {
						b, err := impl.MBCase.DeleteRequest(um.ID, backupID)

						if err != nil {
							if errors.Is(err, domain.ErrInternalDomain) {
								return backup.NewInitializeGroupBackupInternalServerError().
									WithPayload(ToRestError(err))
							} else {
								return backup.NewInitializeGroupBackupBadRequest().
									WithPayload(ToRestError(err))
							}
						}

						return backup.NewInitializeGroupBackupDeletionAccepted().
							WithPayload(MapToBackup(b))
					})
				})
			})
		})
	})
}

func NewGroupBackupDeletionImpl(mbuc usecase.ManageBackupInterface, ja restapi.JwtAuthentication) backup.InitializeGroupBackupDeletionHandler {
	return &groupBackupDeletionImpl{
		MBCase: mbuc,
		JAuth:  ja,
	}
}
