package sqliterepo

import (
	"database/sql"
	"errors"
	"github.com/mattn/go-sqlite3"
	"gorm.io/gorm"
	"mnimidamonbackend/domain/constants"
	"mnimidamonbackend/domain/repository"
)

func toRepositoryError(err error) repository.ErrRepo {
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

		constants.Log("Unknown SQLite error %v: %v", sqliteErr.ExtendedCode, sqliteErr)
	}

	constants.Log("Unknown SQLiteUserRepository error: %t --> %v", err, err)
	return repository.UnknownRepositoryError
}