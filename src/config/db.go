package config

import (
	"log"
	"os"

	"github.com/go-pg/pg/v9"
	"github.com/go-redis/redis"
	"github.com/janakhpon/gopherscom/src/controllers"
	"github.com/joho/godotenv"
)

func Connect() *pg.DB {
	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}


	opts := &pg.Options{
		User:     os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Addr:     os.Getenv("DB_HOST"),
		Database: os.Getenv("DB_DATABASE"),
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
	// controllers.RemoveTagTable(db)
	// controllers.RemoveOtherTable(db)
	// controllers.RemoveLibraryTable(db)
	// controllers.RemoveApptypeTable(db)
	// controllers.RemoveLanguageTable(db)
	// controllers.RemoveFrameworkTable(db)
	// controllers.RemovePlatformTable(db)
	// controllers.RemoveBootcampTable(db)

	controllers.CreateBlogTable(db)
	controllers.CreateProfileTable(db)
	controllers.CreateUserTable(db)
	controllers.CreateCompanyTable(db)
	controllers.CreateCompanyBranchTable(db)
	controllers.CreateApptypeTable(db)
	controllers.CreateLibraryTable(db)
	controllers.CreateOtherTable(db)
	controllers.CreateTagTable(db)
	controllers.CreateLanguageTable(db)
	controllers.CreateFrameworkTable(db)
	controllers.CreatePlatformTable(db)
	controllers.CreateBootcampTable(db)

	controllers.InitiateDB(db)

	return db
}

func ConnectRedis() *redis.Client {
	godotenv.Load()

	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	controllers.InitiateRedis(client)

	return client
}
