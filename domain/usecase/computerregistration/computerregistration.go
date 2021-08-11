package computerregistration

import (
	"errors"
	. "mnimidamonbackend/domain"
	"mnimidamonbackend/domain/model"
	"mnimidamonbackend/domain/repository"
	"mnimidamonbackend/domain/usecase"
	"mnimidamonbackend/domain/usecase/payload"
	"mnimidamonbackend/domain/usecase/procedure"
)

type computerRegistrationUseCase struct {
	CRepo  repository.ComputerRepository
	BRepo  repository.BackupRepository
	CBRepo repository.ComputerBackupRepository
	GCRepo repository.GroupComputerRepository
	URepo  repository.UserRepository
}

func (cr computerRegistrationUseCase) UnregisterComputer(computerID uint)  error {
	c, err := cr.CRepo.FindById(computerID)
	if err != nil {
		return ToDomainError(err)
	}

	ts := repository.NewTransactionStack()
	ctx := cr.CRepo.BeginTx(); ts.Add(ctx)
	brtx := cr.BRepo.ContinueTx(ctx); ts.Add(brtx)
	gctx := cr.GCRepo.ContinueTx(ctx); ts.Add(gctx)
	cbtx :=  cr.CBRepo.ContinueTx(ctx); ts.Add(cbtx)

	defer ts.RollbackUnlessCommitted()

	if err := procedure.DeleteComputer(c, gctx, cbtx, ctx, brtx); err != nil {
		return err
	}

	// Commit the changes
	ts.Commit()
	return nil
}

func (cr computerRegistrationUseCase) RegisterComputer(p payload.ComputerCredentialsPayload) (*model.Computer, error) {
	// Find the owner
	u, err := cr.URepo.FindById(p.OwnerID)
	if err != nil {
		return nil,ToDomainError(err)
	}

	// Find if name is unique
	_, err = cr.CRepo.FindByName(p.Name, u.ID)

	if err == nil {
		return nil, ErrNameNotUnique
	} else {
		if !errors.Is(err, repository.ErrNotFound) {
			return nil, ToDomainError(err)
		}
	}

	// Create
	cm := &model.Computer{
		Entity:  model.Entity{},
		OwnerID: u.ID,
		Name:    p.Name,
		Owner:   nil,
	}

	if err := cr.CRepo.Create(cm, u.ID); err != nil {
		return nil, ToDomainError(err)
	}

	return cm, nil
}

func NewUseCase(cr repository.ComputerRepository, ur repository.UserRepository, cbr repository.ComputerBackupRepository, gcr repository.GroupComputerRepository, br repository.BackupRepository) usecase.ComputerRegistrationInterface {
	return computerRegistrationUseCase{
		CRepo:  cr,
		CBRepo: cbr,
		GCRepo: gcr,
		URepo:  ur,
		BRepo: br,
	}
}
