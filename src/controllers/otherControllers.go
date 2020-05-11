package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/janakhpon/gopherscom/src/models"
)

func GetOtherList(c *gin.Context) {
	var otherList []models.Other
	err := dbConnect.Model(&otherList).Select()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": otherList,
	})
	return
}

func GetOther(c *gin.Context) {
	id := c.Request.URL.Query().Get("id")
	other := &models.Other{ID: id}
	err := dbConnect.Select(other)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "failed to fetch",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "succeed",
		"data": other,
	})
	return
}

func CreateOther(c *gin.Context) {
	var otherBody models.Other
	c.BindJSON(&otherBody)

	other := models.Other{
		ID:          uuid.New().String(),
		NAME:        otherBody.NAME,
		DESCRIPTION: otherBody.DESCRIPTION,
		AUTHOR:      c.Request.URL.Query().Get("authorid"),
		CREATEDAT:   time.Now(),
		UPDATEDAT:   time.Now(),
	}

	insertError := dbConnect.Insert(&other)
	if insertError != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": insertError,
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "created",
		"data":    &other,
	})

	return
}

func UpdateOther(c *gin.Context) {
	id := c.Request.URL.Query().Get("id")
	var otherBody models.Other
	c.BindJSON(&otherBody)
	reother := &models.Other{ID: id}

	err := dbConnect.Select(reother)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "failed to fetch",
		})
		return
	}

	other := models.Other{
		ID:          id,
		NAME:        otherBody.NAME,
		DESCRIPTION: otherBody.DESCRIPTION,
		AUTHOR:      reother.AUTHOR,
		CREATEDAT:   reother.CREATEDAT,
		UPDATEDAT:   time.Now(),
	}
	updateError := dbConnect.Update(&other)

	if updateError != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": updateError,
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "updated",
		"data":    &other,
	})
	return
}

func DeleteOther(c *gin.Context) {
	id := c.Request.URL.Query().Get("id")
	var otherBody models.Other
	c.BindJSON(&otherBody)
	other := &models.Other{ID: id}

	err := dbConnect.Delete(other)
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
