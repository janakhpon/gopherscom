package config

import (
	"log"
	"os"

	"github.com/go-pg/pg/v9"
	"github.com/janakhpon/gopherscom/src/controllers"
	"github.com/joho/godotenv"
)

func Connect() *pg.DB {
	godotenv.Load()

	opts := &pg.Options{
		User:     os.Getenv("DBUSER"),
		Password: os.Getenv("PASSWORD"),
		Addr:     os.Getenv("HOST"),
		Database: os.Getenv("DB"),
	}

	var db *pg.DB = pg.Connect(opts)
	if db == nil {
		log.Printf("Failed to connect")
		os.Exit(100)
	}

	log.Printf("Connected to db")
	// controllers.RemoveBlogTable(db)
	// controllers.RemoveProfileTable(db)
	// controllers.RemoveUserTable(db)
	// controllers.RemoveCompanyTable(db)
	// controllers.RemoveCompanyBranchTable(db)
	controllers.CreateBlogTable(db)
	controllers.CreateProfileTable(db)
	controllers.CreateUserTable(db)
	controllers.CreateCompanyTable(db)
	controllers.CreateCompanyBranchTable(db)
	controllers.InitiateDB(db)

	return db
}
