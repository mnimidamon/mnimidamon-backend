package testsuites

import (
	"errors"
	"mnimidamonbackend/domain/model"
	"mnimidamonbackend/domain/repository"
	"testing"
)

// TransactionSuiteTestInterface implementation for repository.UserRepository
type UserRepositoryTransactionSuiteImpl struct {
	Repo repository.UserRepository
	User model.User
}

func (u *UserRepositoryTransactionSuiteImpl) BeginTx() TransactionSuiteTestTxInterface {
	utx := u.Repo.BeginTx()
	return &UserRepositoryTransactionSuiteTxImpl{
		Repo: utx,
		User: u.User,
	}
}

func (u *UserRepositoryTransactionSuiteImpl) Find() error {
	_, err := u.Repo.FindByUsername(u.User.Username)
	return err
}

type UserRepositoryTransactionSuiteTxImpl struct {
	Repo repository.UserRepositoryTx
	User model.User
}

func (utx *UserRepositoryTransactionSuiteTxImpl) Rollback() error {
	return utx.Repo.Rollback()
}

func (utx *UserRepositoryTransactionSuiteTxImpl) Commit() error {
	return utx.Repo.Commit()
}

func (utx *UserRepositoryTransactionSuiteTxImpl) Create() error {
	return utx.Repo.Create(&utx.User)
}

func (utx *UserRepositoryTransactionSuiteTxImpl) Find() error {
	_, err := utx.Repo.FindByUsername(utx.User.Username)
	return err
}

func (utx *UserRepositoryTransactionSuiteTxImpl) CorrectCheck(t *testing.T) {
	if utx.User.ID == 0 {
		t.Errorf("Expected user.ID > 0, got %v", utx.User)
	}
}

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

	t.Run("DeleteSuccessful", func(t *testing.T) {
		if err := ur.Delete(&marmiha); err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		t.Run("FindByIdFail", func(t *testing.T) {
			if m, err := ur.FindById(marmiha.ID); !errors.Is(repository.ErrNotFound, err) {
				t.Errorf("Expected %v, got err:%v user:%v", repository.ErrNotFound, err, m)
			}
		})
	})

	peter := &model.User{
		Entity:       model.Entity{},
		Username:     "peter",
		PasswordHash: "peters_hashed_pass",
	}

	// Common transaction implementation testing.
	urts := UserRepositoryTransactionSuiteImpl{
		Repo: ur,
		User: *peter,
	}

	runTransactionRollbackSuccessSuite(&urts, t)
	runTransactionCommitSuccessSuite(&urts, t)
}
