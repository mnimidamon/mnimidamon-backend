package testsuites

import (
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
	panic("implement me")
}

func (grct *GroupComputerRepositoryTester) SaveSuccessfulTests(t *testing.T) {
	panic("implement me")
}

func (grct *GroupComputerRepositoryTester) FindAfterSaveTests(t *testing.T) {
	panic("implement me")
}

func (grct *GroupComputerRepositoryTester) ConstraintsTest(t *testing.T) {
	panic("implement me")
}

func (grct *GroupComputerRepositoryTester) UpdateTests(t *testing.T) {
	panic("implement me")
}

func (grct *GroupComputerRepositoryTester) SpecificTests(t *testing.T) {
	panic("implement me")
}

func (grct *GroupComputerRepositoryTester) DeleteTests(t *testing.T) {
	panic("implement me")
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
		Owner:   model.User{},
	}

	mnimidamons_thinkpad := &model.GroupComputer{
		Entity:      model.Entity{},
		GroupID:     0,
		ComputerID:  0,
		Group:       model.Group{},
		Computer:    model.Computer{},
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
