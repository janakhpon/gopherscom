package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/janakhpon/gopherscom/src/models"
)

func GetLibraryList(c *gin.Context) {
	var libraryList []models.Library
	var library models.Library

	keys := rdbClient.Keys("library*")
	keyres := keys.Val()

	for _, key := range keyres {
		val, err := rdbClient.Get(key).Result()
		if err != nil {
			c.JSON(http.StatusAccepted, gin.H{
				"msg": "failed to get user from cache",
			})
			return
		}
		err = json.Unmarshal([]byte(val), &library)
		if library.AUTHOR != "" {
			libraryList = append(libraryList, library)
		}
	}

	if len(libraryList) != 0 {
		c.JSON(http.StatusOK, gin.H{
			"data":   libraryList,
			"status": "from redis",
		})
		return
	}

	err := dbConnect.Model(&libraryList).Select()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "something went wrong",
		})
		return
	}
	for _, key := range libraryList {
		library := models.Library{
			ID:          key.ID,
			NAME:        key.NAME,
			DESCRIPTION: key.DESCRIPTION,
			AUTHOR:      key.AUTHOR,
			CREATEDAT:   key.CREATEDAT,
			UPDATEDAT:   key.UPDATEDAT,
		}
		cachelibrary, err := json.Marshal(library)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err,
			})
			return
		}
		err = rdbClient.Set("library"+library.ID, cachelibrary, 604800*time.Second).Err()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err,
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"data": libraryList,
	})
	return
}

func GetLibrary(c *gin.Context) {
	id := c.Request.URL.Query().Get("id")
	library := &models.Library{ID: id}
	val, err := rdbClient.Get("library" + id).Result()
	if err != nil {
		c.JSON(http.StatusAccepted, gin.H{
			"msg": "failed to get user from cache",
		})
	} else {
		err = json.Unmarshal([]byte(val), &library)
		if library != nil {
			c.JSON(http.StatusOK, gin.H{
				"msg":    "succeed",
				"data":   library,
				"status": "from redis",
			})
			return
		}
	}

	err = dbConnect.Select(library)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "failed to fetch",
		})
		return
	}

	cachelibrary, err := json.Marshal(library)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	err = rdbClient.Set("library"+library.ID, cachelibrary, 604800*time.Second).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
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
	cachelibrary, err := json.Marshal(library)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	err = rdbClient.Set("library"+library.ID, cachelibrary, 604800*time.Second).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
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

	cachelibrary, err := json.Marshal(library)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	err = rdbClient.Set("library"+library.ID, cachelibrary, 604800*time.Second).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
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
	err = rdbClient.Del("library" + id).Err()
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Deleted!",
	})
	return
}
