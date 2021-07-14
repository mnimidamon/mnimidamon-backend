package testsuites

import (
	"errors"
	"fmt"
	"mnimidamonbackend/domain/model"
	"mnimidamonbackend/domain/repository"
	"testing"
	"time"
)

type InviteRepositoryTester struct {
	Repo repository.InviteRepository
	GRepo repository.GroupRepository
	URepo repository.UserRepository

	Invite *model.Invite
	User *model.User
	Group *model.Group
}

func (irt *InviteRepositoryTester) Setup(t *testing.T) {
	_, gr, ur := irt.Repo, irt.GRepo, irt.URepo
	_, u, g := irt.Invite, irt.User, irt.Group

	t.Run("PreSaveOperations", func(t *testing.T) {
		err := ur.Create(u)
		if err != nil {
			t.Error(expectedGot("no error upon creating user", err))
		}

		err = gr.Create(g)
		if err != nil {
			t.Error(expectedGot("no error upon creating group", err))
		}
	})
}

func (irt *InviteRepositoryTester) FindBeforeSaveTests(t *testing.T) {
	_, u, g := irt.Invite, irt.User, irt.Group
	ir := irt.Repo
	t.Run("FindByIdFail", func(t *testing.T) {
		i , err := ir.FindById(u.ID, g.ID)

		if !errors.Is(repository.ErrNotFound, err) {
			t.Errorf(expectedGot(repository.ErrNotFound, fmt.Sprintf("Error: %v, Invite: %v", err, i)))
		}
	})

	t.Run("FindAllOfGroupEmpty", func(t *testing.T) {
		invites , err := ir.FindAllOfGroup(g.ID)

		if err != nil {
			t.Errorf(expectedGot("empty slice", err))
		}

		if len(invites) != 0 {
			t.Errorf(expectedGot("empty slice", invites))
		}
	})

	t.Run("FindAllOfUserEmpty", func(t *testing.T) {
		invites , err := ir.FindAllOfUser(u.ID)

		if err != nil {
			t.Errorf(expectedGot("empty slice", err))
		}

		if len(invites) != 0 {
			t.Errorf(expectedGot("empty slice", invites))
		}
	})

	t.Run("ExistsNot", func(t *testing.T) {
		exists, err := ir.Exists(u.ID, g.ID)

		if err != nil {
			t.Error(unexpectedErr(err))
		}

		if exists {
			t.Error("Expected false, got true")
		}
	})
}

func (irt *InviteRepositoryTester) SaveSuccessfulTests(t *testing.T) {
	i, u, g := irt.Invite, irt.User, irt.Group
	ir := irt.Repo

	t.Run("SaveSuccess", func(t *testing.T) {
		i.GroupID = g.ID
		i.UserID = u.ID

		if err := ir.Create(i); err != nil {
			t.Error(unexpectedErr(err))
		}

		if i.UserID != u.ID {
			t.Error(expectedGot("right user_id", i))
		}

		if i.GroupID != g.ID {
			t.Error(expectedGot("right group_id", i))
		}
	})

	t.Run("Exists", func(t *testing.T) {
		exists, err := ir.Exists(u.ID, g.ID)

		if err != nil {
			t.Error(unexpectedErr(err))
		}

		if !exists {
			t.Error("Expected true, got false")
		}
	})
}

func (irt *InviteRepositoryTester) FindAfterSaveTests(t *testing.T) {
	_, u, g := irt.Invite, irt.User, irt.Group
	ir := irt.Repo
	t.Run("FindByIdSuccess", func(t *testing.T) {
		i , err := ir.FindById(u.ID, g.ID)

		if err != nil {
			t.Error(unexpectedErr(err))
		}

		if i.UserID != u.ID {
			t.Error(expectedGot("right user_id", i))
		}

		if i.GroupID != g.ID {
			t.Error(expectedGot("right group_id", i))
		}
	})

	t.Run("FindAllOfGroupSuccess", func(t *testing.T) {
		invites , err := ir.FindAllOfGroup(g.ID)

		if err != nil {
			t.Errorf(expectedGot("slice", err))
		}

		if len(invites) != 1 {
			t.Errorf(expectedGot("slice of 1", invites))
		}
	})

	t.Run("FindAllOfUserSuccess", func(t *testing.T) {
		invites , err := ir.FindAllOfUser(u.ID)

		if err != nil {
			t.Errorf(expectedGot("slice of 1", err))
		}

		if len(invites) != 1 {
			t.Errorf(expectedGot("empty slice", invites))
		}
	})
}

