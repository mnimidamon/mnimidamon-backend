package testsuites

import (
	"errors"
	"mnimidamonbackend/domain/model"
	"mnimidamonbackend/domain/repository"
	"testing"
)

// TransactionSuiteTestInterface implementation for repository.GroupRepository
type GroupRepositoryTransactionSuiteImpl struct {
	Repo repository.GroupRepository
	Group model.Group

}

func (g *GroupRepositoryTransactionSuiteImpl) BeginTx() TransactionSuiteTestTxInterface {
	gtx := g.Repo.BeginTx()
	return &GroupRepositoryTransactionSuiteTxImpl{
		Repo:  gtx,
		Group: g.Group,
	}
}

func (g *GroupRepositoryTransactionSuiteImpl) Find() error {
	_, err := g.Repo.FindByName(g.Group.Name)
	return err
}

type GroupRepositoryTransactionSuiteTxImpl struct {
	Repo repository.GroupRepositoryTx
	Group model.Group
}

func (gtx *GroupRepositoryTransactionSuiteTxImpl) Create() error {
	return gtx.Repo.Create(&gtx.Group)
}

func (gtx *GroupRepositoryTransactionSuiteTxImpl) Find() error {
	_, err := gtx.Repo.FindByName(gtx.Group.Name)
	return err
}

func (gtx *GroupRepositoryTransactionSuiteTxImpl) CorrectCheck(t *testing.T) {
	if gtx.Group.ID == 0 {
		t.Errorf("Expected group.ID > 0, got %v", gtx.Group)
	}
}

func (gtx *GroupRepositoryTransactionSuiteTxImpl) Rollback() error {
	return gtx.Repo.Rollback()
}

func (gtx *GroupRepositoryTransactionSuiteTxImpl) Commit() error {
	return gtx.Repo.Commit()
}

// Tests the repository.GroupRepository interface implementation against common tests.
func GroupRepositoryTestSuite(t *testing.T, gr repository.GroupRepository) {
	guccigang := model.Group{
		Entity: model.Entity{},
		Name:   "guccigang",
	}

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
		_, err := gr.FindByName(guccigang.Name)

		if !errors.Is(repository.ErrNotFound, err) {
			t.Errorf("Expected %v, recieved %v", repository.ErrNotFound, err)
		}
	})

	t.Run("FindByIdErrNotFound", func(t *testing.T) {
		_, err := gr.FindById(1)

		if !errors.Is(repository.ErrNotFound, err) {
			t.Errorf("Expected %v, recieved %v", repository.ErrNotFound, err)
		}
	})

	t.Run("SaveSuccess", func(t *testing.T) {
		err := gr.Create(&guccigang)

		if err != nil {
			t.Errorf("Unexpected error %v", err)
		}

		if guccigang.ID != 1 {
			t.Errorf("Expected ID of 1, got %v", guccigang.ID)
		}
	})

	t.Run("FindByNameSuccess", func(t *testing.T) {
		gg, err := gr.FindByName("guccig")

		if err != nil {
			t.Errorf("Expected no error, recieved %v", err)
		}

		if gg.Name != guccigang.Name {
			t.Errorf("Expected %v, got %v", guccigang, gg)
		}

	})

	t.Run("FindByIdSuccess", func(t *testing.T) {
		gg, err := gr.FindById(1)

		if err != nil {
			t.Errorf("Expected no error, recieved %v", err)
		}

		if gg.ID != guccigang.ID {
			t.Errorf("Expected %v, got %v", guccigang, gg)
		}
	})

	t.Run("SaveNonUniqueNameFail", func(t *testing.T) {
		if err := gr.Create(&guccigang); !errors.Is(repository.ErrUniqueConstraintViolation, err) {
			t.Errorf("Expected %v, got %v", repository.ErrUniqueConstraintViolation, err)
		}

	})

	mnimidamons := &model.Group{
		Entity:       model.Entity{},
		Name:     "mnimidamons",
	}

	grts := GroupRepositoryTransactionSuiteImpl{
		Repo:  gr,
		Group: *mnimidamons,
	}

	runTransactionRollbackSuccessSuite(&grts, t)
	runTransactionCommitSuccessSuite(&grts, t)
}
