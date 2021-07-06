package testsuites

import (
	"errors"
	"mnimidamonbackend/domain/model"
	"mnimidamonbackend/domain/repository"
	"testing"
)

// TransactionSuiteTestInterface implementation for repository.GroupRepository
type GroupRepositoryTester struct {
	Repo  repository.GroupRepository
	URepo repository.UserRepository
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

	t.Run("ExistFalse", func(t *testing.T) {
		exists, err := gr.Exists(grt.Group.ID)

		if err != nil {
			t.Error(expectedNoError(err))
		}

		if exists {
			t.Error("Expected false, got true")
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

	t.Run("Exists", func(t *testing.T) {
		exists, err := gr.Exists(grt.Group.ID)

		if err != nil {
			t.Error(expectedNoError(err))
		}

		if !exists {
			t.Error("Expected true, got false")
		}
	})
}

func (grt *GroupRepositoryTester) UpdateTests(t *testing.T) {
	group, gr := grt.Group, grt.Repo

	group.Name = "mangodemons"

	if err := gr.Update(&group); err != nil {
		t.Error(expectedNoError(err))
	}

	if group.Name != "mangodemons" {
		t.Error(expectedGot("Group.Name mangodemons", group))
	}

	g, err := gr.FindById(group.ID)
	if err != nil {
		t.Error(expectedNoError(err))
	}

	if g.Name != group.Name {
		t.Error(expectedGot(group, g))
	}
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
	gr, ur := grt.Repo, grt.URepo

	u, g := model.User{
		Entity:       model.Entity{},
		Username:     "membertest",
		PasswordHash: "membertest_hash",
	}, grt.Group

	t.Run("UserIsNotMember", func(t *testing.T) {
		isMember, err := gr.IsMemberOf(1, g.ID)

		if err != nil {
			t.Error(expectedNoError(err))
		}

		if isMember {
			t.Error("Should return false, got true")
		}
	})

	t.Run("UserSavedSuccessfully", func(t *testing.T) {
		if err := ur.Create(&u); err != nil {
			t.Error(expectedGot("no error", err))
		}
	})

	t.Run("UserAddedNonExistentGroupFail", func(t *testing.T) {
		group, err := gr.AddMember(u.ID, 42)

		if err == nil {
			t.Error(expectedGot("an error", group))
		}
	})

	t.Run("AddUnExistingUserToGroup", func(t *testing.T) {
		group, err := gr.AddMember(42, g.ID)

		if err == nil {
			t.Error(expectedGot("an error", group))
		}
	})

	t.Run("UserAddedAsMemberSuccess", func(t *testing.T) {
		group, err := gr.AddMember(u.ID, g.ID)

		if err != nil {
			t.Error(expectedNoError(err))
		}

		if group.ID != g.ID && group.Name == g.Name	{
			t.Error(expectedGot(u, group))
		}
	})

	t.Run("UserAddedAsMemberSecondTimeFail", func(t *testing.T) {
		_, err := gr.AddMember(u.ID, g.ID)

		if !errors.Is(repository.ErrUserAlreadyInGroupViolation, err) {
			t.Error(expectedGot(repository.ErrUserAlreadyInGroupViolation, err))
		}
	})

	t.Run("UserIsMember", func(t *testing.T) {
		isMember, err := gr.IsMemberOf(u.ID, g.ID)

		if err != nil {
			t.Error(expectedNoError(err))
		}
		if !isMember {
			t.Error("Should return true, got false")
		}
	})


}

func (grt *GroupRepositoryTester) DeleteTests(t *testing.T) {
	g, gr := grt.Group, grt.Repo

	t.Run("DeleteSuccessful", func(t *testing.T) {
		if err := gr.Delete(g.ID); err != nil {
			t.Error(expectedNoError(err))
		}
	})

	t.Run("FindByIdFail", func(t *testing.T) {
		if m, err := gr.FindById(g.ID); !errors.Is(repository.ErrNotFound, err) {
			t.Errorf("Expected %v, got err:%v group:%v", repository.ErrNotFound, err, m)
		}
	})
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
	Repo  repository.GroupRepositoryTx
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
func GroupRepositoryTestSuite(t *testing.T, gr repository.GroupRepository, ur repository.UserRepository) {
	guccigang, mnimidamons := model.Group{
		Entity: model.Entity{},
		Name:   "guccigang",
	}, model.Group{
		Entity: model.Entity{},
		Name:   "mnimidamons",
	}

	grt := &GroupRepositoryTester{
		Repo:  gr,
		URepo: ur,
		Group: guccigang,
	}

	runCommonRepositoryTests(grt, t)

	grt.Group = mnimidamons
	runTransactionRollbackSuccessSuite(grt, t)
}
