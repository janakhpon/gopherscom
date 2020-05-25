package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/janakhpon/gopherscom/src/models"
)

func GetCompanyList(c *gin.Context) {
	var companieslist []models.Company
	err := dbConnect.Model(&companieslist).Select()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": companieslist,
	})
	return
}

func AddCompany(c *gin.Context) {
	var companyBody models.Company
	c.BindJSON(&companyBody)

	company := models.Company{
		ID:          uuid.New().String(),
		NAME:        companyBody.NAME,
		PRODUCTS:    companyBody.PRODUCTS,
		EMPLOYEE:    companyBody.EMPLOYEE,
		FRAMEWORKS:  companyBody.FRAMEWORKS,
		LANGUAGES:   companyBody.LANGUAGES,
		PLATFORMS:   companyBody.PLATFORMS,
		DATABASES:   companyBody.DATABASES,
		OTHERS:      companyBody.OTHERS,
		ADDRESS:     companyBody.ADDRESS,
		ZIPCODE:     companyBody.ZIPCODE,
		CITY:        companyBody.CITY,
		STATE:       companyBody.STATE,
		COUNTRY:     companyBody.COUNTRY,
		LAT:         companyBody.LAT,
		LON:         companyBody.LON,
		FOUNDEDYEAR: companyBody.FOUNDEDYEAR,
		CREATEDAT:   time.Now(),
		UPDATEDAT:   time.Now(),
	}

	insertError := dbConnect.Insert(&company)
	if insertError != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": insertError,
		})
		return
	}

	cachecompany, err := json.Marshal(company)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	err = rdbClient.Set(company.ID, cachecompany, 604800*time.Second).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "created",
		"data":    &company,
	})
	return
}
