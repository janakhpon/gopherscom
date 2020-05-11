package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/janakhpon/gopherscom/src/models"
)

func GetTagList(c *gin.Context) {
	var tagList []models.Tag
	err := dbConnect.Model(&tagList).Select()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": tagList,
	})
	return
}

func GetTag(c *gin.Context) {
	id := c.Request.URL.Query().Get("id")
	tag := &models.Tag{ID: id}
	err := dbConnect.Select(tag)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "failed to fetch",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "succeed",
		"data": tag,
	})
	return
}

func CreateTag(c *gin.Context) {
	var tagBody models.Tag
	c.BindJSON(&tagBody)

	tag := models.Tag{
		ID:          uuid.New().String(),
		NAME:        tagBody.NAME,
		DESCRIPTION: tagBody.DESCRIPTION,
		AUTHOR:      c.Request.URL.Query().Get("authorid"),
		CREATEDAT:   time.Now(),
		UPDATEDAT:   time.Now(),
	}

	insertError := dbConnect.Insert(&tag)
	if insertError != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": insertError,
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "created",
		"data":    &tag,
	})

	return
}

func UpdateTag(c *gin.Context) {
	id := c.Request.URL.Query().Get("id")
	var tagBody models.Tag
	c.BindJSON(&tagBody)
	retag := &models.Tag{ID: id}

	err := dbConnect.Select(retag)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "failed to fetch",
		})
		return
	}

	tag := models.Tag{
		ID:          id,
		NAME:        tagBody.NAME,
		DESCRIPTION: tagBody.DESCRIPTION,
		AUTHOR:      retag.AUTHOR,
		CREATEDAT:   retag.CREATEDAT,
		UPDATEDAT:   time.Now(),
	}
	updateError := dbConnect.Update(&tag)

	if updateError != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": updateError,
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "updated",
		"data":    &tag,
	})
	return
}

func DeleteTag(c *gin.Context) {
	id := c.Request.URL.Query().Get("id")
	var tagBody models.Tag
	c.BindJSON(&tagBody)
	tag := &models.Tag{ID: id}

	err := dbConnect.Delete(tag)
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
