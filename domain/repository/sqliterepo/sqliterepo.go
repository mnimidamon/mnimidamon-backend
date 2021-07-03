package sqliterepo

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	. "mnimidamonbackend/domain/repository/sqliterepo/modelsql"
)

// Models that are used for our SQLite Gorm repository implementations.
var models = []interface{} {
	&User{},
	&Computer{},
	&Invite{},
	&Group{},
	&GroupComputer{},
	&GroupComputerBackup{},
	&Backup{},
}

// Open a database connection to the SQLite database.
func newDatabaseConnection(databaseFilePath string, opts gorm.Option) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(databaseFilePath), opts)

	if err != nil {
		return nil, err
	}

	return db, nil
}

// Initialize the database and return the pointer.
func Initialize(databaseFilePath string, opts gorm.Option, dropTable bool)  (*gorm.DB, error) {
	db, err := newDatabaseConnection(databaseFilePath, opts)
	if err != nil {
		return nil, err
	}

	// Drop the tables if this is specified.
	if dropTable {

		err := db.
			Migrator().
			DropTable(models...)

		if err != nil {
			return nil, err
		}
	}

	// Migrate the tables if these do not exist.
	if err := db.AutoMigrate(models...); err != nil {
		return nil, fmt.Errorf("migrating gorm models produced an error: %w", err)
	}

	return db, nil
}