func (irt *InviteRepositoryTester) ConstraintsTest(t *testing.T) {
	i := irt.Invite
	ir := irt.Repo
	t.Run("DuplicateInviteSaveFail", func(t *testing.T) {
		if err := ir.Create(i); !errors.Is(repository.ErrAlreadyExists, err) {
			t.Error(expectedGot(repository.ErrAlreadyExists, err))
		} else if err != nil && !errors.Is(repository.ErrAlreadyExists, err){
			t.Error(unexpectedErr(err))
		}
	})
}

func (irt *InviteRepositoryTester) UpdateTests(t *testing.T) {
	t.Skip("No specific update tests")
}

func (irt *InviteRepositoryTester) SpecificTests(t *testing.T) {
	t.Skip("No specific tests")
}

func (irt *InviteRepositoryTester) DeleteTests(t *testing.T) {
	_, u, g := irt.Invite, irt.User, irt.Group
	ir := irt.Repo

	t.Run("DeleteSuccess", func(t *testing.T) {
		err := ir.Delete(u.ID, g.ID)

		if err != nil {
			t.Error(unexpectedErr(err))
		}
	})

	t.Run("FindByIdFail", func(t *testing.T) {
		i , err := ir.FindById(u.ID, g.ID)

		if !errors.Is(repository.ErrNotFound, err) {
			t.Errorf(expectedGot(repository.ErrNotFound, fmt.Sprintf("Error: %v, Invite: %v", err, i)))
		}
	})
}

func (irt *InviteRepositoryTester) BeginTx() TransactionSuiteTestTxInterface {
	irtx := irt.Repo.BeginTx()
	return &InviteRepositoryTesterTx{
		Repo:   irtx,
		Invite: irt.Invite,
		User:   irt.User,
		Group:  irt.Group,
	}
}

func (irt *InviteRepositoryTester) Find() error {
	_, err := irt.Repo.FindById(irt.User.ID, irt.Group.ID)
	return err
}

type InviteRepositoryTesterTx struct {
	Repo repository.InviteRepositoryTx
	Invite *model.Invite
	User *model.User
	Group *model.Group
}

func (irtx *InviteRepositoryTesterTx) Create() error {
	return irtx.Repo.Create(irtx.Invite)
}

func (irtx *InviteRepositoryTesterTx) Find() error {
	_, err := irtx.Repo.FindById(irtx.User.ID, irtx.Group.ID)
	return err
}

func (irtx *InviteRepositoryTesterTx) CorrectCheck(t *testing.T) {
	i, u, g := irtx.Invite, irtx.User, irtx.Group
	if !(i.UserID == u.ID && i.GroupID == g.ID) {
		t.Error(expectedGot("matching id's of invite group and user", i))
	}
}

func (irtx *InviteRepositoryTesterTx) Rollback() error {
	return irtx.Repo.Rollback()
}

func (irtx *InviteRepositoryTesterTx) Commit() error {
	return irtx.Repo.Commit()
}

func InviteRepositoryTestSuite(t *testing.T, ir repository.InviteRepository, gr repository.GroupRepository, ur repository.UserRepository) {
	invite := &model.Invite{
		UserID:    0,
		GroupID:   0,
		User:      nil,
		Group:     nil,
		CreatedAt: time.Time{},
	}

	marmiha := &model.User{
		Entity:       model.Entity{},
		Username:     "marmiha",
		PasswordHash: "marmiha_hash",
	}

	mnimidamons := &model.Group{
		Entity: model.Entity{},
		Name:   "mnimidamons",
	}
	
	irt := &InviteRepositoryTester{
		Repo:   ir,
		GRepo:  gr,
		URepo:  ur,
		Invite: invite,
		User:   marmiha,
		Group:  mnimidamons,
	}

	runCommonRepositoryTests(irt, t)
}