package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/janakhpon/gopherscom/src/models"
)

func GetProfileList(c *gin.Context) {
	var profilelist []models.Profile
	err := dbConnect.Model(&profilelist).Select()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": profilelist,
	})
	return
}

func CreateProfile(c *gin.Context) {
	var profileBody models.Profile
	c.BindJSON(&profileBody)

	profile := models.Profile{
		ID:         uuid.New().String(),
		USER:       c.Request.URL.Query().Get("user"),
		CAREER:     profileBody.CAREER,
		LANGUAGES:  profileBody.LANGUAGES,
		FRAMEWORKS: profileBody.FRAMEWORKS,
		DATABASES:  profileBody.DATABASES,
		SEX:        profileBody.SEX,
		BIRTHDATE:  profileBody.BIRTHDATE,
		ADDRESS:    profileBody.ADDRESS,
		ZIPCODE:    profileBody.ZIPCODE,
		CITY:       profileBody.CITY,
		STATE:      profileBody.STATE,
		COUNTRY:    profileBody.COUNTRY,
		CREATEDAT:  time.Now(),
		UPDATEDAT:  time.Now(),
	}

	insertError := dbConnect.Insert(&profile)
	if insertError != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": insertError,
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "created",
		"data":    &profile,
	})

	return
}

func UpdateProfile(c *gin.Context) {
	userid := c.Request.URL.Query().Get("userid")
	var profileBody models.Profile
	c.BindJSON(&profileBody)
	resprofile := &models.Profile{USER: userid}

	err := dbConnect.Select(resprofile)

	fmt.Println(err)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "failed to fetch",
		})
		return
	}

	profile := models.Profile{
		ID:         resprofile.ID,
		USER:       userid,
		CAREER:     profileBody.CAREER,
		LANGUAGES:  profileBody.LANGUAGES,
		FRAMEWORKS: profileBody.FRAMEWORKS,
		DATABASES:  profileBody.DATABASES,
		SEX:        profileBody.SEX,
		BIRTHDATE:  profileBody.BIRTHDATE,
		ADDRESS:    profileBody.ADDRESS,
		ZIPCODE:    profileBody.ZIPCODE,
		CITY:       profileBody.CITY,
		STATE:      profileBody.STATE,
		COUNTRY:    profileBody.COUNTRY,
		CREATEDAT:  resprofile.CREATEDAT,
		UPDATEDAT:  time.Now(),
	}

	updateError := dbConnect.Update(&profile)

	if updateError != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": updateError,
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "updated",
		"data":    &profile,
	})
	return
}

func GetByID(c *gin.Context) {
	id := c.Request.URL.Query().Get("id")
	profile := models.Profile{USER: id}
	err := dbConnect.Select(&profile)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "failed to fetch",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "succeed",
		"data": profile,
	})
	return
}

func GetProfileByUser(c *gin.Context) {
	userid := c.Request.URL.Query().Get("userid")
	resprofile := &models.Profile{USER: userid}

	err := dbConnect.Select(resprofile)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "failed to fetch",
		})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "failed to fetch",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "succeed",
		"data": resprofile,
	})
	return
}
