package testsuites

import (
	"errors"
	"fmt"
	"mnimidamonbackend/domain/model"
	"mnimidamonbackend/domain/repository"
	"testing"
)

type BackupRepositoryTester struct {
	Repo  repository.BackupRepository
	GRepo repository.GroupRepository
	URepo repository.UserRepository

	Backup       model.Backup
	SecondBackup model.Backup

	User       model.User
	SecondUser model.User

	Group       model.Group
	SecondGroup model.Group
}

func (brt *BackupRepositoryTester) FindBeforeSaveTests(t *testing.T) {
	br, _, _ := brt.Repo, brt.GRepo, brt.URepo

	g := brt.Group

	t.Run("FindAllByGroupIdEmpty", func(t *testing.T) {
		backups, err := br.FindAll(g.ID)

		if err != nil {
			t.Errorf(expectedGot("empty slice", err))
		}

		if len(backups) != 0 {
			t.Errorf(expectedGot("empty slice", backups))
		}
	})

	t.Run("FindByIdNotFound", func(t *testing.T) {
		b, err := br.FindById(1)
		if !errors.Is(repository.ErrNotFound, err) {
			t.Errorf(expectedGot(repository.ErrNotFound, fmt.Sprintf("Error: %v, Backup: %v", err, b)))
		}
	})
}

func (brt *BackupRepositoryTester) SaveSuccessfulTests(t *testing.T) {
	br, gr, ur := brt.Repo, brt.GRepo, brt.URepo

	u, g, b := brt.User, brt.Group, brt.Backup
	us, gs, bs := brt.SecondUser, brt.SecondGroup, brt.SecondBackup


	t.Run("PreSaveOperations", func(t *testing.T) {
		err := ur.Create(&u)
		if err != nil {
			t.Error(expectedGot("no error upon creating first user", err))
		}

		err = ur.Create(&us)
		if err != nil {
			t.Error(expectedGot("no error upon creating second user", err))
		}

		err = gr.Create(&g)
		if err != nil {
			t.Error(expectedGot("no error upon creating first group", err))
		}
		err = gr.Create(&gs)
		if err != nil {
			t.Error(expectedGot("no error upon creating second group", err))
		}

		// First user is owner the first backup
		// Second user is owner of the second backup
		b.OwnerID = u.ID
		bs.OwnerID = us.ID

		// First backup is of the first group.
		// Second backup is also a part of the first group.
		b.GroupID = g.ID
		bs.GroupID = g.ID
	})

	t.Run("SuccessfulSave", func(t *testing.T) {
		if err := br.Create(&b); err != nil {
			t.Error(expectedGot("no error on first backup creation", err))
		}

		if err := br.Create(&bs); err != nil {
			t.Error(expectedGot("no error on second backup creation", err))
		}

		if b.ID == 0 {
			t.Error(expectedGot("FirstGroup.ID > 0", b))
		}

		if bs.ID == 0 {
			t.Error(expectedGot("SecondGroup.ID > 0", bs))
		}
	})
}

func (brt *BackupRepositoryTester) FindAfterSaveTests(t *testing.T) {
	panic("implement me")
}

func (brt *BackupRepositoryTester) ConstraintsTest(t *testing.T) {
	panic("implement me")
}

func (brt *BackupRepositoryTester) UpdateTests(t *testing.T) {
	panic("implement me")
}

func (brt *BackupRepositoryTester) SpecificTests(t *testing.T) {
	panic("implement me")
}

func (brt *BackupRepositoryTester) DeleteTests(t *testing.T) {
	panic("implement me")
}

func (brt *BackupRepositoryTester) BeginTx() TransactionSuiteTestTxInterface {
	btx := brt.Repo.BeginTx()
	return &BackupRepositoryTesterTx{
		Repo:   btx,
		Backup: brt.Backup,
	}
}

func (brt *BackupRepositoryTester) Find() error {
	_, err := brt.Repo.FindById(brt.Backup.ID)
	return err
}

type BackupRepositoryTesterTx struct {
	Repo   repository.BackupRepositoryTx
	Backup model.Backup
}

func (btx *BackupRepositoryTesterTx) Create() error {
	return btx.Repo.Create(&btx.Backup)
}

func (btx *BackupRepositoryTesterTx) Find() error {
	_, err := btx.Repo.FindById(btx.Backup.ID)
	return err
}

func (btx *BackupRepositoryTesterTx) CorrectCheck(t *testing.T) {
	if btx.Backup.ID == 0 {
		t.Errorf("Expected group.ID > 0, got %v", btx.Backup)
	}
}

func (btx *BackupRepositoryTesterTx) Rollback() error {
	return btx.Repo.Rollback()
}

func (btx *BackupRepositoryTesterTx) Commit() error {
	return btx.Repo.Commit()
}

// Tests the repository.GroupRepository interface implementation against common tests.
func BackupRepositoryTestSuite(t *testing.T, br repository.BackupRepository, gr repository.GroupRepository, ur repository.UserRepository) {
	marmiha, marpeter := model.User{
		Entity:       model.Entity{},
		Username:     "marmiha",
		PasswordHash: "marmiha_hash",
	}, model.User{
		Entity:       model.Entity{},
		Username:     "marpeter",
		PasswordHash: "marpeter_hash",
	}

	mnimidamons, guccigang := model.Group{
		Entity: model.Entity{},
		Name:   "mnimidamons",
	}, model.Group{
		Entity: model.Entity{},
		Name:   "guccigang",
	}

	pictures, documents := model.Backup{
		Entity:        model.Entity{},
		FileName:      "pictures",
		Size:          1024,
		Hash:          "pictures_hash",
		UploadRequest: true,
		DeleteRequest: false,
		OnServer:      false,
		OwnerID:       0,
		GroupID:       0,
		Owner:         model.User{},
		Group:         model.Group{},
	}, model.Backup{
		Entity:        model.Entity{},
		FileName:      "documents",
		Size:          1000,
		Hash:          "documents_hash",
		UploadRequest: true,
		DeleteRequest: false,
		OnServer:      false,
		OwnerID:       0,
		GroupID:       0,
		Owner:         model.User{},
		Group:         model.Group{},
	}

	brt := &BackupRepositoryTester{
		Repo:         br,
		GRepo:        gr,
		URepo:        ur,
		Backup:       pictures,
		SecondBackup: documents,
		User:         marmiha,
		SecondUser:   marpeter,
		Group:        mnimidamons,
		SecondGroup:  guccigang,
	}

	runCommonRepositoryTests(brt, t)
	runTransactionTestSuite(brt, t)
}
