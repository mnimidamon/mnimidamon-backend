package managegroupcomputer

import (
	"errors"
	"mnimidamonbackend/domain"
	"mnimidamonbackend/domain/model"
	"mnimidamonbackend/domain/repository"
	"mnimidamonbackend/domain/usecase"
)

type manageGroupComputerUseCase struct {
	GCRepo repository.GroupComputerRepository
	CRepo  repository.ComputerRepository
	GRepo  repository.GroupRepository
	BRepo  repository.BackupRepository
	CBRepo repository.ComputerBackupRepository
}

func (mgc manageGroupComputerUseCase) Update(groupComputerID uint, size uint) (*model.GroupComputer, error) {
	panic("implement me") // TODO
}

func (mgc manageGroupComputerUseCase) JoinGroup(computerID uint, size uint, groupID uint) (*model.GroupComputer, error) {
	// Get the computer
	c, err := mgc.CRepo.FindById(computerID)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return nil, domain.ErrComputerNotFound
		}
		return nil, domain.ToDomainError(err)
	}

	// Check if user is member of group
	isMember, err := mgc.GRepo.IsMemberOf(c.OwnerID, groupID)
	if err != nil {
		return nil, domain.ToDomainError(err)
	}

	if !isMember {
		return nil, domain.ErrUserNotGroupMember
	}


	// Check if computer is already member of group, also stands for group exist check
	_, err = mgc.GCRepo.FindById(groupID, computerID)
	if err == nil {
		return nil, domain.ErrComputerAlreadyMember
	} else if !errors.Is(err, repository.ErrNotFound){
		return nil, domain.ToDomainError(err)
	}

	gcm := &model.GroupComputer{
		Entity:      model.Entity{},
		GroupID:     groupID,
		ComputerID:  computerID,
		Group:       nil,
		Computer:    nil,
		StorageSize: size,
	}

	gcrtx := mgc.GCRepo.BeginTx()
	if err := gcrtx.Create(gcm); err != nil {
		return nil, domain.ToDomainError(err)
	}

	backups, err := mgc.BRepo.FindAll(groupID)
	if err != nil {
		return nil, err
	}

	var onServerSize uint = 0
	var uploadRequestSize uint = 0

	for _, b := range backups {
		// Skip the backups ready for deletion.
		if b.DeleteRequest {
			continue
		}

		if b.UploadRequest {
			uploadRequestSize += b.Size
		} else if b.OnServer {
			onServerSize += b.Size
		}
	}

	// Check if enough backups are on the server or requested to fill the space.
	if onServerSize + uploadRequestSize >= size {
		gcrtx.Commit()
		return gcm, nil
	}

	brtx := mgc.BRepo.ContinueTx(gcrtx)
	// Else update request to upload for enough backups.
	var aditionalUploadRequestSize uint = 0
	for _, b := range backups {
		// Skip these.
		if b.DeleteRequest || b.UploadRequest || b.OnServer{
			continue
		}

		// Request the backup upload.
		b.UploadRequest = true
		if err := brtx.Update(b); err != nil {
			// Rollback all changes if it fails.
			brtx.Rollback()
			gcrtx.Rollback()
			return nil, domain.ToDomainError(err)
		}

		// Check if enough size is requested.
		aditionalUploadRequestSize += b.Size
		if uploadRequestSize + onServerSize + aditionalUploadRequestSize >= size {
			break
		}
	}

	// Commit the changes.
	gcrtx.Commit()
	brtx.Commit()

	return gcm, nil
}

func (mgc manageGroupComputerUseCase) LeaveGroup(computerID uint, size uint) error {
	panic("implement me") // TODO
}

func NewUseCase(gcr repository.GroupComputerRepository, cr repository.ComputerRepository, gr repository.GroupRepository, br repository.BackupRepository, cbr repository.ComputerBackupRepository) usecase.ManageGroupComputerInterface {
	return manageGroupComputerUseCase{
		GCRepo: gcr,
		GRepo:  gr,
		BRepo:  br,
		CBRepo: cbr,
		CRepo:  cr,
	}
}
