package testsuites

import (
	"errors"
	"mnimidamonbackend/domain/model"
	"mnimidamonbackend/domain/repository"
	"testing"
)

type ComputerRepositoryTester struct {
	Repo  repository.ComputerRepository
	URepo repository.UserRepository

	Owner    *model.User
	Computer *model.Computer
}

func (crt *ComputerRepositoryTester) Setup(t *testing.T) {
	err := crt.URepo.Create(crt.Owner)

	if err != nil {
		t.Errorf("Expected owner creation but got error %v", err)
	}
}

func (crt *ComputerRepositoryTester) FindBeforeSaveTests(t *testing.T) {
	cr := crt.Repo
	t.Run("FindAllEmpty", func(t *testing.T) {
		computers, err := cr.FindAll(crt.Owner.ID)

		if err != nil {
			t.Errorf(expectedGot("empty slice", err))
		}

		if len(computers) != 0 {
			t.Errorf(expectedGot("empty slice", computers))
		}
	})

	t.Run("FindByIdNotFound", func(t *testing.T) {
		_, err := cr.FindById(crt.Computer.ID)

		if !errors.Is(repository.ErrNotFound, err) {
			t.Errorf("Expected %v, recieved %v", repository.ErrNotFound, err)
		}
	})

	t.Run("FindByNameNotFound", func(t *testing.T) {
		_, err := cr.FindByName(crt.Computer.Name, crt.Owner.ID)

		if !errors.Is(repository.ErrNotFound, err) {
			t.Errorf("Expected %v, recieved %v", repository.ErrNotFound, err)
		}
	})
}

func (crt *ComputerRepositoryTester) SaveSuccessfulTests(t *testing.T) {
	cr := crt.Repo

	t.Run("SaveSuccess", func(t *testing.T) {
		err := cr.Create(crt.Computer, crt.Owner.ID)

		if err != nil {
			t.Errorf(unexpectedErr(err))
		}

		if crt.Computer.ID == 0 {
			t.Errorf("Expected ID greater 1, got %v", crt.Computer)
		}
	})
}

func (crt *ComputerRepositoryTester) FindAfterSaveTests(t *testing.T) {
	cr := crt.Repo
	t.Run("FindAll", func(t *testing.T) {
		computers, err := cr.FindAll(crt.Owner.ID)

		if err != nil {
			t.Errorf(expectedGot("computers", err))
		}

		if len(computers) == 0 {
			t.Errorf(expectedGot("non empty", computers))
		}
	})

	t.Run("FindById", func(t *testing.T) {
		c, err := cr.FindById(crt.Computer.ID)

		if err != nil {
			t.Errorf(unexpectedErr(err))
		}

		if c.ID != crt.Computer.ID || c.OwnerID != crt.Computer.OwnerID || c.Name != crt.Computer.Name{
			t.Errorf(expectedGot(crt.Computer, c))
		}
	})

	t.Run("FindByName", func(t *testing.T) {
		c, err := cr.FindByName(crt.Computer.Name, crt.Owner.ID)

		if err != nil {
			t.Errorf(unexpectedErr(err))
		}

		if c.ID != crt.Computer.ID || c.OwnerID != crt.Computer.OwnerID || c.Name != crt.Computer.Name{
			t.Errorf(expectedGot(crt.Computer, c))
		}
	})
}

func (crt *ComputerRepositoryTester) ConstraintsTest(t *testing.T) {
	cr := crt.Repo

	t.Run("SaveSameOwnerSameNameFails", func(t *testing.T) {
		err := cr.Create(crt.Computer, crt.Owner.ID)

		if !errors.Is(repository.ErrUniqueConstraintViolation, err) {
			t.Errorf(expectedGot(repository.ErrUniqueConstraintViolation, err))
		}

		computers, err := cr.FindAll(crt.Owner.ID)

		if err != nil {
			t.Errorf(expectedGot("computers", err))
		}

		if len(computers) != 1 {
			t.Errorf(expectedGot("expected 1 computer", computers))
		}
	})
}

func (crt *ComputerRepositoryTester) UpdateTests(t *testing.T) {
	cr := crt.Repo
	c := crt.Computer
	t.Run("UpdateNameSuccess", func(t *testing.T) {
		c.Name = "mac"
		c.OwnerID = 100

		err := cr.Update(c)

		if err != nil {
			t.Errorf(unexpectedErr(err))
		}

		if c.OwnerID != crt.Owner.ID {
			t.Errorf("Should not update ownerID")
		}

		if c.Name != "mac" {
			t.Errorf("Name was not updated")
		}
	})
}

func (crt *ComputerRepositoryTester) SpecificTests(t *testing.T) {
	t.Skip("No specific tests needed")
}

func (crt *ComputerRepositoryTester) DeleteTests(t *testing.T) {
	cr := crt.Repo
	c := crt.Computer

	t.Run("DeleteSuccessful", func(t *testing.T) {
		if err := cr.Delete(c.ID); err != nil {
			t.Error(unexpectedErr(err))
		}
	})

	t.Run("FindByIdFail", func(t *testing.T) {
		if m, err := cr.FindById(c.ID); !errors.Is(repository.ErrNotFound, err) {
			t.Errorf("Expected %v, got err:%v computer:%v", repository.ErrNotFound, err, m)
		}
	})
}

func (crt *ComputerRepositoryTester) BeginTx() TransactionSuiteTestTxInterface {
	crtx := crt.Repo.BeginTx()
	return &ComputerRepositoryTesterTx{
		Repo:     crtx,
		Computer: crt.Computer,
		Owner:    crt.Owner,
	}
}

func (crt *ComputerRepositoryTester) Find() error {
	_, err := crt.Repo.FindById(crt.Computer.ID)
	return err
}

type ComputerRepositoryTesterTx struct {
	Repo     repository.ComputerRepositoryTx
	Computer *model.Computer
	Owner    *model.User
}

func (crtx *ComputerRepositoryTesterTx) Create() error {
	return crtx.Repo.Create(crtx.Computer, crtx.Owner.ID)
}

func (crtx *ComputerRepositoryTesterTx) Find() error {
	_, err := crtx.Repo.FindById(crtx.Computer.ID)
	return err
}

func (crtx *ComputerRepositoryTesterTx) CorrectCheck(t *testing.T) {
	if crtx.Computer.ID == 0 {
		t.Errorf("Expected computer.ID > 0, got %v", crtx.Computer)
	}
}

func (crtx *ComputerRepositoryTesterTx) Rollback() error {
	return crtx.Repo.Rollback()
}

func (crtx *ComputerRepositoryTesterTx) Commit() error {
	return crtx.Repo.Commit()
}

func ComputerRepositoryTestSuite(t *testing.T, cr repository.ComputerRepository, ur repository.UserRepository) {
	marmiha := model.User{
		Entity:       model.Entity{},
		Username:     "marmiha",
		PasswordHash: "marmiha_hashed_pass",
	}

	thinkpad := model.Computer{
		Entity:  model.Entity{},
		OwnerID: 0,
		Name:    "thinkpad",
		Owner:   model.User{},
	}

	brt := &ComputerRepositoryTester{
		Repo:     cr,
		URepo:    ur,
		Owner:    &marmiha,
		Computer: &thinkpad,
	}

	runCommonRepositoryTests(brt, t)
}
