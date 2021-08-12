package handlers

import (
	"errors"
	"github.com/go-openapi/runtime/middleware"
	"mnimidamonbackend/adapter/restapi/authentication"
	"mnimidamonbackend/adapter/restapi/endpoints/operations/backup"
	"mnimidamonbackend/domain"
	"mnimidamonbackend/domain/model"
	"mnimidamonbackend/domain/usecase"
)

type downloadBackupImpl struct {
	MFCase usecase.ManageFileInterface
	JAuth  authentication.JwtAuthentication
}

func (impl *downloadBackupImpl) Handle(p backup.DownloadBackupParams, _ interface{}) middleware.Responder {
	return impl.JAuth.ExtractUserFromApiKey(p.HTTPRequest, func(um *model.User) middleware.Responder {
		return impl.JAuth.ExtractComputerFromApiKey(p.HTTPRequest, um.ID, func(cm *model.Computer) middleware.Responder {
			backupID, groupID := uint(p.BackupID), uint(p.GroupID)
			return impl.JAuth.WithGroup(um, groupID, func(gm *model.Group) middleware.Responder {
				return impl.JAuth.WithGroupComputer(cm, gm, func(gcm *model.GroupComputer) middleware.Responder {
					return impl.JAuth.WithBackup(um, gm, backupID, func(bm *model.Backup) middleware.Responder {
						rc, err := impl.MFCase.DownloadBackup(backupID)

						if err != nil {
							if errors.Is(err, domain.ErrInternalDomain) {
								return backup.NewDownloadBackupInternalServerError().
									WithPayload(ToRestError(err))
							}
							return backup.NewDownloadBackupBadRequest().
								WithPayload(ToRestError(err))
						}

						return backup.NewDownloadBackupOK().
							WithPayload(rc)
					})
				})
			})
		})
	})
}

func NewDownloadBackupImpl(mfuc usecase.ManageFileInterface, ja authentication.JwtAuthentication) backup.DownloadBackupHandler {
	return &downloadBackupImpl{
		MFCase: mfuc,
		JAuth: ja,
	}
}