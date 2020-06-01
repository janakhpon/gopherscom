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
	var companiesList []models.Company
	var company models.Company

	keys := rdbClient.Keys("company*")
	keyres := keys.Val()

	for _, key := range keyres {
		val, err := rdbClient.Get(key).Result()
		if err != nil {
			c.JSON(http.StatusAccepted, gin.H{
				"msg": "failed to get user from cache",
			})
			return
		}
		err = json.Unmarshal([]byte(val), &company)
		if company.NAME != "" {
			companiesList = append(companiesList, company)
		}
	}

	if len(companiesList) != 0 {
		c.JSON(http.StatusOK, gin.H{
			"data":   companiesList,
			"status": "from redis",
		})
		return
	}

	err := dbConnect.Model(&companiesList).Select()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "something went wrong",
		})
		return
	}
	for _, key := range companiesList {
		company := models.Company{
			ID:          key.ID,
			NAME:        key.NAME,
			PRODUCTS:    key.PRODUCTS,
			EMPLOYEE:    key.EMPLOYEE,
			FRAMEWORKS:  key.FRAMEWORKS,
			LANGUAGES:   key.LANGUAGES,
			PLATFORMS:   key.PLATFORMS,
			DATABASES:   key.DATABASES,
			OTHERS:      key.OTHERS,
			ADDRESS:     key.ADDRESS,
			ZIPCODE:     key.ZIPCODE,
			CITY:        key.CITY,
			STATE:       key.STATE,
			COUNTRY:     key.COUNTRY,
			LAT:         key.LAT,
			LON:         key.LON,
			FOUNDEDYEAR: key.FOUNDEDYEAR,
			CREATEDAT:   key.CREATEDAT,
			UPDATEDAT:   key.UPDATEDAT,
		}
		cachecompany, err := json.Marshal(company)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err,
			})
			return
		}
		err = rdbClient.Set("company"+company.ID, cachecompany, 604800*time.Second).Err()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err,
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"data": companiesList,
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

	err = rdbClient.Set("company"+company.ID, cachecompany, 604800*time.Second).Err()
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
