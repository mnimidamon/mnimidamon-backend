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

type logComputerBackupImpl struct {
	MGBCase usecase.ManageComputerBackupInterface
	JAuth   restapi.JwtAuthentication
}

func (impl *logComputerBackupImpl) Handle(p backup.LogComputerBackupParams, _ interface{}) middleware.Responder {
	return impl.JAuth.ExtractUserFromApiKey(p.HTTPRequest, func(um *model.User) middleware.Responder {
		return impl.JAuth.ExtractComputerFromApiKey(p.HTTPRequest, um.ID, func(cm *model.Computer) middleware.Responder {
			backupID, groupID := uint(p.BackupID), uint(p.GroupID)
			return impl.JAuth.WithGroup(um, groupID, func(gm *model.Group) middleware.Responder {
				return impl.JAuth.WithGroupComputer(cm, gm, func(gcm *model.GroupComputer) middleware.Responder {
					return impl.JAuth.WithBackup(um, gm, backupID, func(bm *model.Backup) middleware.Responder {
						hash, prepend := *p.Body.Hash, *p.Body.PrependString
						cbm, err := impl.MGBCase.LogDownload(bm.ID, cm.ID, prepend, hash)

						if err != nil {
							if errors.Is(err, domain.ErrInternalDomain) {
								return backup.NewLogComputerBackupInternalServerError().
									WithPayload(ToRestError(err))
							}
							return backup.NewLogComputerBackupBadRequest().
								WithPayload(ToRestError(err))
						}

						return backup.NewLogComputerBackupOK().
							WithPayload(MapToComputerBackup(cbm))
					})
				})
			})
		})
	})
}

func NewLogComputerBackupHandler(mgbuc usecase.ManageComputerBackupInterface, ja restapi.JwtAuthentication) backup.LogComputerBackupHandler {
	return &logComputerBackupImpl{
		MGBCase: mgbuc,
		JAuth:   ja,
	}
}
