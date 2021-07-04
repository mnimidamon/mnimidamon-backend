package testsuites

import (
	"errors"
	"mnimidamonbackend/domain/model"
	"mnimidamonbackend/domain/repository"
	"testing"
)

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

	t.Run("TransactionRollbackSuccess", func(t *testing.T) {
		grx := gr.BeginTx()

		if err := grx.Create(mnimidamons); err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if mnimidamons.ID == 0 {
			t.Errorf("Expected mnimidamons.ID of 0, got %v", mnimidamons.ID)
		}

		if err := grx.Rollback(); err != nil {
			t.Errorf("Expected no error on rollback, got %v", err)
		}

		if _, err := grx.FindByName(mnimidamons.Name); !errors.Is(repository.ErrTxAlreadyRolledBack, err) {
			t.Errorf("Expected %v, recieved %v", repository.ErrTxAlreadyRolledBack, err)
		}

		if _, err := gr.FindByName(mnimidamons.Name); !errors.Is(repository.ErrNotFound, err) {
			t.Errorf("Expected %v, got %v", repository.ErrNotFound, err)
		}
	})

	t.Run("TransactionCommitSuccess", func(t *testing.T) {
		grx := gr.BeginTx()

		if err := grx.Create(mnimidamons); err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if mnimidamons.ID == 0 {
			t.Errorf("Expected mnimidamons.ID of 0, got %v", mnimidamons.ID)
		}

		if err := grx.Commit(); err != nil {
			t.Errorf("Expected no error on rollback, got %v", err)
		}

		if _, err := grx.FindByName(mnimidamons.Name); !errors.Is(repository.ErrTxAlreadyRolledBack, err) {
			t.Errorf("Expected %v, recieved %v", repository.ErrTxAlreadyRolledBack, err)
		}

		if _, err := gr.FindByName(mnimidamons.Name); err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
	})
}
