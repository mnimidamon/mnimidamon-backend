package testsuites

import (
	"errors"
	"fmt"
	"mnimidamonbackend/domain/model"
	"mnimidamonbackend/domain/repository"
	"testing"
	"time"
)

type ComputerBackupRepositoryTester struct {
	Repo repository.ComputerBackupRepository

	GCRepo repository.GroupComputerRepository
	GRepo  repository.GroupRepository
	URepo  repository.UserRepository
	CRepo  repository.ComputerRepository
	BRepo  repository.BackupRepository

	*model.ComputerBackup
	GroupComputer *model.GroupComputer
	User          *model.User
	Group         *model.Group
	*model.Computer
	Backup *model.Backup
}

func (cbrt *ComputerBackupRepositoryTester) Setup(t *testing.T) {
	_, gcr, gr, ur, cr, br := cbrt.Repo, cbrt.GCRepo, cbrt.GRepo, cbrt.URepo, cbrt.CRepo, cbrt.BRepo
	_, gc, u, g, c, b := cbrt.ComputerBackup, cbrt.GroupComputer, cbrt.User, cbrt.Group, cbrt.Computer, cbrt.Backup

	/*
		cbr, gcr, gr, ur, cr, br := cbrt.Repo, cbrt.GCRepo, cbrt.GRepo, cbrt.URepo, cbrt.CRepo, cbrt.BRepo
		cb, gc, u, g, c, b := cbrt., cbrt.GroupComputer, cbrt.User, cbrt.Group, cbrt., cbrt.Backup
	*/

	t.Run("PreSaveOperations", func(t *testing.T) {
		if err := ur.Create(u); err != nil {
			t.Error(expectedGot("no error upon creating user", err))
		}

		if err := gr.Create(g); err != nil {
			t.Error(expectedGot("no error upon creating group", err))
		}

		c.OwnerID = u.ID
		if err := cr.Create(c, c.OwnerID); err != nil {
			t.Error(expectedGot("no error upon creating computer", err))
		}

		if _, err := gr.AddMember(u.ID, g.ID); err != nil {
			t.Error(expectedGot("no error upon adding user to group", err))
		}

		gc.ComputerID = c.ID
		gc.GroupID = g.ID
		if err := gcr.Create(gc); err != nil {
			t.Error(expectedGot("no error upon creating group computer", err))
		}

		b.OwnerID = u.ID
		b.GroupID = g.ID
		if err := br.Create(b); err != nil {
			t.Error(expectedGot("no error upon creating group backup", err))
		}
	})
}

func (cbrt *ComputerBackupRepositoryTester) FindBeforeSaveTests(t *testing.T) {
	cbr := cbrt.Repo
	gc, b := cbrt.GroupComputer, cbrt.Backup

	t.Run("FindByIdFail", func(t *testing.T) {
		gc, err := cbr.FindById(gc.ID, b.ID)

		if !errors.Is(repository.ErrNotFound, err) {
			t.Errorf(expectedGot(repository.ErrNotFound, fmt.Sprintf("Error: %v, ComputerBackup: %v", err, gc)))
		}
	})

	t.Run("FindAllOfGroupComputerEmpty", func(t *testing.T) {
		cbs, err := cbr.FindAllOfGroupComputer(gc.ID)

		if err != nil {
			t.Errorf(expectedGot("empty slice", err))
		}

		if len(cbs) != 0 {
			t.Errorf(expectedGot("empty slice", cbs))
		}
	})

	t.Run("FindAllOfBackupEmpty", func(t *testing.T) {
		computers, err := cbr.FindAllOfBackup(b.ID)

		if err != nil {
			t.Errorf(expectedGot("empty slice", err))
		}

		if len(computers) != 0 {
			t.Errorf(expectedGot("empty slice", computers))
		}
	})

	t.Run("ExistsNot", func(t *testing.T) {
		exists, err := cbr.Exists(gc.ID, b.ID)

		if err != nil {
			t.Error(unexpectedErr(err))
		}

		if exists {
			t.Error("Expected false, got true")
		}
	})
}

