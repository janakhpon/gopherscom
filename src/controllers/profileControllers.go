package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/janakhpon/gopherscom/src/models"
)

func GetProfileList(c *gin.Context) {
	var profileList []models.Profile
	var profile models.Profile

	keys := rdbClient.Keys("profile*")
	keyres := keys.Val()

	for _, key := range keyres {
		val, err := rdbClient.Get(key).Result()
		if err != nil {
			c.JSON(http.StatusAccepted, gin.H{
				"msg": "failed to get user from cache",
			})
			return
		}
		err = json.Unmarshal([]byte(val), &profile)
		if profile.ID != "" {
			profileList = append(profileList, profile)
		}
	}

	if len(profileList) != 0 {
		c.JSON(http.StatusOK, gin.H{
			"data":   profileList,
			"status": "from redis",
		})
		return
	}
	err := dbConnect.Model(&profileList).Select()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "something went wrong",
		})
		return
	}

	for _, key := range profileList {
		profile := models.Profile{
			ID:         key.ID,
			USERID:     key.USERID,
			CAREER:     key.CAREER,
			FRAMEWORKS: key.FRAMEWORKS,
			LANGUAGES:  key.LANGUAGES,
			PLATFORMS:  key.PLATFORMS,
			DATABASES:  key.DATABASES,
			OTHERS:     key.OTHERS,
			SEX:        key.SEX,
			BIRTHDATE:  key.BIRTHDATE,
			ADDRESS:    key.ADDRESS,
			ZIPCODE:    key.ZIPCODE,
			CITY:       key.CITY,
			STATE:      key.STATE,
			COUNTRY:    key.COUNTRY,
			LAT:        key.LAT,
			LON:        key.LON,
			CREATEDAT:  key.CREATEDAT,
			UPDATEDAT:  key.UPDATEDAT,
		}
		cacheprofile, err := json.Marshal(profile)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err,
			})
			return
		}
		err = rdbClient.Set("profile"+profile.ID, cacheprofile, 604800*time.Second).Err()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err,
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"data": profileList,
	})
	return
}

func CreateProfile(c *gin.Context) {
	var profileBody models.Profile
	c.BindJSON(&profileBody)

	profile := models.Profile{
		ID:         uuid.New().String(),
		USERID:     c.Request.URL.Query().Get("userid"),
		CAREER:     profileBody.CAREER,
		FRAMEWORKS: profileBody.FRAMEWORKS,
		LANGUAGES:  profileBody.LANGUAGES,
		PLATFORMS:  profileBody.PLATFORMS,
		DATABASES:  profileBody.DATABASES,
		OTHERS:     profileBody.OTHERS,
		SEX:        profileBody.SEX,
		BIRTHDATE:  profileBody.BIRTHDATE,
		ADDRESS:    profileBody.ADDRESS,
		ZIPCODE:    profileBody.ZIPCODE,
		CITY:       profileBody.CITY,
		STATE:      profileBody.STATE,
		COUNTRY:    profileBody.COUNTRY,
		LAT:        profileBody.LAT,
		LON:        profileBody.LON,
		CREATEDAT:  time.Now(),
		UPDATEDAT:  time.Now(),
	}

	insertError := dbConnect.Insert(&profile)
	if insertError != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": insertError,
		})
		return
	}
	cacheprofile, err := json.Marshal(profile)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}
	err = rdbClient.Set("profile"+profile.ID, cacheprofile, 604800*time.Second).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "created",
		"data":    &profile,
	})

	return
}

func UpdateProfile(c *gin.Context) {
	userid := c.Request.URL.Query().Get("userid")
	var profileBody models.Profile
	c.BindJSON(&profileBody)
	resprofile := &models.Profile{}

	err := dbConnect.Model(resprofile).Where("userid = ?", userid).Select()

	fmt.Println(err)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "failed to fetch",
		})
		return
	}

	profile := models.Profile{
		ID:         resprofile.ID,
		USERID:     userid,
		CAREER:     profileBody.CAREER,
		FRAMEWORKS: profileBody.FRAMEWORKS,
		LANGUAGES:  profileBody.LANGUAGES,
		PLATFORMS:  profileBody.PLATFORMS,
		DATABASES:  profileBody.DATABASES,
		OTHERS:     profileBody.OTHERS,
		SEX:        profileBody.SEX,
		BIRTHDATE:  profileBody.BIRTHDATE,
		ADDRESS:    profileBody.ADDRESS,
		ZIPCODE:    profileBody.ZIPCODE,
		CITY:       profileBody.CITY,
		STATE:      profileBody.STATE,
		COUNTRY:    profileBody.COUNTRY,
		LAT:        profileBody.LAT,
		LON:        profileBody.LON,
		CREATEDAT:  resprofile.CREATEDAT,
		UPDATEDAT:  time.Now(),
	}
	_, err = dbConnect.Model(&profile).Where("userid = ?", userid).Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "Something went wrong",
		})
		return
	}

	// updateError := dbConnect.Update(&profile)

	// if updateError != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{
	// 		"message": updateError,
	// 	})
	// 	return
	// }

	cacheprofile, err := json.Marshal(profile)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}
	err = rdbClient.Set("profile"+profile.ID, cacheprofile, 604800*time.Second).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "updated",
		"data":    &profile,
	})
	return
}

func GetByID(c *gin.Context) {
	id := c.Request.URL.Query().Get("id")
	profile := models.Profile{ID: id}

	val, err := rdbClient.Get("profile" + id).Result()
	if err != nil {
		c.JSON(http.StatusAccepted, gin.H{
			"msg": "failed to get user from cache",
		})
	} else {
		err = json.Unmarshal([]byte(val), &profile)
		if profile.ID != "" {
			c.JSON(http.StatusOK, gin.H{
				"msg":    "succeed",
				"data":   profile,
				"status": "from redis",
			})
			return
		}
	}

	err = dbConnect.Select(&profile)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "failed to fetch",
		})
		return
	}
	cacheprofile, err := json.Marshal(profile)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}
	err = rdbClient.Set("profile"+profile.ID, cacheprofile, 604800*time.Second).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "succeed",
		"data": profile,
	})
	return
}

func GetProfileByUser(c *gin.Context) {
	userid := c.Request.URL.Query().Get("userid")
	resprofile := &models.Profile{USERID: userid}

	err := dbConnect.Select(resprofile)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "failed to fetch",
		})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "failed to fetch",
		})
		return
	}

	cacheprofile, err := json.Marshal(resprofile)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}
	err = rdbClient.Set("profile"+resprofile.ID, cacheprofile, 604800*time.Second).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "succeed",
		"data": resprofile,
	})
	return
}

func ResetProfileCache(c *gin.Context) {
	var profileList []models.Profile
	var profile models.Profile

	keys := rdbClient.Keys("profile*")
	keyres := keys.Val()

	for _, key := range keyres {
		val, err := rdbClient.Get(key).Result()
		if err != nil {
			c.JSON(http.StatusAccepted, gin.H{
				"msg": "failed to get user from cache",
			})
			return
		}
		err = json.Unmarshal([]byte(val), &profile)
		if profile.ID != "" {
			profileList = append(profileList, profile)
		}
	}

	if len(profileList) != 0 {
		for _, key := range profileList {
			err := rdbClient.Del("profile" + key.ID).Err()
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
