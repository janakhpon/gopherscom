package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/janakhpon/gopherscom/src/models"
)

func GetPlatformList(c *gin.Context) {
	var platformList []models.Platform
	err := dbConnect.Model(&platformList).Select()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": platformList,
	})
	return
}

func GetPlatform(c *gin.Context) {
	id := c.Request.URL.Query().Get("id")
	platform := &models.Platform{ID: id}
	err := dbConnect.Select(platform)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "failed to fetch",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "succeed",
		"data": platform,
	})
	return
}

func CreatePlatform(c *gin.Context) {
	var platformBody models.Platform
	c.BindJSON(&platformBody)

	platform := models.Platform{
		ID:          uuid.New().String(),
		NAME:        platformBody.NAME,
		DESCRIPTION: platformBody.DESCRIPTION,
		AUTHOR:      platformBody.AUTHOR,
		CREATEDAT:   time.Now(),
		UPDATEDAT:   time.Now(),
	}

	insertError := dbConnect.Insert(&platform)
	if insertError != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": insertError,
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "created",
		"data":    &platform,
	})

	return
}

func UpdatePlatform(c *gin.Context) {
	id := c.Request.URL.Query().Get("id")
	var platformBody models.Platform
	c.BindJSON(&platformBody)
	replatform := &models.Platform{ID: id}

	err := dbConnect.Select(replatform)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "failed to fetch",
		})
		return
	}

	platform := models.Platform{
		ID:          id,
		NAME:        platformBody.NAME,
		DESCRIPTION: platformBody.DESCRIPTION,
		AUTHOR:      platformBody.AUTHOR,
		CREATEDAT:   replatform.CREATEDAT,
		UPDATEDAT:   time.Now(),
	}
	updateError := dbConnect.Update(&platform)

	if updateError != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": updateError,
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "updated",
		"data":    &platform,
	})
	return
}

func DeletePlatform(c *gin.Context) {
	id := c.Request.URL.Query().Get("id")
	var platformBody models.Platform
	c.BindJSON(&platformBody)
	platform := &models.Platform{ID: id}

	err := dbConnect.Delete(platform)
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
