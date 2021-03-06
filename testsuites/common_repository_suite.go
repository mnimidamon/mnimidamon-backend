package testsuites

import (
	"errors"
	"fmt"
	"mnimidamonbackend/domain/repository"
	"testing"
)

// Logging the unimplemented test suites.
var unimplemented = "Unimplemented Tests"

// For testing common repository procedures. For more consistency and easier code consistency.
type CommonRepositoryTestSuiteInterface interface {
	Setup(t *testing.T)               // Setting up for the testing, repository initializations, dependent models insertion.
	FindBeforeSaveTests(t *testing.T) // Find functionalities testing before save.
	SaveSuccessfulTests(t *testing.T) // Successful saving tests.
	FindAfterSaveTests(t *testing.T)  // Finding after successful save.
	ConstraintsTest(t *testing.T)     // Repository model specific constraint testing.
	UpdateTests(t *testing.T)         // Updating tests.
	SpecificTests(t *testing.T)       // Repository specific tests.
	DeleteTests(t *testing.T)         // Deletion testing.

	TransactionSuiteTestInterface
}

// For testing the common Transaction implementation.
type TransactionSuiteTestInterface interface {
	BeginTx() TransactionSuiteTestTxInterface // Begin a transaction.
	Find() error                              // Find something outside of the transaction.
}

// For testing the common Transaction implementation when transaction has already begun.
type TransactionSuiteTestTxInterface interface {
	Create() error             // Create something inside a transaction.
	Find() error               // Find something inside the transaction.
	CorrectCheck(t *testing.T) // Check the correctness of the found thing.
	repository.Transaction     // Rollback and Commit functionalities.
}

// Run common repository testing suite.
func runCommonRepositoryTests(crtsi CommonRepositoryTestSuiteInterface, t *testing.T) {
	t.Run("TestingSuiteSetup", func(t *testing.T) {
		crtsi.Setup(t)
	})

	t.Run("FindBeforeSaveTests", func(t *testing.T) {
		crtsi.FindBeforeSaveTests(t)
	})

	t.Run("SaveSuccessfulTests", func(t *testing.T) {
		crtsi.SaveSuccessfulTests(t)
	})

	t.Run("FindAfterSaveTests", func(t *testing.T) {
		crtsi.FindAfterSaveTests(t)
	})

	t.Run("ConstraintsTest", func(t *testing.T) {
		crtsi.ConstraintsTest(t)
	})

	t.Run("UpdateTests", func(t *testing.T) {
		crtsi.UpdateTests(t)
	})

	t.Run("SpecificTests", func(t *testing.T) {
		crtsi.SpecificTests(t)
	})

	t.Run("DeleteTests", func(t *testing.T) {
		crtsi.DeleteTests(t)
	})

	t.Run("TransactionTests", func(t *testing.T) {
		runTransactionTestSuite(crtsi, t)
	})
}

func runTransactionTestSuite(ti TransactionSuiteTestInterface, t *testing.T) {
	runTransactionRollbackSuccessSuite(ti, t)
	runTransactionCommitSuccessSuite(ti, t)
}

func runTransactionRollbackSuccessSuite(ti TransactionSuiteTestInterface, t *testing.T) {
	t.Run("TransactionRollbackSuccess", func(t *testing.T) {
		tix := ti.BeginTx()

		if err := tix.Create(); err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		tix.CorrectCheck(t)

		if err := tix.Rollback(); err != nil {
			t.Errorf("Expected no error on rollback, got %v", err)
		}

		if err := tix.Find(); !errors.Is(repository.ErrTxAlreadyRolledBack, err) {
			t.Errorf("Expected %v, recieved %v", repository.ErrTxAlreadyRolledBack, err)
		}

		if err := ti.Find(); !errors.Is(repository.ErrNotFound, err) {
			t.Errorf("Expected %v, got %v", repository.ErrNotFound, err)
		}
	})
}

func runTransactionCommitSuccessSuite(ti TransactionSuiteTestInterface, t *testing.T) {
	t.Run("TransactionCommitSuccess", func(t *testing.T) {
		tix := ti.BeginTx()

		if err := tix.Create(); err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		tix.CorrectCheck(t)

		if err := tix.Commit(); err != nil {
			t.Errorf("Expected no error on rollback, got %v", err)
		}

		if err := tix.Find(); !errors.Is(repository.ErrTxAlreadyRolledBack, err) {
			t.Errorf("Expected %v, recieved %v", repository.ErrTxAlreadyRolledBack, err)
		}

		if err := ti.Find(); err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
	})
}

// Helper functions for errors.
func expectedGot(exp interface{}, got interface{}) string {
	return fmt.Sprintf("Expected %v, got %v", exp, got)
}

func unexpectedErr(got interface{}) string {
	return fmt.Sprintf("Expected no error, got %v", got)
}
