package main

import (
	"flag"
	"gorm.io/gorm"
	"log"
	"mnimidamonbackend/domain/repository/sqliterepo"
)

var (
	inMemoryFlag  = false
	dropTableFlag = false

	inMemoryDb = "file::memory:?cache=shared"
	fileDb     = "../databasefiles/mnimidamon.db"
)

func main() {

	flag.BoolVar(&inMemoryFlag, "memory-db", false, "in memory database")
	flag.BoolVar(&dropTableFlag, "drop-table", false, "drop table on start")
	flag.Parse()

	dbLocation := fileDb
	if inMemoryFlag {
		dbLocation = inMemoryDb
	}

	_, err := sqliterepo.Initialize(dbLocation, &gorm.Config{}, dropTableFlag)

	if err != nil {
		log.Fatalf("Database initialization gone wrong: %s", err)
	}

	log.Print("Successfully started the database.")
}
