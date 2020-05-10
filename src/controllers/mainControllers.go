package controllers

import (
	"log"

	"github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/orm"
	"github.com/janakhpon/gopherscom/src/models"
)

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

func CreateUserTable(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	createError := db.CreateTable(&models.User{}, opts)
	if createError != nil {
		log.Printf("Error %v\n", createError)
		return createError
	}
	log.Printf("Created USER Table ")
	return nil
}

func CreateCompanyTable(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	createError := db.CreateTable(&models.Company{}, opts)
	if createError != nil {
		log.Printf("Error %v\n", createError)
		return createError
	}
	log.Printf("Created COMPANY Table ")
	return nil
}

func CreateCompanyBranchTable(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	createError := db.CreateTable(&models.Branch{}, opts)
	if createError != nil {
		log.Printf("Error %v\n", createError)
		return createError
	}
	log.Printf("Created COMPANY BRANCH Table ")
	return nil
}

func CreateTagTable(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	createError := db.CreateTable(&models.Tag{}, opts)
	if createError != nil {
		log.Printf("Error %v\n", createError)
		return createError
	}
	log.Printf("Created TAG Table ")
	return nil
}

func CreateApptypeTable(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	createError := db.CreateTable(&models.Apptype{}, opts)
	if createError != nil {
		log.Printf("Error %v\n", createError)
		return createError
	}
	log.Printf("Created APPTYPE Table ")
	return nil
}

func CreateLibraryTable(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	createError := db.CreateTable(&models.Library{}, opts)
	if createError != nil {
		log.Printf("Error %v\n", createError)
		return createError
	}
	log.Printf("Created LIBRARY Table ")
	return nil
}

func CreateOtherTable(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	createError := db.CreateTable(&models.Other{}, opts)
	if createError != nil {
		log.Printf("Error %v\n", createError)
		return createError
	}
	log.Printf("Created OTHER Table ")
	return nil
}

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

func RemoveUserTable(db *pg.DB) error {
	opts := &orm.DropTableOptions{
		IfExists: true,
		Cascade:  true,
	}
	createError := db.DropTable(&models.User{}, opts)
	if createError != nil {
		log.Printf("Error %v\n", createError)
		return createError
	}
	log.Printf("removed User Table from Database ")
	return nil
}

func RemoveCompanyTable(db *pg.DB) error {
	opts := &orm.DropTableOptions{
		IfExists: true,
		Cascade:  true,
	}
	createError := db.DropTable(&models.Company{}, opts)
	if createError != nil {
		log.Printf("Error %v\n", createError)
		return createError
	}
	log.Printf("removed Company Table from Database ")
	return nil
}

func RemoveCompanyBranchTable(db *pg.DB) error {
	opts := &orm.DropTableOptions{
		IfExists: true,
		Cascade:  true,
	}
	createError := db.DropTable(&models.Branch{}, opts)
	if createError != nil {
		log.Printf("Error %v\n", createError)
		return createError
	}
	log.Printf("removed Company Branch Table from Database ")
	return nil
}

func RemoveApptypeTable(db *pg.DB) error {
	opts := &orm.DropTableOptions{
		IfExists: true,
		Cascade:  true,
	}
	createError := db.DropTable(&models.Apptype{}, opts)
	if createError != nil {
		log.Printf("Error %v\n", createError)
		return createError
	}
	log.Printf("removed Apptype Table from Database ")
	return nil
}

func RemoveLibraryTable(db *pg.DB) error {
	opts := &orm.DropTableOptions{
		IfExists: true,
		Cascade:  true,
	}
	createError := db.DropTable(&models.Library{}, opts)
	if createError != nil {
		log.Printf("Error %v\n", createError)
		return createError
	}
	log.Printf("removed Library Table from Database ")
	return nil
}

func RemoveOtherTable(db *pg.DB) error {
	opts := &orm.DropTableOptions{
		IfExists: true,
		Cascade:  true,
	}
	createError := db.DropTable(&models.Other{}, opts)
	if createError != nil {
		log.Printf("Error %v\n", createError)
		return createError
	}
	log.Printf("removed Other Table from Database ")
	return nil
}

func RemoveTagTable(db *pg.DB) error {
	opts := &orm.DropTableOptions{
		IfExists: true,
		Cascade:  true,
	}
	createError := db.DropTable(&models.Tag{}, opts)
	if createError != nil {
		log.Printf("Error %v\n", createError)
		return createError
	}
	log.Printf("removed Tag Table from Database ")
	return nil
}

var dbConnect *pg.DB

func InitiateDB(db *pg.DB) {
	dbConnect = db
}
