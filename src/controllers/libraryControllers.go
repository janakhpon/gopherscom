package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/janakhpon/gopherscom/src/models"
)

func GetLibraryList(c *gin.Context) {
	var libraryList []models.Library
	err := dbConnect.Model(&libraryList).Select()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": libraryList,
	})
	return
}

func GetLibrary(c *gin.Context) {
	id := c.Request.URL.Query().Get("id")
	library := &models.Library{ID: id}
	err := dbConnect.Select(library)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "failed to fetch",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "succeed",
		"data": library,
	})
	return
}

func CreateLibrary(c *gin.Context) {
	var libraryBody models.Library
	c.BindJSON(&libraryBody)

	library := models.Library{
		ID:          uuid.New().String(),
		NAME:        libraryBody.NAME,
		DESCRIPTION: libraryBody.DESCRIPTION,
		LANGUAGES:   libraryBody.LANGUAGES,
		AUTHOR:      c.Request.URL.Query().Get("authorid"),
		CREATEDAT:   time.Now(),
		UPDATEDAT:   time.Now(),
	}

	insertError := dbConnect.Insert(&library)
	if insertError != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": insertError,
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "created",
		"data":    &library,
	})

	return
}

func UpdateLibrary(c *gin.Context) {
	id := c.Request.URL.Query().Get("id")
	var libraryBody models.Library
	c.BindJSON(&libraryBody)
	relibrary := &models.Library{ID: id}

	err := dbConnect.Select(relibrary)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "failed to fetch",
		})
		return
	}

	library := models.Library{
		ID:          id,
		NAME:        libraryBody.NAME,
		DESCRIPTION: libraryBody.DESCRIPTION,
		LANGUAGES:   libraryBody.LANGUAGES,
		AUTHOR:      relibrary.AUTHOR,
		CREATEDAT:   relibrary.CREATEDAT,
		UPDATEDAT:   time.Now(),
	}
	updateError := dbConnect.Update(&library)

	if updateError != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": updateError,
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "updated",
		"data":    &library,
	})
	return
}

func DeleteLibrary(c *gin.Context) {
	id := c.Request.URL.Query().Get("id")
	var libraryBody models.Library
	c.BindJSON(&libraryBody)
	library := &models.Library{ID: id}

	err := dbConnect.Delete(library)
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
