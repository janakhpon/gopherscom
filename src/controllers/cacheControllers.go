package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/janakhpon/gopherscom/src/models"
)

func GetCachedUser(c *gin.Context) {
	user := &models.User{}
	val, err := rdbClient.Get("userinfo").Result()
	if err != nil {
		c.JSON(http.StatusAccepted, gin.H{
			"msg": "failed to get user from cache",
		})
		return
	}
	err = json.Unmarshal([]byte(val), &user)
	fmt.Printf("%+v\n", user)

	fmt.Println(user.EMAIL)

	fmt.Println(val)

	c.JSON(http.StatusAccepted, gin.H{
		"msg":  "succed",
		"user": user,
	})

}

func GetCachedProfile(c *gin.Context) {
	profile := &models.Profile{}
	val, err := rdbClient.Get("profileinfo").Result()
	if err != nil {
		c.JSON(http.StatusAccepted, gin.H{
			"msg": "failed to get user from cache",
		})
		return
	}
	err = json.Unmarshal([]byte(val), &profile)

	c.JSON(http.StatusAccepted, gin.H{
		"msg":  "succed",
		"user": profile,
	})

}

func SetKeys(c *gin.Context) {
	var apptypeBody models.Apptype
	c.BindJSON(&apptypeBody)

	apptype := models.Apptype{
		ID:          uuid.New().String(),
		NAME:        "testing",
		DESCRIPTION: "testing",
		AUTHOR:      c.Request.URL.Query().Get("authorid"),
		CREATEDAT:   time.Now(),
		UPDATEDAT:   time.Now(),
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

func GetKeys(c *gin.Context) {
	keys := rdbClient.Keys("apptype*")
	keyres := keys.Val()
	// if err != nil {
	// 	// handle error
	// }
	// for _, key := range keyres {
	// 	fmt.Println(key)
	// }
	fmt.Println(reflect.TypeOf(keys))
	fmt.Print(keys)
	fmt.Println(keyres)
	fmt.Println(keyres[0])
	for i, element := range keyres {
		// Convert element to string to display it.
		fmt.Println(i, string(element))
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "created",
		"data":    &keyres,
	})

	return

}
