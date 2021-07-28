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

type uploadBackupImpl struct {
	MFCase usecase.ManageFileInterface
	JAuth  authentication.JwtAuthentication
}

func (impl *uploadBackupImpl) Handle(p backup.UploadBackupParams, _ interface{}) middleware.Responder {
	return impl.JAuth.ExtractUserFromApiKey(p.HTTPRequest, func(um *model.User) middleware.Responder {
		return impl.JAuth.ExtractComputerFromApiKey(p.HTTPRequest, um.ID, func(cm *model.Computer) middleware.Responder {
			backupID, groupID := uint(p.BackupID), uint(p.GroupID)
			return impl.JAuth.WithGroup(um, groupID, func(gm *model.Group) middleware.Responder {
				return impl.JAuth.WithGroupComputer(cm, gm, func(gcm *model.GroupComputer) middleware.Responder {
					return impl.JAuth.WithBackup(um, gm, backupID, func(bm *model.Backup) middleware.Responder {
						b, err := impl.MFCase.UploadBackup(backupID, p.BackupData)

						if err != nil {
							if errors.Is(err, domain.ErrInternalDomain) {
								return backup.NewUploadBackupInternalServerError().
									WithPayload(ToRestError(err))
							}
							return backup.NewUploadBackupBadRequest().
								WithPayload(ToRestError(err))
						}

						return backup.NewUploadBackupOK().
							WithPayload(MapToBackup(b))
					})
				})
			})
		})
	})
}

func NewUploadBackupHandler(mfuc usecase.ManageFileInterface, ja authentication.JwtAuthentication) backup.UploadBackupHandler {
	return &uploadBackupImpl{
		MFCase: mfuc,
		JAuth:  ja,
	}
}
