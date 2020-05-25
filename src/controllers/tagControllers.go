package controllers

import (
	"encoding/json"
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
	val, err := rdbClient.Get(id).Result()
	if err != nil {

	}
	err = json.Unmarshal([]byte(val), &tag)
	if tag != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "succeed",
			"data": tag,
		})
		return
	}
	err = dbConnect.Select(tag)
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
	cachetag, err := json.Marshal(tag)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}
	err = rdbClient.Set(tag.ID, cachetag, 604800*time.Second).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
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
	cachetag, err := json.Marshal(tag)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}
	err = rdbClient.Set(tag.ID, cachetag, 604800*time.Second).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
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
	err = rdbClient.Del(id).Err()
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Deleted!",
	})
	return
}
