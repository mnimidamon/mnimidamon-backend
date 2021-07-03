package sqliterepo_test

import (
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"mnimidamonbackend/domain/repository/sqliterepo"
	"mnimidamonbackend/testsuite"
	"testing"
)

var inMemoryDb = "file::memory:?cache=shared"
func TestSQLiteUserRepository(t *testing.T) {
	db, err := sqliterepo.Initialize("../../../../databasefiles/mnimidamon.db", &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	}, true)

	if err != nil {
		t.Errorf("Error occured with new datbase connection: %v", err)
	}

	ur := sqliterepo.NewUserRepository(db)

	// Call the interface testing suite.
	testsuite.UserRepositoryTestSuite(t, ur)
}