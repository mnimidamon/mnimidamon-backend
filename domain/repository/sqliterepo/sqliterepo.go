package sqliterepo

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	. "mnimidamonbackend/domain/repository/sqliterepo/modelsql"
)

// Models that are used for our SQLite Gorm repository implementations.
var models = []interface{} {
	&Group{},
	&GroupComputer{},
	&ComputerBackup{},
	&Computer{},
	&Invite{},
	&User{},
	&Backup{},
}

var relationTables = []string {
	"group_members",
	"group_invites",
	"group_computers",
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

		// Model tables.
		err := db.
			Migrator().
			DropTable(models...)

		if err != nil {
			return nil, err
		}

		// Relational tables.
		err = db.
			Migrator().
			DropTable(
				"group_members",
				"group_invites",
				"group_computers",
		)

		if err != nil {
			return nil, err
		}
	}


	// Migrate the tables if these do not exist.
	if err := db.AutoMigrate(models...); err != nil {
		return nil, fmt.Errorf("migrating gorm models produced an error: %w", err)
	}

	// Ensure foreign key constraints to be checked. SQLite 3.x backwards compatibility.
	db.Exec("PRAGMA foreign_keys = ON")
	return db, nil
}