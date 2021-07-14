package testsuites

import (
	"errors"
	"fmt"
	"mnimidamonbackend/domain/model"
	"mnimidamonbackend/domain/repository"
	"testing"
)

type GroupComputerRepositoryTester struct {
	Repo  repository.GroupComputerRepository
	GRepo repository.GroupRepository
	URepo repository.UserRepository
	CRepo repository.ComputerRepository

	GroupComputer *model.GroupComputer
	User          *model.User
	Group         *model.Group
	Computer      *model.Computer
}

func (grct *GroupComputerRepositoryTester) Setup(t *testing.T) {
	// Inserting user, his computer and the group that the user will belong to.
	_, gr, ur, cr := grct.Repo, grct.GRepo, grct.URepo, grct.CRepo
	_, u, g, c := grct.GroupComputer, grct.User, grct.Group, grct.Computer
	/*
	gcr, gr, ur, cr := grct.Repo, grct.GRepo, grct.URepo, grct.CRepo
	gc, u, g, c := grct.GroupComputer, grct.User, grct.Group, grct.Computer
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
	})
}

func (grct *GroupComputerRepositoryTester) FindBeforeSaveTests(t *testing.T) {
	gcr := grct.Repo
	g, c := grct.Group, grct.Computer

	t.Run("FindByIdFail", func(t *testing.T) {
		gc, err := gcr.FindById(g.ID, c.ID)

		if !errors.Is(repository.ErrNotFound, err) {
			t.Errorf(expectedGot(repository.ErrNotFound, fmt.Sprintf("Error: %v, Invite: %v", err, gc)))
		}
	})

	t.Run("FindAllOfGroupEmpty", func(t *testing.T) {
		computers, err := gcr.FindAllOfGroup(g.ID)

		if err != nil {
			t.Errorf(expectedGot("empty slice", err))
		}

		if len(computers) != 0 {
			t.Errorf(expectedGot("empty slice", computers))
		}
	})

	t.Run("FindAllOfUserEmpty", func(t *testing.T) {
		computers, err := gcr.FindAllOfComputer(c.ID)

		if err != nil {
			t.Errorf(expectedGot("empty slice", err))
		}

		if len(computers) != 0 {
			t.Errorf(expectedGot("empty slice", computers))
		}
	})

	t.Run("ExistsNot", func(t *testing.T) {
		exists, err := gcr.Exists(g.ID, c.ID)

		if err != nil {
			t.Error(unexpectedErr(err))
		}

		if exists {
			t.Error("Expected false, got true")
		}
	})
}

func (grct *GroupComputerRepositoryTester) SaveSuccessfulTests(t *testing.T) {
	gcr := grct.Repo
	gc, g, c := grct.GroupComputer, grct.Group, grct.Computer

	t.Run("SaveSuccess", func(t *testing.T) {
		gc.GroupID = g.ID
		gc.ComputerID = c.ID

		if err := gcr.Create(gc); err != nil {
			t.Error(unexpectedErr(err))
		}

		if gc.GroupID != g.ID {
			t.Error(expectedGot("right group_id", gc))
		}

		if gc.ComputerID != c.ID {
			t.Error(expectedGot("right computer_id", gc))
		}
	})

	t.Run("Exists", func(t *testing.T) {
		exists, err := gcr.Exists(g.ID, c.ID)

		if err != nil {
			t.Error(unexpectedErr(err))
		}

		if !exists {
			t.Error("Expected true, got false")
		}
	})
}

func (grct *GroupComputerRepositoryTester) FindAfterSaveTests(t *testing.T) {
	gcr := grct.Repo
	_, g, c := grct.GroupComputer, grct.Group, grct.Computer

	t.Run("FindByIdSuccess", func(t *testing.T) {
		cn, err := gcr.FindById(g.ID, c.ID)

		if err != nil {
			t.Error(unexpectedErr(err))
		}

		if cn.ComputerID != c.ID {
			t.Error(expectedGot("right computer_id", cn))
		}

		if cn.GroupID != g.ID {
			t.Error(expectedGot("right group_id", cn))
		}
	})

	t.Run("FindAllOfGroupSuccess", func(t *testing.T) {
		computers, err := gcr.FindAllOfGroup(g.ID)

		if err != nil {
			t.Errorf(expectedGot("slice", err))
		}

		if len(computers) != 1 {
			t.Errorf(expectedGot("slice of 1", computers))
		}
	})

	t.Run("FindAllOfUserSuccess", func(t *testing.T) {
		computers, err := gcr.FindAllOfComputer(c.ID)

		if err != nil {
			t.Errorf(expectedGot("slice", err))
		}

		if len(computers) != 1 {
			t.Errorf(expectedGot("slice of 1", computers))
		}
	})
}

func (grct *GroupComputerRepositoryTester) ConstraintsTest(t *testing.T) {
	gcr := grct.Repo
	gc :=  grct.GroupComputer

	t.Run("DuplicateInviteSaveFail", func(t *testing.T) {
		if err := gcr.Create(gc); !errors.Is(repository.ErrAlreadyExists, err) {
			t.Error(expectedGot(repository.ErrAlreadyExists, err))
		} else if err != nil && !errors.Is(repository.ErrAlreadyExists, err){
			t.Error(unexpectedErr(err))
		}
	})
}

func (grct *GroupComputerRepositoryTester) UpdateTests(t *testing.T) {
	gcr := grct.Repo
	cn, g, c := grct.GroupComputer, grct.Group, grct.Computer

	t.Run("UpdateSize", func(t *testing.T) {
		cn.ComputerID = 200
		cn.GroupID = 10

		cn.StorageSize = 200

		if err := gcr.Update(cn); err != nil {
			t.Error(unexpectedErr(err))
		}

		if cn.ComputerID != c.ID {
			t.Error(expectedGot("right computer_id", cn))
		}

		if cn.GroupID != g.ID {
			t.Error(expectedGot("right group_id", cn))
		}

		if cn.StorageSize != 200 {
			t.Error(expectedGot("storage size of 200", cn))
		}
	})
}

func (grct *GroupComputerRepositoryTester) SpecificTests(t *testing.T) {
	t.Skip("No specific tests")
}

func (grct *GroupComputerRepositoryTester) DeleteTests(t *testing.T) {
	gcr := grct.Repo
	cn := grct.GroupComputer

	t.Run("DeleteSuccess", func(t *testing.T) {
		err := gcr.Delete(cn.GroupID, cn.ComputerID)

		if err != nil {
			t.Error(unexpectedErr(err))
		}
	})

	t.Run("FindByIdFail", func(t *testing.T) {
		ggg, err := gcr.FindById(cn.GroupID, cn.ComputerID)

		if !errors.Is(repository.ErrNotFound, err) {
			t.Errorf(expectedGot(repository.ErrNotFound, fmt.Sprintf("Error: %v, GroupComputer: %v", err, ggg)))
		}
	})
}

func (grct *GroupComputerRepositoryTester) BeginTx() TransactionSuiteTestTxInterface {
	return &GroupComputerRepositoryTesterTx{
		Repo:          grct.Repo.BeginTx(),
		GroupComputer: grct.GroupComputer,
		User:          grct.User,
		Group:         grct.Group,
		Computer:      grct.Computer,
	}
}

func (grct *GroupComputerRepositoryTester) Find() error {
	_, err := grct.Repo.FindById(grct.Group.ID, grct.Computer.ID)
	return err
}

type GroupComputerRepositoryTesterTx struct {
	Repo repository.GroupComputerRepositoryTx

	GroupComputer *model.GroupComputer
	User          *model.User
	Group         *model.Group
	Computer      *model.Computer
}

func (grctx *GroupComputerRepositoryTesterTx) Create() error {
	return grctx.Repo.Create(grctx.GroupComputer)
}

func (grctx *GroupComputerRepositoryTesterTx) Find() error {
	_, err := grctx.Repo.FindById(grctx.Group.ID, grctx.Computer.ID)
	return err
}

func (grctx *GroupComputerRepositoryTesterTx) CorrectCheck(t *testing.T) {
	gc, c, g := grctx.GroupComputer, grctx.Computer, grctx.Group
	if !(gc.GroupID == g.ID && gc.ComputerID == c.ID) {
		t.Error(expectedGot("matching id's of group and computer", gc))
	}
}

func (grctx *GroupComputerRepositoryTesterTx) Rollback() error {
	return grctx.Repo.Rollback()
}

func (grctx *GroupComputerRepositoryTesterTx) Commit() error {
	return grctx.Repo.Commit()
}

func GroupComputerRepositoryTestSuite(t *testing.T, gcr repository.GroupComputerRepository, gr repository.GroupRepository, ur repository.UserRepository, cr repository.ComputerRepository) {
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

	gcrt := &GroupComputerRepositoryTester{
		Repo:          gcr,
		GRepo:         gr,
		URepo:         ur,
		CRepo:         cr,
		GroupComputer: mnimidamons_thinkpad,
		User:          marmiha,
		Group:         mnimidamons,
		Computer:      thinkpad,
	}

	runCommonRepositoryTests(gcrt, t)
}
