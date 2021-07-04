package testsuites

import (
	"errors"
	"mnimidamonbackend/domain/model"
	"mnimidamonbackend/domain/repository"
	"testing"
)

// Tests the repository.UserRepository interface implementation against common tests.
func UserRepositoryTestSuite(t *testing.T, ur repository.UserRepository) {
	marmiha := model.User{
		Entity:       model.Entity{},
		Username:     "marmiha",
		PasswordHash: "marmiha_hashed_pass",
	}

	t.Run("FindAllEmpty", func(t *testing.T) {
		users, err := ur.FindAll()

		if err != nil {
			t.Errorf("Expected empty slice, got an error %v", err)
		}

		if len(users) != 0 {
			t.Errorf("Expected empty slice, recieved %v", users)
		}
	})

	t.Run("FindByUsernameErrNotFound", func(t *testing.T) {
		_, err := ur.FindByUsername("marmiha")

		if !errors.Is(repository.ErrNotFound, err) {
			t.Errorf("Expected %v, recieved %v", repository.ErrNotFound, err)
		}
	})

	t.Run("FindByIdErrNotFound", func(t *testing.T) {
		_, err := ur.FindById(1)

		if !errors.Is(repository.ErrNotFound, err) {
			t.Errorf("Expected %v, recieved %v", repository.ErrNotFound, err)
		}
	})

	t.Run("SaveSuccess", func(t *testing.T) {
		err := ur.Create(&marmiha)

		if err != nil {
			t.Errorf("Unexpected error %v", err)
		}

		if marmiha.ID != 1 {
			t.Errorf("Expected ID of 1, got %v", marmiha.ID)
		}
	})

	t.Run("FindByUsernameSuccess", func(t *testing.T) {
		m, err := ur.FindByUsername("marm")

		if err != nil {
			t.Errorf("Expected no error, recieved %v", err)
		}

		if m.Username != marmiha.Username {
			t.Errorf("Expected %v, got %v", marmiha, m)
		}

	})

	t.Run("FindByIdSuccess", func(t *testing.T) {
		m, err := ur.FindById(1)

		if err != nil {
			t.Errorf("Expected no error, recieved %v", err)
		}

		if m.ID != marmiha.ID {
			t.Errorf("Expected %v, got %v", marmiha, m)
		}
	})

	t.Run("SaveNonUniqueNameFail", func(t *testing.T) {
		if err := ur.Create(&marmiha); !errors.Is(repository.ErrUniqueConstraintViolation, err) {
			t.Errorf("Expected %v, got %v", repository.ErrUniqueConstraintViolation, err)
		}

	})

	peter := &model.User{
		Entity:       model.Entity{},
		Username:     "peter",
		PasswordHash: "peters_hashed_pass",
	}

	t.Run("TransactionRollbackSuccess", func(t *testing.T) {
		urx := ur.BeginTx()

		if err := urx.Create(peter); err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if peter.ID == 0 {
			t.Errorf("Expected peter.ID of 0, got %v", peter.ID)
		}

		if err := urx.Rollback(); err != nil {
			t.Errorf("Expected no error on rollback, got %v", err)
		}

		if _, err := urx.FindByUsername(peter.Username); !errors.Is(repository.ErrTxAlreadyRolledBack, err) {
			t.Errorf("Expected %v, recieved %v", repository.ErrTxAlreadyRolledBack, err)
		}

		if _, err := ur.FindByUsername(peter.Username); !errors.Is(repository.ErrNotFound, err) {
			t.Errorf("Expected %v, got %v", repository.ErrNotFound, err)
		}
	})

	t.Run("TransactionCommitSuccess", func(t *testing.T) {
		urx := ur.BeginTx()

		if err := urx.Create(peter); err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if peter.ID == 0 {
			t.Errorf("Expected peter.ID of 0, got %v", peter.ID)
		}

		if err := urx.Commit(); err != nil {
			t.Errorf("Expected no error on rollback, got %v", err)
		}

		if _, err := urx.FindByUsername(peter.Username); !errors.Is(repository.ErrTxAlreadyRolledBack, err) {
			t.Errorf("Expected %v, recieved %v", repository.ErrTxAlreadyRolledBack, err)
		}

		if _, err := ur.FindByUsername(peter.Username); err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
	})
}
