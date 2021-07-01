package sqliterepo

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Open a database connection to the SQLite database.
func newDatabaseConnection(databaseFilePath string, opts gorm.Option) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(databaseFilePath), opts)

	if err != nil {
		return nil, err
	}

	return db, nil
}

// Initialize the database and return the pointer.
func Initialize(databaseFilePath string, opts gorm.Option)  (*gorm.DB, error) {
	db, err := newDatabaseConnection(databaseFilePath, opts)

	if err != nil {
		return nil, err
	}

	// TODO: Initialize database models.

	return db, nil
}