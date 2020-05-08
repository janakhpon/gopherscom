package controllers

import (
	"log"

	"github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/orm"
	"github.com/janakhpon/gopherscom/src/models"
)

// Create User Table
func CreateBlogTable(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	createError := db.CreateTable(&models.Blog{}, opts)
	if createError != nil {
		log.Printf("Error %v\n", createError)
		return createError
	}
	log.Printf("Created BLOG Table ")
	return nil
}

// Create User Table
func CreateProfileTable(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	createError := db.CreateTable(&models.Profile{}, opts)
	if createError != nil {
		log.Printf("Error %v\n", createError)
		return createError
	}
	log.Printf("Created PROFILE Table ")
	return nil
}

// Create User Table
func RemoveBlogTable(db *pg.DB) error {
	opts := &orm.DropTableOptions{
		IfExists: true,
		Cascade:  true,
	}
	createError := db.DropTable(&models.Blog{}, opts)
	if createError != nil {
		log.Printf("Error %v\n", createError)
		return createError
	}
	log.Printf("removed Blog Table from Database ")
	return nil
}

// Create User Table
func RemoveProfileTable(db *pg.DB) error {
	opts := &orm.DropTableOptions{
		IfExists: true,
		Cascade:  true,
	}
	createError := db.DropTable(&models.Profile{}, opts)
	if createError != nil {
		log.Printf("Error %v\n", createError)
		return createError
	}
	log.Printf("removed Profile Table from Database ")
	return nil
}

//db connection intialized
var dbConnect *pg.DB

func InitiateDB(db *pg.DB) {
	dbConnect = db
}
