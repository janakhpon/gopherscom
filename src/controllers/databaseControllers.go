package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/janakhpon/gopherscom/src/models"
)

func GetDatabaseList(c *gin.Context) {
	var databaseList []models.Database
	err := dbConnect.Model(&databaseList).Select()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": databaseList,
	})
	return
}

func GetDatabase(c *gin.Context) {
	id := c.Request.URL.Query().Get("id")
	database := &models.Database{ID: id}
	err := dbConnect.Select(database)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "failed to fetch",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "succeed",
		"data": database,
	})
	return
}

func CreateDatabase(c *gin.Context) {
	var databaseBody models.Database
	c.BindJSON(&databaseBody)

	database := models.Database{
		ID:          uuid.New().String(),
		NAME:        databaseBody.NAME,
		DESCRIPTION: databaseBody.DESCRIPTION,
		AUTHOR:      c.Request.URL.Query().Get("authorid"),
		CREATEDAT:   time.Now(),
		UPDATEDAT:   time.Now(),
	}

	insertError := dbConnect.Insert(&database)
	if insertError != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": insertError,
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "created",
		"data":    &database,
	})

	return
}

func UpdateDatabase(c *gin.Context) {
	id := c.Request.URL.Query().Get("id")
	var databaseBody models.Database
	c.BindJSON(&databaseBody)
	redatabase := &models.Database{ID: id}

	err := dbConnect.Select(redatabase)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "failed to fetch",
		})
		return
	}

	database := models.Database{
		ID:          id,
		NAME:        databaseBody.NAME,
		DESCRIPTION: databaseBody.DESCRIPTION,
		AUTHOR:      redatabase.AUTHOR,
		CREATEDAT:   redatabase.CREATEDAT,
		UPDATEDAT:   time.Now(),
	}
	updateError := dbConnect.Update(&database)

	if updateError != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": updateError,
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "updated",
		"data":    &database,
	})
	return
}

func DeleteDatabase(c *gin.Context) {
	id := c.Request.URL.Query().Get("id")
	var databaseBody models.Database
	c.BindJSON(&databaseBody)
	database := &models.Database{ID: id}

	err := dbConnect.Delete(database)
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
