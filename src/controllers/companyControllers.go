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

func GetCompany(c *gin.Context) {
	id := c.Request.URL.Query().Get("id")
	company := models.Company{ID: id}
	res, err := rdbClient.Get("company" + id).Result()

	if err != nil {
		c.JSON(http.StatusAccepted, gin.H{
			"msg": "failed to get user from cache",
		})
	} else {
		err = json.Unmarshal([]byte(res), &company)
		if company.NAME != "" {
			c.JSON(http.StatusOK, gin.H{
				"msg":    "succeed",
				"data":   company,
				"status": "from redis",
			})
			return
		}
	}

	err = dbConnect.Select(company)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "failed to fetch",
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

	c.JSON(http.StatusOK, gin.H{
		"msg":  "succeed",
		"data": company,
	})
	return
}

func UpdateCompany(c *gin.Context) {
	id := c.Request.URL.Query().Get("id")
	var companyBody models.Company
	c.BindJSON(&companyBody)
	recompany := &models.Company{ID: id}

	err := dbConnect.Select(recompany)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "failed to fetch",
		})
		return
	}

	company := models.Company{
		ID:          id,
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
		CREATEDAT:   recompany.CREATEDAT,
		UPDATEDAT:   time.Now(),
	}
	updateError := dbConnect.Update(&company)
	if updateError != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": updateError,
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
		"message": "updated",
		"data":    &company,
	})
	return
}

func DeleteCompany(c *gin.Context) {
	id := c.Request.URL.Query().Get("id")
	var companyBody models.Company
	c.BindJSON(&companyBody)
	company := &models.Company{ID: id}

	err := dbConnect.Delete(company)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	err = rdbClient.Del("company" + id).Err()
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Deleted!",
	})
	return
}

func ResetCompanyCache(c *gin.Context) {
	var companyList []models.Company
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
		if company.ID != "" {
			companyList = append(companyList, company)
		}
	}

	if len(companyList) != 0 {
		for _, key := range companyList {
			err := rdbClient.Del("company" + key.ID).Err()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": err,
				})
				return
			}
		}

		c.JSON(http.StatusOK, gin.H{
			"msg":    "resetted cache",
			"status": "from redis",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "failed to reset",
	})
	return
}
