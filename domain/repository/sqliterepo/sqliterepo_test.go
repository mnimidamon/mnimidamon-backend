package sqliterepo_test

import (
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"mnimidamonbackend/domain/repository/sqliterepo"
	"mnimidamonbackend/testsuites"
	"testing"
)

var inMemoryDb = "file::memory:?cache=shared"
var fileDB = "../../../../databasefiles/mnimidamon.db"

func initializeDatabase() (*gorm.DB, error) {
	return sqliterepo.Initialize(targetDatabasePath(), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	}, true)
}

func targetDatabasePath() string {
	return fileDB
}

func TestSQLiteUserRepository(t *testing.T) {
	db, err := initializeDatabase()

	if err != nil {
		t.Errorf("Error occured with new datbase connection: %v", err)
	}

	ur := sqliterepo.NewUserRepository(db)
	// Call the interface testing suite.
	testsuites.UserRepositoryTestSuite(t, ur)
}

func TestSQLiteGroupRepository(t *testing.T) {
	db, err := initializeDatabase()

	if err != nil {
		t.Errorf("Error occured with new datbase connection: %v", err)
	}

	gr := sqliterepo.NewGroupRepository(db)
	ur := sqliterepo.NewUserRepository(db)
	testsuites.GroupRepositoryTestSuite(t, gr, ur)
}


func TestSQLiteBackupRepository(t *testing.T) {
	db, err := initializeDatabase()

	if err != nil {
		t.Errorf("Error occured with new datbase connection: %v", err)
	}

	gr := sqliterepo.NewGroupRepository(db)
	ur := sqliterepo.NewUserRepository(db)
	br := sqliterepo.NewBackupRepository(db)
	testsuites.BackupRepositoryTestSuite(t, br, gr, ur)
}

func TestSQLiteComputerRepository(t *testing.T) {
	db, err := initializeDatabase()

	if err != nil {
		t.Errorf("Error occured with new datbase connection: %v", err)
	}

	cr := sqliterepo.NewComputerRepository(db)
	ur := sqliterepo.NewUserRepository(db)

	testsuites.ComputerRepositoryTestSuite(t, cr, ur)
}

func TestSQLiteInivteRepository(t *testing.T) {
	db, err := initializeDatabase()

	if err != nil {
		t.Errorf("Error occured with new datbase connection: %v", err)
	}
	ur := sqliterepo.NewUserRepository(db)
	gr := sqliterepo.NewGroupRepository(db)
	ir := sqliterepo.NewInviteRepository(db)

	testsuites.InviteRepositoryTestSuite(t, ir, gr, ur)
}

