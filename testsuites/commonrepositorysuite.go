package testsuites

import (
	"errors"
	"mnimidamonbackend/domain/repository"
	"testing"
)

// For testing the common Transaction implementation.
type TransactionSuiteTestInterface interface {
	BeginTx() TransactionSuiteTestTxInterface
	Find() error
}
type TransactionSuiteTestTxInterface interface {
	Create() error
	Find() error
	CorrectCheck(t *testing.T)
	repository.Transaction
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