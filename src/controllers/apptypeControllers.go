package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/janakhpon/gopherscom/src/models"
)

func GetApptypeList(c *gin.Context) {
	var apptypeList []models.Apptype
	var apptype models.Apptype

	keys := rdbClient.Keys("apptype*")
	keyres := keys.Val()

	for _, key := range keyres {
		val, err := rdbClient.Get(key).Result()
		if err != nil {
			c.JSON(http.StatusAccepted, gin.H{
				"msg": "failed to get user from cache",
			})
			return
		}
		err = json.Unmarshal([]byte(val), &apptype)
		if apptype.AUTHOR != "" {
			apptypeList = append(apptypeList, apptype)
		}
	}

	if len(apptypeList) != 0 {
		c.JSON(http.StatusOK, gin.H{
			"data":   apptypeList,
			"status": "from redis",
		})
		return
	}
	err := dbConnect.Model(&apptypeList).Select()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "something went wrong",
		})
		return
	}
	for _, key := range apptypeList {
		apptype := models.Apptype{
			ID:          key.ID,
			NAME:        key.NAME,
			DESCRIPTION: key.DESCRIPTION,
			AUTHOR:      key.AUTHOR,
			CREATEDAT:   key.CREATEDAT,
			UPDATEDAT:   key.UPDATEDAT,
		}
		cacheapptype, err := json.Marshal(apptype)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err,
			})
			return
		}
		err = rdbClient.Set("apptype"+apptype.ID, cacheapptype, 604800*time.Second).Err()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err,
			})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"data": apptypeList,
	})
	return
}

func GetApptype(c *gin.Context) {
	id := c.Request.URL.Query().Get("id")
	apptype := &models.Apptype{ID: id}
	val, err := rdbClient.Get("apptype" + id).Result()
	if err != nil {
		c.JSON(http.StatusAccepted, gin.H{
			"msg": "failed to get user from cache",
		})
	} else {
		err = json.Unmarshal([]byte(val), &apptype)
		if apptype != nil {
			c.JSON(http.StatusOK, gin.H{
				"msg":    "succeed",
				"data":   apptype,
				"status": "from redis",
			})
			return
		}
	}

	err = dbConnect.Select(apptype)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "failed to fetch",
		})
		return
	}

	cacheapptype, err := json.Marshal(apptype)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	err = rdbClient.Set("apptype"+apptype.ID, cacheapptype, 604800*time.Second).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "succeed",
		"data": apptype,
	})
	return
}

func CreateApptype(c *gin.Context) {
	var apptypeBody models.Apptype
	c.BindJSON(&apptypeBody)

	apptype := models.Apptype{
		ID:          uuid.New().String(),
		NAME:        apptypeBody.NAME,
		DESCRIPTION: apptypeBody.DESCRIPTION,
		AUTHOR:      c.Request.URL.Query().Get("authorid"),
		CREATEDAT:   time.Now(),
		UPDATEDAT:   time.Now(),
	}

	insertError := dbConnect.Insert(&apptype)
	if insertError != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": insertError,
		})
		return
	}
	cacheapptype, err := json.Marshal(apptype)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	err = rdbClient.Set("apptype"+apptype.ID, cacheapptype, 604800*time.Second).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "created",
		"data":    &apptype,
	})

	return
}

func UpdateApptype(c *gin.Context) {
	id := c.Request.URL.Query().Get("id")
	var apptypeBody models.Apptype
	c.BindJSON(&apptypeBody)
	reapptype := &models.Apptype{ID: id}

	err := dbConnect.Select(reapptype)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "failed to fetch",
		})
		return
	}

	apptype := models.Apptype{
		ID:          id,
		NAME:        apptypeBody.NAME,
		DESCRIPTION: apptypeBody.DESCRIPTION,
		AUTHOR:      reapptype.AUTHOR,
		CREATEDAT:   reapptype.CREATEDAT,
		UPDATEDAT:   time.Now(),
	}
	updateError := dbConnect.Update(&apptype)

	if updateError != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": updateError,
		})
		return
	}
	cacheapptype, err := json.Marshal(apptype)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	err = rdbClient.Set("apptype"+apptype.ID, cacheapptype, 604800*time.Second).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "updated",
		"data":    &apptype,
	})
	return
}

func DeleteApptype(c *gin.Context) {
	id := c.Request.URL.Query().Get("id")
	var apptypeBody models.Apptype
	c.BindJSON(&apptypeBody)
	apptype := &models.Apptype{ID: id}

	err := dbConnect.Delete(apptype)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	err = rdbClient.Del("apptype" + id).Err()
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Deleted!",
	})
	return
}

func ResetApptypeCache(c *gin.Context) {
	var apptypeList []models.Apptype
	var apptype models.Apptype

	keys := rdbClient.Keys("apptype*")
	keyres := keys.Val()

	for _, key := range keyres {
		val, err := rdbClient.Get(key).Result()
		if err != nil {
			c.JSON(http.StatusAccepted, gin.H{
				"msg": "failed to get user from cache",
			})
			return
		}
		err = json.Unmarshal([]byte(val), &apptype)
		if apptype.AUTHOR != "" {
			apptypeList = append(apptypeList, apptype)
		}
	}

	if len(apptypeList) != 0 {
		for _, key := range apptypeList {
			err := rdbClient.Del("apptype" + key.ID).Err()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": err,
				})
				return
			}
		}

		c.JSON(http.StatusOK, gin.H{
			"msg":    "reset cache successfully",
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
