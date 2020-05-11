package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/janakhpon/gopherscom/src/models"
)

func GetApptypeList(c *gin.Context) {
	var apptypeList []models.Apptype
	err := dbConnect.Model(&apptypeList).Select()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": apptypeList,
	})
	return
}

func GetApptype(c *gin.Context) {
	id := c.Request.URL.Query().Get("id")
	apptype := &models.Apptype{ID: id}
	err := dbConnect.Select(apptype)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "failed to fetch",
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

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Deleted!",
	})
	return
}
