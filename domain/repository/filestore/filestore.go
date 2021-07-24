package filestore

import (
	"io"
	"mnimidamonbackend/domain/constants"
	"mnimidamonbackend/domain/repository"
	"os"
	"path/filepath"
	"strconv"
)

type fileStoreImpl struct {
	directoryPath string
}

func (fs fileStoreImpl) SaveFile(backupID uint, rc io.ReadCloser) error {
	filePath := fs.getBackupFilePath(backupID)
	outFile, err := os.Create(filePath)

	if err != nil {
		constants.Log("FileStore error when creating file for backup %v: %v", filePath, backupID, err)
		return repository.ErrCreateFile
	}

	defer outFile.Close()
	_, err = io.Copy(outFile, rc)

	if err != nil {
		constants.Log("FileStore error when copying ReadCloser to %v: %v", err)
		return repository.ErrSaveFile
	}

	return nil
}

func (fs fileStoreImpl) GetFile(backupID uint) (io.ReadCloser, error) {
	filePath := fs.getBackupFilePath(backupID)
	file, err := os.Open(filePath)

	if err != nil {
		constants.Log("FileStore error when opening file %v: %v", filePath, err)
		return nil, repository.ErrOpenFile
	}

	return file, nil
}

func (fs fileStoreImpl) DeleteFile(backupID uint) error {
	filePath := fs.getBackupFilePath(backupID)
	err := os.Remove(filePath)

	if err != nil {
		constants.Log("FileStore error when deleting file %v: %v", filePath, err)
		return repository.ErrFileDeletion
	}

	return nil
}

func (fs fileStoreImpl) getBackupFilePath(backupID uint) string {
	return filepath.Join(fs.directoryPath, strconv.Itoa(int(backupID)))
}

func New(directoryPath string) repository.FileStore {
	return fileStoreImpl{
		directoryPath: directoryPath,
	}
}
