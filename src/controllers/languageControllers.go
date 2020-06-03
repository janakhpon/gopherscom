package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/janakhpon/gopherscom/src/models"
)

func GetLanguageList(c *gin.Context) {
	var languageList []models.Language
	var language models.Language

	keys := rdbClient.Keys("language*")
	keyres := keys.Val()

	for _, key := range keyres {
		val, err := rdbClient.Get(key).Result()
		if err != nil {
			c.JSON(http.StatusAccepted, gin.H{
				"msg": "failed to get user from cache",
			})
			return
		}
		err = json.Unmarshal([]byte(val), &language)
		if language.AUTHOR != "" {
			languageList = append(languageList, language)
		}
	}

	if len(languageList) != 0 {
		c.JSON(http.StatusOK, gin.H{
			"data":   languageList,
			"status": "from redis",
		})
		return
	}

	err := dbConnect.Model(&languageList).Select()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "something went wrong",
		})
		return
	}

	for _, key := range languageList {
		language := models.Language{
			ID:          key.ID,
			NAME:        key.NAME,
			DESCRIPTION: key.DESCRIPTION,
			AUTHOR:      key.AUTHOR,
			CREATEDAT:   key.CREATEDAT,
			UPDATEDAT:   key.UPDATEDAT,
		}
		cachelanguage, err := json.Marshal(language)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err,
			})
			return
		}
		err = rdbClient.Set("language"+language.ID, cachelanguage, 604800*time.Second).Err()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err,
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"data": languageList,
	})
	return
}

func GetLanguage(c *gin.Context) {
	id := c.Request.URL.Query().Get("id")
	language := &models.Language{ID: id}
	val, err := rdbClient.Get("language" + id).Result()
	if err != nil {
		c.JSON(http.StatusAccepted, gin.H{
			"msg": "failed to get user from cache",
		})
	} else {
		err = json.Unmarshal([]byte(val), &language)
		if language != nil {
			c.JSON(http.StatusOK, gin.H{
				"msg":    "succeed",
				"data":   language,
				"status": "from redis",
			})
			return
		}
	}

	err = dbConnect.Select(language)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "failed to fetch",
		})
		return
	}

	cachelanguage, err := json.Marshal(language)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	err = rdbClient.Set("language"+language.ID, cachelanguage, 604800*time.Second).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
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
	cachelanguage, err := json.Marshal(language)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	err = rdbClient.Set("language"+language.ID, cachelanguage, 604800*time.Second).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
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

	cachelanguage, err := json.Marshal(language)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	err = rdbClient.Set("language"+language.ID, cachelanguage, 604800*time.Second).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
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
	err = rdbClient.Del("language" + id).Err()
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Deleted!",
	})
	return
}

func ResetLanguageCache(c *gin.Context) {
	var languageList []models.Language
	var language models.Language

	keys := rdbClient.Keys("language*")
	keyres := keys.Val()

	for _, key := range keyres {
		val, err := rdbClient.Get(key).Result()
		if err != nil {
			c.JSON(http.StatusAccepted, gin.H{
				"msg": "failed to get user from cache",
			})
			return
		}
		err = json.Unmarshal([]byte(val), &language)
		if language.AUTHOR != "" {
			languageList = append(languageList, language)
		}
	}

	if len(languageList) != 0 {
		for _, key := range languageList {
			err := rdbClient.Del("language" + key.ID).Err()
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
