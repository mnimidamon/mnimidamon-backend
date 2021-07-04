package testsuites

import (
	"errors"
	"mnimidamonbackend/domain/model"
	"mnimidamonbackend/domain/repository"
	"testing"
)

// TransactionSuiteTestInterface implementation for repository.GroupRepository
type GroupRepositoryTester struct {
	Repo repository.GroupRepository
	Group model.Group
}

func (grt *GroupRepositoryTester) FindBeforeSaveTests(t *testing.T) {
	gr := grt.Repo

	t.Run("FindAllEmpty", func(t *testing.T) {
		groups, err := gr.FindAll()

		if err != nil {
			t.Errorf("Expected empty slice, got an error %v", err)
		}

		if len(groups) != 0 {
			t.Errorf("Expected empty slice, recieved %v", groups)
		}
	})

	t.Run("FindByNameErrNotFound", func(t *testing.T) {
		_, err := gr.FindByName(grt.Group.Name)

		if !errors.Is(repository.ErrNotFound, err) {
			t.Errorf("Expected %v, recieved %v", repository.ErrNotFound, err)
		}
	})

	t.Run("FindByIdErrNotFound", func(t *testing.T) {
		_, err := gr.FindById(grt.Group.ID)

		if !errors.Is(repository.ErrNotFound, err) {
			t.Errorf("Expected %v, recieved %v", repository.ErrNotFound, err)
		}
	})
}

func (grt *GroupRepositoryTester) SaveSuccessfulTests(t *testing.T) {
	gr := grt.Repo

	t.Run("SaveSuccess", func(t *testing.T) {
		err := gr.Create(&grt.Group)

		if err != nil {
			t.Errorf("Unexpected error %v", err)
		}

		if grt.Group.ID == 0 {
			t.Errorf("Expected ID greater 1, got %v", grt.Group)
		}
	})
}

func (grt *GroupRepositoryTester) FindAfterSaveTests(t *testing.T) {
	gr := grt.Repo

	t.Run("FindByNameSuccess", func(t *testing.T) {
		gg, err := gr.FindByName(grt.Group.Name[0:2])

		if err != nil {
			t.Errorf("Expected no error, recieved %v", err)
		}

		if gg.Name != grt.Group.Name {
			t.Errorf("Expected %v, got %v", grt.Group, gg)
		}

	})

	t.Run("FindByIdSuccess", func(t *testing.T) {
		gg, err := gr.FindById(1)

		if err != nil {
			t.Errorf("Expected no error, recieved %v", err)
		}

		if gg.ID != grt.Group.ID {
			t.Errorf("Expected %v, got %v", grt.Group, gg)
		}
	})
}

func (grt *GroupRepositoryTester) UpdateTests(t *testing.T) {
	// TODO
}

func (grt *GroupRepositoryTester) ConstraintsTest(t *testing.T) {
	gr := grt.Repo
	t.Run("SaveNonUniqueNameFail", func(t *testing.T) {
		if err := gr.Create(&grt.Group); !errors.Is(repository.ErrUniqueConstraintViolation, err) {
			t.Errorf("Expected %v, got %v", repository.ErrUniqueConstraintViolation, err)
		}

	})
}

func (grt *GroupRepositoryTester) SpecificTests(t *testing.T) {
	// TODO
}

func (grt *GroupRepositoryTester) DeleteTests(t *testing.T) {
	// TODO
}

func (grt *GroupRepositoryTester) BeginTx() TransactionSuiteTestTxInterface {
	gtx := grt.Repo.BeginTx()
	return &GroupRepositoryTesterTx{
		Repo:  gtx,
		Group: grt.Group,
	}
}

func (grt *GroupRepositoryTester) Find() error {
	_, err := grt.Repo.FindByName(grt.Group.Name)
	return err
}

type GroupRepositoryTesterTx struct {
	Repo repository.GroupRepositoryTx
	Group model.Group
}

func (gtx *GroupRepositoryTesterTx) Create() error {
	return gtx.Repo.Create(&gtx.Group)
}

func (gtx *GroupRepositoryTesterTx) Find() error {
	_, err := gtx.Repo.FindByName(gtx.Group.Name)
	return err
}

func (gtx *GroupRepositoryTesterTx) CorrectCheck(t *testing.T) {
	if gtx.Group.ID == 0 {
		t.Errorf("Expected group.ID > 0, got %v", gtx.Group)
	}
}

func (gtx *GroupRepositoryTesterTx) Rollback() error {
	return gtx.Repo.Rollback()
}

func (gtx *GroupRepositoryTesterTx) Commit() error {
	return gtx.Repo.Commit()
}

// Tests the repository.GroupRepository interface implementation against common tests.
func GroupRepositoryTestSuite(t *testing.T, gr repository.GroupRepository) {
	guccigang, mnimidamons := model.Group{
		Entity: model.Entity{},
		Name:   "guccigang",
	}, model.Group{
		Entity:       model.Entity{},
		Name:     "mnimidamons",
	}


	grt := &GroupRepositoryTester{
		Repo:  gr,
		Group: guccigang,
	}

	runCommonRepositoryTests(grt, t)

	grt.Group = mnimidamons
	runTransactionRollbackSuccessSuite(grt, t)
}
