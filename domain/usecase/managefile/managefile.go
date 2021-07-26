package managefile

import (
	"errors"
	"io"
	"mnimidamonbackend/domain"
	"mnimidamonbackend/domain/constants"
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

	// Save it to FileStore.
	err = mf.FStore.SaveFile(bm, rc)
	if err != nil {
		if errors.Is(err, repository.ErrInvalidBackupHash) || errors.Is(err, repository.ErrInvalidSize) {
			constants.Log("%v", err)
			return nil, domain.ErrInvalidFile
		}
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
