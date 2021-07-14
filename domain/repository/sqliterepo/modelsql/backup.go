package modelsql

import (
	"mnimidamonbackend/domain/model"
	"time"
)

// Backup denotes a file that has been uploaded for backup by the Owner inside a Group.
// Flags:
//		- UploadRequest: file is requested to be uploaded from ComputerBackup owners (new member, more space, ...)
//		- DeleteRequest: file will no longer be backed up, all local storages ComputerBackup are up for deletion.
//		- OnServer: file is locally stored on server.
type Backup struct {
	Entity

	FileName string
	Size     uint
	Hash 	 string

	UploadRequest bool `gorm:"default:true"`
	DeleteRequest bool `gorm:"default:false"`
	OnServer      bool `gorm:"default:false"`

	OwnerID uint
	GroupID uint

	Owner *User  `gorm:"foreignKey:OwnerID"`
	Group *Group `gorm:"foreignKey:GroupID"`
}

func NewBackupFromBusinessModel(bm *model.Backup) *Backup {
	if bm == nil {
		return nil
	}

	return &Backup{
		Entity:        Entity{
			ID:        bm.ID,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
		},
		FileName:      bm.FileName,
		Size:          bm.Size,
		Hash:          bm.Hash,
		UploadRequest: bm.UploadRequest,
		DeleteRequest: bm.DeleteRequest,
		OnServer:      bm.OnServer,
		OwnerID:       bm.OwnerID,
		GroupID:       bm.GroupID,
		Owner:         NewUserFromBusinessModel(bm.Owner),
		Group:         NewGroupFromBusinessModel(bm.Group),
	}
}

func (b *Backup) NewBusinessModel() *model.Backup {
	if b == nil {
		return nil
	}

	bm := new(model.Backup)
	b.CopyToBusinessModel(bm)
	return bm
}

func (b *Backup) CopyToBusinessModel(bm *model.Backup) {
	if b == nil {
		bm = nil
		return
	}

	bm.ID = b.ID
	bm.FileName = b.FileName
	bm.Size = b.Size
	bm.Hash = b.Hash
	bm.UploadRequest = b.UploadRequest
	bm.DeleteRequest = b.DeleteRequest
	bm.OnServer = b.OnServer
	bm.OwnerID = b.OwnerID
	bm.GroupID = b.GroupID
	b.Owner.CopyToBusinessModel(bm.Owner)
	b.Group.CopyToBusinessModel(bm.Group)
}