func (cbrt *ComputerBackupRepositoryTester) SaveSuccessfulTests(t *testing.T) {
	gcr := cbrt.Repo
	cb, gc, b := cbrt.ComputerBackup, cbrt.GroupComputer, cbrt.Backup

	t.Run("SaveSuccess", func(t *testing.T) {
		cb.GroupComputerID = gc.ID
		cb.BackupID = b.ID

		if err := gcr.Create(cb); err != nil {
			t.Error(unexpectedErr(err))
		}

		if cb.GroupComputerID != gc.ID {
			t.Error(expectedGot("right group_computer_id", cb))
		}

		if cb.BackupID != b.ID {
			t.Error(expectedGot("right backup_id", cb))
		}
	})

	t.Run("Exists", func(t *testing.T) {
		exists, err := gcr.Exists(gc.ID, b.ID)

		if err != nil {
			t.Error(unexpectedErr(err))
		}

		if !exists {
			t.Error("Expected true, got false")
		}
	})
}

func (cbrt *ComputerBackupRepositoryTester) FindAfterSaveTests(t *testing.T) {
	cbr := cbrt.Repo
	gc, b := cbrt.GroupComputer, cbrt.Backup

	t.Run("FindByIdSuccess", func(t *testing.T) {
		cn, err := cbr.FindById(gc.ID, b.ID)

		if err != nil {
			t.Error(unexpectedErr(err))
		}

		if cn.BackupID != b.ID {
			t.Error(expectedGot("right backup_id", cn))
		}

		if cn.GroupComputerID != gc.ID {
			t.Error(expectedGot("right group_computer_id", cn))
		}
	})

	t.Run("FindAllOfGroupComputerSuccess", func(t *testing.T) {
		computers, err := cbr.FindAllOfGroupComputer(gc.ID)

		if err != nil {
			t.Errorf(expectedGot("slice", err))
		}

		if len(computers) != 1 {
			t.Errorf(expectedGot("slice of 1", computers))
		}
	})

	t.Run("FindAllOfBackupSuccess", func(t *testing.T) {
		computers, err := cbr.FindAllOfBackup(b.ID)

		if err != nil {
			t.Errorf(expectedGot("slice of 1", err))
		}

		if len(computers) != 1 {
			t.Errorf(expectedGot("empty slice", computers))
		}
	})
}

func (cbrt *ComputerBackupRepositoryTester) ConstraintsTest(t *testing.T) {
	cbr := cbrt.Repo
	cb := cbrt.ComputerBackup

	t.Run("DuplicateInviteSaveFail", func(t *testing.T) {
		if err := cbr.Create(cb); !errors.Is(repository.ErrAlreadyExists, err) {
			t.Error(expectedGot(repository.ErrAlreadyExists, err))
		} else if err != nil && !errors.Is(repository.ErrAlreadyExists, err) {
			t.Error(unexpectedErr(err))
		}
	})
}

func (cbrt *ComputerBackupRepositoryTester) UpdateTests(t *testing.T) {
	t.Skip("No update functions, skipping")
}

func (cbrt *ComputerBackupRepositoryTester) SpecificTests(t *testing.T) {
	t.Skip("No specific tests")
}

func (cbrt *ComputerBackupRepositoryTester) DeleteTests(t *testing.T) {
	gcr := cbrt.Repo
	cb := cbrt.ComputerBackup

	t.Run("DeleteSuccess", func(t *testing.T) {
		err := gcr.Delete(cb.GroupComputerID, cb.BackupID)

		if err != nil {
			t.Error(unexpectedErr(err))
		}
	})

	t.Run("FindByIdFail", func(t *testing.T) {
		ggg, err := gcr.FindById(cb.GroupComputerID, cb.BackupID)

		if !errors.Is(repository.ErrNotFound, err) {
			t.Errorf(expectedGot(repository.ErrNotFound, fmt.Sprintf("Error: %v, GroupComputer: %v", err, ggg)))
		}
	})
}

