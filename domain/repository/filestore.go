package repository

import (
	"io"
	"mnimidamonbackend/domain/model"
)

type FileStore interface {
	SaveFile(backup *model.Backup, rc io.ReadCloser) error
	GetFile(backupID uint) (io.ReadCloser, error)
	DeleteFile(backupID uint) error
	CalculateReaderCloserHash(rc io.Reader) (string, error)
}
