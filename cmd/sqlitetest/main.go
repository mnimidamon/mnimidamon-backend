package main

import (
	"gorm.io/gorm"
	"log"
	"mnimidamonbackend/domain/repository/sqliterepo"
)

var inMemoryDb = "file::memory:?cache=shared"
var fileDb = "../databasefiles/mnimidamon.db"

func main() {
	_, err := sqliterepo.Initialize(fileDb, &gorm.Config{}, true)

	if err != nil {
		log.Fatalf("Database initialization gone wrong: %s", err)
	}

	log.Print("Successfully started the database.")
}