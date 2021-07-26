package filestore

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"math"
	"mnimidamonbackend/domain/model"
	"mnimidamonbackend/domain/repository"
	"os"
	"path/filepath"
	"strconv"
)

type fileStoreImpl struct {
	directoryPath string
}

func (fs fileStoreImpl) SaveFile(backup *model.Backup, rc io.ReadCloser) error {
	filePath := fs.getBackupFilePath(backup.ID)
	outFile, err := os.Create(filePath)

	if err != nil {
		return fmt.Errorf("%w: file store error when creating file for backup %v: %v", repository.ErrCreateFile, filePath, err)
	}

	defer outFile.Close()
	defer rc.Close()

	written, err := io.Copy(outFile, rc)
	if err != nil {
		return fmt.Errorf("%w: got err %v when saving file", repository.ErrSaveFile, err)
	}

	sizeKb := uint(math.Ceil(float64(written) / (1024.)))

	// If size is +- %1 * sizeKb
	if  !(uint(math.Ceil(.99 * float64(sizeKb))) <= backup.Size && backup.Size <= uint(math.Ceil(1.01 * float64(sizeKb)))) {
		_ = fs.DeleteFile(backup.ID)
		return fmt.Errorf("%w: invalid file size %v != +- %%1 %v", repository.ErrInvalidSize, sizeKb, backup.Size)
	}

	// Calculate hash
	f, _ := fs.GetFile(backup.ID)
	defer f.Close()

	calculatedHash, err := CalculateReaderCloserHash(f)
	correctHash := backup.Hash

	if err != nil {
		return err
	}

	if  correctHash != calculatedHash {
		_ = fs.DeleteFile(backup.ID)
		return fmt.Errorf("%w: expected %v got %v",repository.ErrInvalidBackupHash, correctHash, calculatedHash)
	}

	return nil
}

func (fs fileStoreImpl) GetFile(backupID uint) (io.ReadCloser, error) {
	filePath := fs.getBackupFilePath(backupID)
	file, err := os.Open(filePath)

	if err != nil {
		return nil, fmt.Errorf("%w: error when opening file %v: %v", repository.ErrOpenFile, filePath, err)
	}

	return file, nil
}

func (fs fileStoreImpl) DeleteFile(backupID uint) error {
	filePath := fs.getBackupFilePath(backupID)
	err := os.Remove(filePath)

	if err != nil {
		return fmt.Errorf("%w: error when deleting file %v: %v", repository.ErrFileDeletion, filePath, err)
	}

	return nil
}

func (fs fileStoreImpl) CalculateReaderCloserHash(rc io.Reader) (string, error) {
	h := sha256.New()

	if _, err := io.Copy(h, rc); err != nil {
		return "", fmt.Errorf("%w: error calculating hash of backup %v: %v", repository.ErrCalculatingHash, err)
	}

	calculatedHash := hex.EncodeToString(h.Sum(nil))
	return calculatedHash, nil
}

func (fs fileStoreImpl) getBackupFilePath(backupID uint) string {
	return filepath.Join(fs.directoryPath, strconv.Itoa(int(backupID)))
}

func New(directoryPath string) repository.FileStore {
	return fileStoreImpl{
		directoryPath: directoryPath,
	}
}
