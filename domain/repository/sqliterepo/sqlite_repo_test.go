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
var fileName = "mnimidamon.db"

func initializeDatabase(t *testing.T) (*gorm.DB, error) {
	return sqliterepo.Initialize(targetDatabasePath(t), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	}, true)
}

func targetDatabasePath(t *testing.T) string {
	return t.TempDir() +"/"+ fileName
}

func TestSQLiteUserRepository(t *testing.T) {
	db, err := initializeDatabase(t)
	if err != nil {
		t.Errorf("Error occured with new datbase connection: %v", err)
	}
	sqlconn, _ := db.DB()
	defer sqlconn.Close()

	ur := sqliterepo.NewUserRepository(db)
	// Call the interface testing suite.
	testsuites.UserRepositoryTestSuite(t, ur)
}

func TestSQLiteGroupRepository(t *testing.T) {
	db, err := initializeDatabase(t)

	if err != nil {
		t.Errorf("Error occured with new datbase connection: %v", err)
	}
	sqlconn, _ := db.DB()
	defer sqlconn.Close()

	gr := sqliterepo.NewGroupRepository(db)
	ur := sqliterepo.NewUserRepository(db)
	testsuites.GroupRepositoryTestSuite(t, gr, ur)
}

func TestSQLiteBackupRepository(t *testing.T) {
	db, err := initializeDatabase(t)

	if err != nil {
		t.Errorf("Error occured with new datbase connection: %v", err)
	}
	sqlconn, _ := db.DB()
	defer sqlconn.Close()

	gr := sqliterepo.NewGroupRepository(db)
	ur := sqliterepo.NewUserRepository(db)
	br := sqliterepo.NewBackupRepository(db)
	testsuites.BackupRepositoryTestSuite(t, br, gr, ur)
}

func TestSQLiteComputerRepository(t *testing.T) {
	db, err := initializeDatabase(t)

	if err != nil {
		t.Errorf("Error occured with new datbase connection: %v", err)
	}
	sqlconn, _ := db.DB()
	defer sqlconn.Close()

	cr := sqliterepo.NewComputerRepository(db)
	ur := sqliterepo.NewUserRepository(db)

	testsuites.ComputerRepositoryTestSuite(t, cr, ur)
}

func TestSQLiteInviteRepository(t *testing.T) {
	db, err := initializeDatabase(t)

	if err != nil {
		t.Errorf("Error occured with new datbase connection: %v", err)
	}
	sqlconn, _ := db.DB()
	defer sqlconn.Close()

	ur := sqliterepo.NewUserRepository(db)
	gr := sqliterepo.NewGroupRepository(db)
	ir := sqliterepo.NewInviteRepository(db)

	testsuites.InviteRepositoryTestSuite(t, ir, gr, ur)
}

func TestSQLiteGroupComputerRepository(t *testing.T) {
	db, err := initializeDatabase(t)

	if err != nil {
		t.Errorf("Error occured with new datbase connection: %v", err)
	}

	sqlconn, _ := db.DB()
	defer sqlconn.Close()

	ur := sqliterepo.NewUserRepository(db)
	gr := sqliterepo.NewGroupRepository(db)
	gcr := sqliterepo.NewGroupComputerRepository(db)
	cr := sqliterepo.NewComputerRepository(db)

	testsuites.GroupComputerRepositoryTestSuite(t, gcr, gr, ur, cr)
}

func TestSQLiteComputerBackupRepository(t *testing.T) {
	db, err := initializeDatabase(t)

	if err != nil {
		t.Errorf("Error occured with new datbase connection: %v", err)
	}
	sqlconn, _ := db.DB()
	defer sqlconn.Close()

	cbr := sqliterepo.NewComputerBackupRepository(db)
	ur := sqliterepo.NewUserRepository(db)
	gr := sqliterepo.NewGroupRepository(db)
	gcr := sqliterepo.NewGroupComputerRepository(db)
	cr := sqliterepo.NewComputerRepository(db)
	br := sqliterepo.NewBackupRepository(db)

	testsuites.ComputerBackupRepositoryTestSuite(t, cbr, gcr, gr, ur, cr, br)
}