func (cbrt *ComputerBackupRepositoryTester) BeginTx() TransactionSuiteTestTxInterface {
	return &ComputerBackupRepositoryTesterTx{
		Repo:           cbrt.Repo.BeginTx(),
		ComputerBackup: cbrt.ComputerBackup,
		GroupComputer:  cbrt.GroupComputer,
		Backup:         cbrt.Backup,
	}
}

func (cbrt *ComputerBackupRepositoryTester) Find() error {
	_, err := cbrt.Repo.FindById(cbrt.ComputerBackup.GroupComputerID, cbrt.ComputerBackup.BackupID)
	return err
}

type ComputerBackupRepositoryTesterTx struct {
	Repo repository.ComputerBackupRepositoryTx

	ComputerBackup *model.ComputerBackup
	GroupComputer  *model.GroupComputer
	Backup         *model.Backup
}

func (cbrtx *ComputerBackupRepositoryTesterTx) Create() error {
	return cbrtx.Repo.Create(cbrtx.ComputerBackup)
}

func (cbrtx *ComputerBackupRepositoryTesterTx) Find() error {
	_, err := cbrtx.Repo.FindById(cbrtx.ComputerBackup.GroupComputerID, cbrtx.ComputerBackup.BackupID)
	return err
}

func (cbrtx *ComputerBackupRepositoryTesterTx) CorrectCheck(t *testing.T) {
	cb, gc, b := cbrtx.ComputerBackup, cbrtx.GroupComputer, cbrtx.Backup
	if !(cb.BackupID == b.ID && cb.GroupComputerID == gc.ID) {
		t.Error(expectedGot("matching id's of backup and group_computer", cb))
	}
}

func (cbrtx *ComputerBackupRepositoryTesterTx) Rollback() error {
	return cbrtx.Repo.Rollback()
}

func (cbrtx *ComputerBackupRepositoryTesterTx) Commit() error {
	return cbrtx.Repo.Commit()
}

func ComputerBackupRepositoryTestSuite(t *testing.T, cbr repository.ComputerBackupRepository, gcr repository.GroupComputerRepository, gr repository.GroupRepository, ur repository.UserRepository, cr repository.ComputerRepository, br repository.BackupRepository) {
	marmiha := &model.User{
		Entity:       model.Entity{},
		Username:     "marmiha",
		PasswordHash: "marmiha_hash",
	}

	mnimidamons := &model.Group{
		Entity: model.Entity{},
		Name:   "mnimidamons",
	}

	thinkpad := &model.Computer{
		Entity:  model.Entity{},
		OwnerID: 0,
		Name:    "thinkpad",
		Owner:   nil,
	}

	mnimidamons_thinkpad := &model.GroupComputer{
		Entity:      model.Entity{},
		GroupID:     0,
		ComputerID:  0,
		Group:       nil,
		Computer:    nil,
		StorageSize: 1024,
	}

	documents := &model.Backup{
		Entity:        model.Entity{},
		FileName:      "pictures",
		Size:          1024,
		Hash:          "pictures_hash",
		UploadRequest: true,
		DeleteRequest: false,
		OnServer:      false,
		OwnerID:       0,
		GroupID:       0,
		Owner:         nil,
		Group:         nil,
	}

	documents_mnimidamons_thinkpad := &model.ComputerBackup{
		BackupID:        0,
		GroupComputerID: 0,
		Backup:          nil,
		GroupComputer:   nil,
		CreatedAt:       time.Time{},
	}

	cbrt := &ComputerBackupRepositoryTester{
		Repo:           cbr,
		GCRepo:         gcr,
		GRepo:          gr,
		URepo:          ur,
		CRepo:          cr,
		BRepo:          br,
		Backup:         documents,
		ComputerBackup: documents_mnimidamons_thinkpad,
		GroupComputer:  mnimidamons_thinkpad,
		User:           marmiha,
		Group:          mnimidamons,
		Computer:       thinkpad,
	}

	runCommonRepositoryTests(cbrt, t)
}
