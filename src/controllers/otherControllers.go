package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/janakhpon/gopherscom/src/models"
)

func GetOtherList(c *gin.Context) {
	var otherList []models.Other
	var other models.Other

	keys := rdbClient.Keys("other*")
	keyres := keys.Val()

	for _, key := range keyres {
		val, err := rdbClient.Get(key).Result()
		if err != nil {
			c.JSON(http.StatusAccepted, gin.H{
				"msg": "failed to get user from cache",
			})
			return
		}
		err = json.Unmarshal([]byte(val), &other)
		if other.AUTHOR != "" {
			otherList = append(otherList, other)
		}
	}

	if len(otherList) != 0 {
		c.JSON(http.StatusOK, gin.H{
			"data":   otherList,
			"status": "from redis",
		})
		return
	}

	err := dbConnect.Model(&otherList).Select()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "something went wrong",
		})
		return
	}

	for _, key := range otherList {
		other := models.Other{
			ID:          key.ID,
			NAME:        key.NAME,
			DESCRIPTION: key.DESCRIPTION,
			AUTHOR:      key.AUTHOR,
			CREATEDAT:   key.CREATEDAT,
			UPDATEDAT:   key.UPDATEDAT,
		}
		cacheother, err := json.Marshal(other)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err,
			})
			return
		}
		err = rdbClient.Set("other"+other.ID, cacheother, 604800*time.Second).Err()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err,
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"data": otherList,
	})
	return
}

func GetOther(c *gin.Context) {
	id := c.Request.URL.Query().Get("id")
	other := &models.Other{ID: id}
	val, err := rdbClient.Get("other" + id).Result()
	if err != nil {
		c.JSON(http.StatusAccepted, gin.H{
			"msg": "failed to get user from cache",
		})
	} else {
		err = json.Unmarshal([]byte(val), &other)
		if other != nil {
			c.JSON(http.StatusOK, gin.H{
				"msg":    "succeed",
				"data":   other,
				"status": "from redis",
			})
			return
		}
	}

	err = dbConnect.Select(other)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "failed to fetch",
		})
		return
	}
	cacheother, err := json.Marshal(other)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	err = rdbClient.Set("other"+other.ID, cacheother, 604800*time.Second).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
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
	cacheother, err := json.Marshal(other)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	err = rdbClient.Set("other"+other.ID, cacheother, 604800*time.Second).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
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
	cacheother, err := json.Marshal(other)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	err = rdbClient.Set("other"+other.ID, cacheother, 604800*time.Second).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
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
	err = rdbClient.Del("other" + id).Err()
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Deleted!",
	})
	return
}

func ResetOtherCache(c *gin.Context) {
	var otherList []models.Other
	var other models.Other

	keys := rdbClient.Keys("other*")
	keyres := keys.Val()

	for _, key := range keyres {
		val, err := rdbClient.Get(key).Result()
		if err != nil {
			c.JSON(http.StatusAccepted, gin.H{
				"msg": "failed to get user from cache",
			})
			return
		}
		err = json.Unmarshal([]byte(val), &other)
		if other.AUTHOR != "" {
			otherList = append(otherList, other)
		}
	}

	if len(otherList) != 0 {
		for _, key := range otherList {
			err := rdbClient.Del("other" + key.ID).Err()
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
