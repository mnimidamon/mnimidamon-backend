package listbackup

import (
	"mnimidamonbackend/domain"
	"mnimidamonbackend/domain/model"
	"mnimidamonbackend/domain/repository"
	"mnimidamonbackend/domain/usecase"
)

type listBackupUseCase struct {
	BRepo repository.BackupRepository
}

func NewUseCase(br repository.BackupRepository) usecase.ListBackupInterface {
	return listBackupUseCase{
		BRepo: br,
	}
}

func (lb listBackupUseCase) FindById(backupID uint) (*model.Backup, error) {
	b, err := lb.BRepo.FindById(backupID)

	if err != nil {
		return nil, domain.ToDomainError(err)
	}

	return b, nil
}

func (lb listBackupUseCase) FindGroupBackups(groupID uint) ([]*model.Backup, error) {
	bl, err := lb.BRepo.FindAll(groupID)

	if err != nil {
		return nil, domain.ToDomainError(err)
	}

	return bl, nil
}




