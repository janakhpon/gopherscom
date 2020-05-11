package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/janakhpon/gopherscom/src/models"
)

func GetLanguageList(c *gin.Context) {
	var languageList []models.Language
	err := dbConnect.Model(&languageList).Select()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": languageList,
	})
	return
}

func GetLanguage(c *gin.Context) {
	id := c.Request.URL.Query().Get("id")
	language := &models.Language{ID: id}
	err := dbConnect.Select(language)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "failed to fetch",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "succeed",
		"data": language,
	})
	return
}

func CreateLanguage(c *gin.Context) {
	var languageBody models.Language
	c.BindJSON(&languageBody)

	language := models.Language{
		ID:          uuid.New().String(),
		NAME:        languageBody.NAME,
		DESCRIPTION: languageBody.DESCRIPTION,
		AUTHOR:      c.Request.URL.Query().Get("authorid"),
		CREATEDAT:   time.Now(),
		UPDATEDAT:   time.Now(),
	}

	insertError := dbConnect.Insert(&language)
	if insertError != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": insertError,
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "created",
		"data":    &language,
	})

	return
}

func UpdateLanguage(c *gin.Context) {
	id := c.Request.URL.Query().Get("id")
	var languageBody models.Language
	c.BindJSON(&languageBody)
	relanguage := &models.Language{ID: id}

	err := dbConnect.Select(relanguage)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "failed to fetch",
		})
		return
	}

	language := models.Language{
		ID:          id,
		NAME:        languageBody.NAME,
		DESCRIPTION: languageBody.DESCRIPTION,
		AUTHOR:      relanguage.AUTHOR,
		CREATEDAT:   relanguage.CREATEDAT,
		UPDATEDAT:   time.Now(),
	}
	updateError := dbConnect.Update(&language)

	if updateError != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": updateError,
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "updated",
		"data":    &language,
	})
	return
}

func DeleteLanguage(c *gin.Context) {
	id := c.Request.URL.Query().Get("id")
	var languageBody models.Language
	c.BindJSON(&languageBody)
	language := &models.Language{ID: id}

	err := dbConnect.Delete(language)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Deleted!",
	})
	return
}
