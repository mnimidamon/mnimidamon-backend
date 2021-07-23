package handlers

import (
	"errors"
	"github.com/go-openapi/runtime/middleware"
	"mnimidamonbackend/adapter/restapi"
	"mnimidamonbackend/adapter/restapi/endpoints/operations/backup"
	"mnimidamonbackend/domain"
	"mnimidamonbackend/domain/model"
	"mnimidamonbackend/domain/usecase"
	"mnimidamonbackend/domain/usecase/payload"
)

type initializeGroupBackupImpl struct {
	MBCase usecase.ManageBackupInterface
	JAuth  restapi.JwtAuthentication
}

func (impl *initializeGroupBackupImpl) Handle(p backup.InitializeGroupBackupParams, _ interface{}) middleware.Responder {
	return impl.JAuth.ExtractUserFromApiKey(p.HTTPRequest, func(um *model.User) middleware.Responder {
		return impl.JAuth.ExtractComputerFromApiKey(p.HTTPRequest, um.ID, func(cm *model.Computer) middleware.Responder {
			groupID := uint(p.GroupID)

			return impl.JAuth.WithGroup(um, groupID, func(gm *model.Group) middleware.Responder {
				return impl.JAuth.WithGroupComputer(cm, gm, func(gcm *model.GroupComputer) middleware.Responder {
					fileName, size, hash := *p.Body.FileName, uint(*p.Body.Size), *p.Body.Hash

					p := payload.InitializeBackupPayload{
						FileName: fileName,
						Size:     size,
						Hash:     hash,
						OwnerID:  um.ID,
						GroupID:  gm.ID,
					}

					b, err := impl.MBCase.InitializeBackup(p)

					if err != nil {
						if errors.Is(err, domain.ErrInternalDomain) {
							return backup.NewInitializeGroupBackupInternalServerError().
								WithPayload(ToRestError(err))
						} else {
							return backup.NewInitializeGroupBackupBadRequest().
								WithPayload(ToRestError(err))
						}
					}

					return backup.NewInitializeGroupBackupOK().
						WithPayload(MapToBackup(b))
				})
			})
		})
	})
}

func NewInitializeGroupBackupHandler(mbuc usecase.ManageBackupInterface, ja restapi.JwtAuthentication) backup.InitializeGroupBackupHandler {
	return &initializeGroupBackupImpl{
		MBCase: mbuc,
		JAuth:  ja,
	}
}
