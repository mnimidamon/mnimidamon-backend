package sqliterepo

import (
	"errors"
	"github.com/mattn/go-sqlite3"
	"gorm.io/gorm"
	"log"
	"mnimidamonbackend/domain/repository"
)

func toBusinessLogicError(err error) error {
	// No record found.
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return repository.ErrNotFound
	}

	// Unique constraint violation.
	if errors.As(sqlite3.ErrConstraint, &err) {
		return repository.ErrUniqueConstraintViolation
	}

	log.Printf("Unknown SQLite error: %t --> %v", err, err)
	return repository.UnknownRepositoryError
}