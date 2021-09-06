package managecomputerbackup

import (
	"errors"
	"fmt"
	"io"
	"mnimidamonbackend/domain"
	"mnimidamonbackend/domain/model"
	"mnimidamonbackend/domain/repository"
	"mnimidamonbackend/domain/usecase"
	"time"
)

type manageComputerBackupUseCase struct {
	BRepo  repository.BackupRepository
	GCRepo repository.GroupComputerRepository
	CBRepo repository.ComputerBackupRepository
	FStore repository.FileStore
}

func (mgb manageComputerBackupUseCase) LogDownload(backupID uint, computerID uint, prefix string, hash string) (*model.ComputerBackup, error) {
	bm, err := mgb.BRepo.FindById(backupID)

	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return nil, domain.ErrBackupNotFound
		}
		return nil, domain.ToDomainError(err)
	}

	if !bm.OnServer {
		return nil, domain.ErrBackupNotOnServer
	}

	gc, err := mgb.GCRepo.FindById(bm.GroupID, computerID)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return nil, domain.ErrComputerNotFound
		}
		return nil, domain.ToDomainError(err)
	}

	rc, _ := mgb.FStore.GetFile(backupID)
	prc := NewPrefixReaderCloser(rc, []byte(prefix))

	hashCalculated, err := mgb.FStore.CalculateReaderHash(prc)
	rc.Close()

	if err != nil {
		return nil, fmt.Errorf("%w: %v -> %v", domain.ErrInternalDomain, err, "prefixed reader closer hash calculation error")
	}

	if hash != hashCalculated {
		return nil, domain.ErrInvalidPrefixedHash
	}

	cbm := &model.ComputerBackup{
		BackupID:        bm.ID,
		GroupComputerID: gc.ID,
		Backup:          nil,
		GroupComputer:   nil,
		CreatedAt:       time.Time{},
	}


	ts := repository.NewTransactionStack()
	// Start Transactions.
	cbrtx := mgb.CBRepo.BeginTx(); ts.Add(cbrtx)
	brtx := mgb.BRepo.ContinueTx(cbrtx); ts.Add(brtx)
	gcmtx := mgb.GCRepo.ContinueTx(cbrtx); ts.Add(gcmtx)

	defer ts.RollbackUnlessCommitted()


	err = cbrtx.Create(cbm)
	if err != nil {
		return nil, domain.ToDomainError(err)
	}

	if bm.OnServer {
		// Check if file should be deleted off the server.
		gcms, err := gcmtx.FindAllOfGroup(gc.GroupID)

		if err != nil {
			return nil, domain.ToDomainError(err)
		}

		shouldDelete := true
		for _, gcm := range gcms {
			// Check if the computer has it stored.
			exists, err := cbrtx.Exists(gcm.ID, bm.ID)

			if err != nil {
				return nil, domain.ToDomainError(err)
			}

			if exists {
				// This computer already has it stored.
				continue
			}

			// Get the stored size of the computer that does not have it stored.
			storedSize, err := cbrtx.FindStoredSizeOf(gcm.ID)

			if err != nil {
				return nil, domain.ToDomainError(err)
			}

			// This computer can also download it, so it should stay on the server until then.
			if storedSize + bm.Size <= gcm.StorageSize {
				shouldDelete = false
				break
			}
		}

		// Everyone has the file or does not have enough space.
		if shouldDelete {
			// Update the flags.
			bm.OnServer = false
			bm.UploadRequest = false


			err := brtx.Update(bm)
			if err != nil {
				return nil, domain.ToDomainError(err)
			}

			// Delete the file.
			err = mgb.FStore.DeleteFile(bm.ID)
			if err != nil {
				return nil, domain.ToDomainError(err)
			}
		}
	}

	ts.Commit()
	return cbm, nil
}

func NewUseCase(br repository.BackupRepository, gcr repository.GroupComputerRepository, cbr repository.ComputerBackupRepository, fs repository.FileStore) usecase.ManageComputerBackupInterface {
	return manageComputerBackupUseCase{
		BRepo:  br,
		GCRepo: gcr,
		FStore: fs,
		CBRepo: cbr,
	}
}


func NewPrefixReaderCloser(rc io.ReadCloser, prefix []byte) io.ReadCloser{
	return &prefixedReaderCloser{
		RC:     rc,
		Prefix: prefix,
		i:      0,
	}
}


type prefixedReaderCloser struct {
	RC     io.ReadCloser
	Prefix []byte
	i      int
}

func (prc *prefixedReaderCloser) Close() error {
	return prc.RC.Close()
}

func (prc *prefixedReaderCloser) Read(p []byte) (n int, err error) {
	toRead := len(p)

	// Prefix has already been read.
	if len(prc.Prefix) < prc.i + 1 {
		return prc.RC.Read(p)
	}

	// Copy prefix to byte
	n = copy(p, prc.Prefix[prc.i:])
	prc.i += n
	if n < toRead {
		x, err := prc.RC.Read(p[n:])

		if err != nil {
			return x + n, err
		}

		prc.i += x + n
		return x + n, nil
	}

	prc.i += n
	return n, nil
}
