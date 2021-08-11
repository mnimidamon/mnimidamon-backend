package managebackup

import (
	"errors"
	"mnimidamonbackend/domain"
	"mnimidamonbackend/domain/constants"
	"mnimidamonbackend/domain/model"
	"mnimidamonbackend/domain/repository"
	"mnimidamonbackend/domain/usecase"
	"mnimidamonbackend/domain/usecase/payload"
)

type manageBackupUseCase struct {
	BRepo  repository.BackupRepository
	URepo  repository.UserRepository
	GRepo  repository.GroupRepository
	CBRepo repository.ComputerBackupRepository
	CRepo  repository.ComputerRepository
	GCRepo repository.GroupComputerRepository
	FStore repository.FileStore
}

func NewUseCase(fs repository.FileStore, br repository.BackupRepository, ur repository.UserRepository, gr repository.GroupRepository, cr repository.ComputerRepository, gcr repository.GroupComputerRepository, cbr repository.ComputerBackupRepository) usecase.ManageBackupInterface {
	return manageBackupUseCase{
		BRepo:  br,
		URepo:  ur,
		GRepo:  gr,
		CRepo:  cr,
		GCRepo: gcr,
		CBRepo: cbr,
		FStore: fs,
	}
}

func (mb manageBackupUseCase) InitializeBackup(p payload.InitializeBackupPayload) (*model.Backup, error) {
	// Check if size is ok
	if p.Size <= 0 {
		return nil, domain.ErrInvalidBackupSize
	}

	// Filename should not be empty
	if p.FileName == "" {
		return nil, domain.ErrInvalidFileName
	}

	// Check if owner exists
	u, err := mb.URepo.FindById(p.OwnerID)
	if errors.Is(domain.ErrNotFound, err) {
		return nil, domain.ErrUserNotFound
	}

	if err != nil {
		return nil, domain.ToDomainError(err)
	}

	// TODO: Check if owner has a group computer inside the group

	// Check if group exists
	g, err := mb.GRepo.FindById(p.GroupID)
	if errors.Is(domain.ErrNotFound, err) {
		return nil, domain.ErrGroupNotFound
	}

	if err != nil {
		return nil, domain.ToDomainError(err)
	}

	b := &model.Backup{
		Entity:        model.Entity{},
		FileName:      p.FileName,
		Size:          p.Size,
		Hash:          p.Hash,
		UploadRequest: true,
		DeleteRequest: false,
		OnServer:      false,
		OwnerID:       p.OwnerID,
		GroupID:       p.GroupID,
		Owner:         nil,
		Group:         nil,
	}

	// Initialize the backup
	if err := mb.BRepo.Create(b); err != nil {
		return nil, domain.ToDomainError(err)
	}

	b.Owner = u
	b.Group = g

	return b, nil
}

func (mb manageBackupUseCase) UploadRequest(ownerID uint, backupID uint) (*model.Backup, error) {
	// Get user
	u, err := mb.URepo.FindById(ownerID)
	if errors.Is(domain.ErrNotFound, err) {
		return nil, domain.ErrUserNotFound
	}

	if err != nil {
		return nil, domain.ToDomainError(err)
	}

	// Get backup
	b, err := mb.BRepo.FindById(backupID)
	if errors.Is(domain.ErrNotFound, err) {
		return nil, domain.ErrBackupNotFound
	}

	if err != nil {
		return nil, domain.ToDomainError(err)
	}

	// Check if right owner
	if u.ID != b.OwnerID {
		return nil, domain.ErrUserNotOwner
	}

	// Check if backup is not on server already
	if b.OnServer {
		return nil, domain.ErrBackupAlreadyOnServer
	}

	// Check if backup is not requested for deletion
	if b.DeleteRequest {
		return nil, domain.ErrBackupWaitingForDeletion
	}

	// Enable the upload flag
	b.UploadRequest = true

	// Update backup
	if err := mb.BRepo.Update(b); err != nil {
		return nil, domain.ToDomainError(err)
	}

	b.Owner = u
	return b, nil
}

func (mb manageBackupUseCase) DeleteRequest(userID uint, backupID uint) error {
	// Get user
	_, err := mb.URepo.FindById(userID)
	if errors.Is(domain.ErrNotFound, err) {
		return domain.ErrUserNotFound
	}

	if err != nil {
		return domain.ToDomainError(err)
	}

	// Get backup
	b, err := mb.BRepo.FindById(backupID)
	if errors.Is(domain.ErrNotFound, err) {
		return domain.ErrBackupNotFound
	}

	if err != nil {
		return domain.ToDomainError(err)
	}

	if b.OwnerID != userID {
		return domain.ErrUserNotOwner
	}

	ts := repository.NewTransactionStack()

	cbtx := mb.CBRepo.BeginTx(); ts.Add(cbtx)
	brtx := mb.BRepo.ContinueTx(cbtx); ts.Add(brtx)

	defer ts.RollbackUnlessCommitted()

	// Find all the Computer backups.
	cbs, err := cbtx.FindAllOfBackup(backupID)
	if err != nil {
		return domain.ToDomainError(err)
	}

	// Delete every computer backups of every user.
	for _, cb := range cbs {
		if err := cbtx.Delete(cb.GroupComputerID, cb.BackupID); err != nil {
			constants.Log("error when deleting computer backup %v, %v", cb, err)
			return domain.ToDomainError(err)
		}
	}

	// Delete the backup.
	if err := brtx.Delete(b.ID); err != nil {
		constants.Log("error deleting backup %v, %v", b, err)
		return domain.ToDomainError(err)
	}

	// Delete the local stored file.
	if b.OnServer {
		if err := mb.FStore.DeleteFile(b.ID); err != nil {
			constants.Log("error deleting backup file %v, %v", b, err)
		}
	}

	// TODO Recompute the Upload flags.

	// If everything went ok commit.
	ts.Commit()
	return nil
}

func (mb manageBackupUseCase) findGroupComputersOfUserAndGroup(userID uint, groupID uint, cr repository.ComputerRepository, gc repository.GroupComputerRepository) ([]*model.GroupComputer, error) {
	// If the user is owner get his computers.
	computers, err := cr.FindAll(userID)
	if err != nil {
		return nil, domain.ToDomainError(err)
	}

	var cIDS []uint
	for _, c := range computers {
		cIDS = append(cIDS, c.ID)
	}

	// For each computer get the group member computer of the backups group
	groupComputers, err := gc.FindAllOfGroupAndComputers(groupID, cIDS...)
	if err != nil {
		return nil, domain.ToDomainError(err)
	}

	return groupComputers, nil
}
