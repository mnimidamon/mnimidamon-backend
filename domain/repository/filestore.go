package repository

import (
	"io"
)

type FileStore interface {
	SaveFile(backupID uint, rc io.ReadCloser) error
	GetFile(backupID uint) (io.ReadCloser, error)
	DeleteFile(backupID uint) error
}
