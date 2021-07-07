package testsuites

import (
	"errors"
	"mnimidamonbackend/domain/model"
	"mnimidamonbackend/domain/repository"
	"testing"
)

// TransactionSuiteTestInterface and CommonRepositoryTestSuiteInterface implementation for repository.UserRepository
type UserRepositoryTester struct {
	Repo repository.UserRepository
	User model.User
}

func (urt *UserRepositoryTester) Setup(t *testing.T) {
	t.Skip("No specific setup needed")
}

func (urt *UserRepositoryTester) FindBeforeSaveTests(t *testing.T) {
	ur := urt.Repo
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
		_, err := ur.FindByUsername(urt.User.Username[0:2])

		if !errors.Is(repository.ErrNotFound, err) {
			t.Errorf("Expected %v, recieved %v", repository.ErrNotFound, err)
		}
	})

	t.Run("FindByIdErrNotFound", func(t *testing.T) {
		_, err := ur.FindById(urt.User.ID)

		if !errors.Is(repository.ErrNotFound, err) {
			t.Errorf("Expected %v, recieved %v", repository.ErrNotFound, err)
		}
	})

	t.Run("ExistFalse", func(t *testing.T) {
		exists, err := ur.Exists(1)

		if err != nil {
			t.Error(unexpectedErr(err))
		}

		if exists {
			t.Error("Expected false, got true")
		}
	})
}

func (urt *UserRepositoryTester) SaveSuccessfulTests(t *testing.T) {
	ur := urt.Repo
	t.Run("SaveSuccess", func(t *testing.T) {
		err := ur.Create(&urt.User)

		if err != nil {
			t.Errorf("Unexpected error %v", err)
		}

		if urt.User.ID == 0 {
			t.Errorf("Expected ID greater than 0, got %v", urt.User.ID)
		}
	})
}

func (urt *UserRepositoryTester) FindAfterSaveTests(t *testing.T) {
	ur := urt.Repo

	t.Run("FindByUsernameSuccess", func(t *testing.T) {
		m, err := ur.FindByUsername(urt.User.Username[0:2])

		if err != nil {
			t.Errorf("Expected no error, recieved %v", err)
		}

		if m.Username != urt.User.Username {
			t.Errorf("Expected %v, got %v", urt.User, m)
		}

	})

	t.Run("FindByIdSuccess", func(t *testing.T) {
		m, err := ur.FindById(urt.User.ID)

		if err != nil {
			t.Errorf("Expected no error, recieved %v", err)
		}

		if m.ID != urt.User.ID {
			t.Errorf("Expected %v, got %v", urt.User.ID, m)
		}
	})

	t.Run("Exists", func(t *testing.T) {
		exists, err := ur.Exists(urt.User.ID)

		if err != nil {
			t.Error(unexpectedErr(err))
		}

		if !exists {
			t.Error("Expected true, got false")
		}
	})
}

func (urt *UserRepositoryTester) UpdateTests(t *testing.T) {
	user, ur := urt.User, urt.Repo

	user.Username = "markez"
	user.PasswordHash = "new_password_hash"

	if err := ur.Update(&user); err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if user.Username != "markez" {
		t.Errorf("Expected user.Username == markez, got %v", user)
	}

	if user.PasswordHash != "new_password_hash" {
		t.Errorf("Expected user.PasswordHash == new_password_hash, got %v", user)
	}

	u, err := ur.FindById(user.ID)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if u.Username != user.Username || u.PasswordHash != user.PasswordHash {
		t.Errorf("Expected %v, got %v", user, u)
	}
}

func (urt *UserRepositoryTester) ConstraintsTest(t *testing.T) {
	ur := urt.Repo
	t.Run("SaveNonUniqueNameFail", func(t *testing.T) {
		if err := ur.Create(&urt.User); !errors.Is(repository.ErrUniqueConstraintViolation, err) {
			t.Errorf("Expected %v, got %v", repository.ErrUniqueConstraintViolation, err)
		}
	})
}

func (urt *UserRepositoryTester) SpecificTests(t *testing.T) {
	t.Skip("No specific tests")
}

func (urt *UserRepositoryTester) DeleteTests(t *testing.T) {
	ur := urt.Repo

	t.Run("DeleteSuccessful", func(t *testing.T) {
		if err := ur.Delete(urt.User.ID); err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
	})

	t.Run("FindByIdFail", func(t *testing.T) {
		if m, err := ur.FindById(urt.User.ID); !errors.Is(repository.ErrNotFound, err) {
			t.Errorf("Expected %v, got err:%v user:%v", repository.ErrNotFound, err, m)
		}
	})
}

func (urt *UserRepositoryTester) BeginTx() TransactionSuiteTestTxInterface {
	utx := urt.Repo.BeginTx()
	return &UserRepositoryTesterTx{
		Repo: utx,
		User: urt.User,
	}
}

func (urt *UserRepositoryTester) Find() error {
	_, err := urt.Repo.FindByUsername(urt.User.Username)
	return err
}

type UserRepositoryTesterTx struct {
	Repo repository.UserRepositoryTx
	User model.User
}

func (utx *UserRepositoryTesterTx) Rollback() error {
	return utx.Repo.Rollback()
}

func (utx *UserRepositoryTesterTx) Commit() error {
	return utx.Repo.Commit()
}

func (utx *UserRepositoryTesterTx) Create() error {
	return utx.Repo.Create(&utx.User)
}

func (utx *UserRepositoryTesterTx) Find() error {
	_, err := utx.Repo.FindByUsername(utx.User.Username)
	return err
}

func (utx *UserRepositoryTesterTx) CorrectCheck(t *testing.T) {
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

	urt := &UserRepositoryTester{
		Repo: ur,
		User: marmiha,
	}

	// Common repository implementation testing.
	runCommonRepositoryTests(urt, t)
}
