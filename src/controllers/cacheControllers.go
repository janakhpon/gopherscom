package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
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
