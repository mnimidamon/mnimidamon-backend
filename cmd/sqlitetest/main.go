package main

import (
	"flag"
	"gorm.io/gorm"
	"log"
	"mnimidamonbackend/domain/model"
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

	db, err := sqliterepo.Initialize(dbLocation, &gorm.Config{}, dropTableFlag)

	if err != nil {
		log.Fatalf("Database initialization gone wrong: %s", err)
	}

	log.Print("Successfully started the database.")

	ur := sqliterepo.NewUserRepository(db)

	users, err := ur.FindAll()
	if err != nil {
		log.Fatalf("Finding all users gone wrong: %s", err)
	} else {
		log.Printf("UserRepo.FindAll: %v", users)
	}


	user, err := ur.FindByUsername("marmiha")
	if err != nil {
		log.Printf("Error UserRepo.FindByUsername: %v", err)
	} else {
		log.Printf("UserRepo.FindByUsername: %v", user)
	}


	user = &model.User{
		Entity:       model.Entity{},
		Username:     "marmiha",
		PasswordHash: "xxxxxxx12312",
	}

	err = ur.Save(user)
	if err != nil {
		log.Printf("Error UserRepo.Save: %v", err)
	} else {
		log.Printf("UserRepo.Save: %v", user)
	}


	user, err = ur.FindByUsername("marmiha")
	if err != nil {
		log.Printf("Error UserRepo.FindByUsername: %v", err)
	} else {
		log.Printf("UserRepo.FindByUsername: %v", user)
	}

}
