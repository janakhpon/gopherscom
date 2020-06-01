package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/janakhpon/gopherscom/src/models"
)

func GetFrameworkList(c *gin.Context) {
	var frameworkList []models.Framework

	var framework models.Framework

	keys := rdbClient.Keys("framework*")
	keyres := keys.Val()

	for _, key := range keyres {
		val, err := rdbClient.Get(key).Result()
		if err != nil {
			c.JSON(http.StatusAccepted, gin.H{
				"msg": "failed to get user from cache",
			})
			return
		}
		err = json.Unmarshal([]byte(val), &framework)
		if framework.AUTHOR != "" {
			frameworkList = append(frameworkList, framework)
		}
	}

	if len(frameworkList) != 0 {
		c.JSON(http.StatusOK, gin.H{
			"data":   frameworkList,
			"status": "from redis",
		})
		return
	}

	err := dbConnect.Model(&frameworkList).Select()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "something went wrong",
		})
		return
	}

	for _, key := range frameworkList {
		framework := models.Framework{
			ID:          key.ID,
			NAME:        key.NAME,
			DESCRIPTION: key.DESCRIPTION,
			AUTHOR:      key.AUTHOR,
			CREATEDAT:   key.CREATEDAT,
			UPDATEDAT:   key.UPDATEDAT,
		}
		cacheframework, err := json.Marshal(framework)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err,
			})
			return
		}
		err = rdbClient.Set("framework"+framework.ID, cacheframework, 604800*time.Second).Err()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err,
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"data": frameworkList,
	})
	return
}

func GetFramework(c *gin.Context) {
	id := c.Request.URL.Query().Get("id")
	framework := &models.Framework{ID: id}
	val, err := rdbClient.Get("framework" + id).Result()
	if err != nil {
		c.JSON(http.StatusAccepted, gin.H{
			"msg": "failed to get user from cache",
		})
	} else {
		err = json.Unmarshal([]byte(val), &framework)
		if framework != nil {
			c.JSON(http.StatusOK, gin.H{
				"msg":    "succeed",
				"data":   framework,
				"status": "from redis",
			})
			return
		}
	}

	err = dbConnect.Select(framework)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "failed to fetch",
		})
		return
	}

	cacheframework, err := json.Marshal(framework)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	err = rdbClient.Set("framework"+framework.ID, cacheframework, 604800*time.Second).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "succeed",
		"data": framework,
	})
	return
}

func CreateFramework(c *gin.Context) {
	var frameworkBody models.Framework
	c.BindJSON(&frameworkBody)

	framework := models.Framework{
		ID:          uuid.New().String(),
		NAME:        frameworkBody.NAME,
		DESCRIPTION: frameworkBody.DESCRIPTION,
		AUTHOR:      c.Request.URL.Query().Get("authorid"),
		CREATEDAT:   time.Now(),
		UPDATEDAT:   time.Now(),
	}

	insertError := dbConnect.Insert(&framework)
	if insertError != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": insertError,
		})
		return
	}

	cacheframework, err := json.Marshal(framework)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	err = rdbClient.Set("framework"+framework.ID, cacheframework, 604800*time.Second).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "created",
		"data":    &framework,
	})
	return
}

func UpdateFramework(c *gin.Context) {
	id := c.Request.URL.Query().Get("id")
	var frameworkBody models.Framework
	c.BindJSON(&frameworkBody)
	reframework := &models.Framework{ID: id}

	err := dbConnect.Select(reframework)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "failed to fetch",
		})
		return
	}

	framework := models.Framework{
		ID:          id,
		NAME:        frameworkBody.NAME,
		DESCRIPTION: frameworkBody.DESCRIPTION,
		AUTHOR:      reframework.AUTHOR,
		CREATEDAT:   reframework.CREATEDAT,
		UPDATEDAT:   time.Now(),
	}
	updateError := dbConnect.Update(&framework)

	if updateError != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": updateError,
		})
		return
	}
	cacheframework, err := json.Marshal(framework)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	err = rdbClient.Set("framework"+framework.ID, cacheframework, 604800*time.Second).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "updated",
		"data":    &framework,
	})
	return
}

func DeleteFramework(c *gin.Context) {
	id := c.Request.URL.Query().Get("id")
	var frameworkBody models.Framework
	c.BindJSON(&frameworkBody)
	framework := &models.Framework{ID: id}

	err := dbConnect.Delete(framework)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}
	err = rdbClient.Del("framework" + id).Err()
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Deleted!",
	})
	return
}
