package managefile

import (
	"io"
	"mnimidamonbackend/domain"
	"mnimidamonbackend/domain/model"
	"mnimidamonbackend/domain/repository"
	"mnimidamonbackend/domain/usecase"
)

type manageFileUseCase struct {
	BRepo repository.BackupRepository
	FStore repository.FileStore
}

func (mf manageFileUseCase) UploadBackup(backupID uint, rc io.ReadCloser) (*model.Backup, error) {
	// Get Backup model.
	bm, err := mf.BRepo.FindById(backupID)

	if err != nil {
		return nil, domain.ToDomainError(err)
	}

	// If upload request is on.
	if !bm.UploadRequest {
		return nil, domain.ErrUploadNotRequested
	}

	/* TODO
	// Hash checking
	h := sha256.New()
	if _, err := io.Copy(h, rc); err != nil {
		constants.Log("Error calculating sha256 hash of backup %v: %v", backupID, err)
		return nil, domain.ErrCalculatingHash
	}

	correctHash := []byte(bm.Hash)
	calculatedHash := h.Sum(nil)

	constants.Log("CORR: %v CALC: %v", correctHash, calculatedHash)
	if  bytes.Compare(correctHash, calculatedHash) != 0 {
		return nil, domain.ErrInvalidBackupHash
	}
	*/

	// Save it to FileStore.
	err = mf.FStore.SaveFile(backupID, rc)
	if err != nil {
		return nil, domain.ToDomainError(err)
	}

	// Update on server flag
	bm.OnServer = true
	bm.UploadRequest = false

	err = mf.BRepo.Update(bm)
	if err != nil {
		return nil, domain.ToDomainError(err)
	}

	return bm, err
}

func (mf manageFileUseCase) DownloadBackup(backupID uint) (io.ReadCloser, error) {
	// Get Backup model.
	bm, err := mf.BRepo.FindById(backupID)

	if err != nil {
		return nil, domain.ToDomainError(err)
	}

	if !bm.OnServer {
		return nil, domain.ErrBackupNotOnServer
	}

	rc, err := mf.FStore.GetFile(backupID)
	if err != nil {
		return nil, domain.ToDomainError(err)
	}

	return rc, nil
}

func NewUseCase(fs repository.FileStore, br repository.BackupRepository) usecase.ManageFileInterface {
	return manageFileUseCase{
		BRepo:  br,
		FStore: fs,
	}
}
