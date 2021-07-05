package sqliterepo

import (
	"database/sql"
	"errors"
	"github.com/mattn/go-sqlite3"
	"gorm.io/gorm"
	"log"
	"mnimidamonbackend/domain/repository"
)

func toBusinessLogicError(err error) error {
	// No record found.
	if errors.Is(gorm.ErrRecordNotFound, err) {
		return repository.ErrNotFound
	}

	// Transaction is already rolled back, operations not permitted.
	if errors.Is(sql.ErrTxDone, err) {
		return repository.ErrTxAlreadyRolledBack
	}

	// SQLite errors.
	if sqliteErr, ok := err.(sqlite3.Error); ok {
		switch sqliteErr.ExtendedCode {
		case sqlite3.ErrConstraintUnique:
			// Unique constraint violation on model fields.
			return repository.ErrUniqueConstraintViolation
		case sqlite3.ErrConstraintPrimaryKey:
			// Unique constraint violation primary key insertion.
			return repository.ErrUniquePrimaryKeyConstraintViolation
		case sqlite3.ErrConstraintForeignKey:
			return repository.ErrForeignKeyConstraintViolation
		}

		log.Printf("Unknow SQLite error %v: %v", sqliteErr.ExtendedCode, sqliteErr)
	}

	log.Printf("Unknown SQLiteUserRepository error: %t --> %v", err, err)
	return repository.UnknownRepositoryError
}