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

func RemoveLikeByIndex(s []models.Like, index int) []models.Like {
	return append(s[:index], s[index+1:]...)
}

func RemoveEnrollmentByIndex(s []models.Enroller, index int) []models.Enroller {
	return append(s[:index], s[index+1:]...)
}

func GetBootcampList(c *gin.Context) {
	var bootcampList []models.Bootcamp
	var bootcamp models.Bootcamp

	keys := rdbClient.Keys("bootcamp*")
	keyres := keys.Val()

	for _, key := range keyres {
		val, err := rdbClient.Get(key).Result()
		if err != nil {
			c.JSON(http.StatusAccepted, gin.H{
				"msg": "failed to get user from cache",
			})
			return
		}
		err = json.Unmarshal([]byte(val), &bootcamp)
		if bootcamp.AUTHOR != "" {
			bootcampList = append(bootcampList, bootcamp)
		}
	}

	if len(bootcampList) != 0 {
		c.JSON(http.StatusOK, gin.H{
			"data":   bootcampList,
			"status": "from redis",
		})
		return
	}

	err := dbConnect.Model(&bootcampList).Select()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "something went wrong",
		})
		return
	}

	for _, key := range bootcampList {
		bootcamp := models.Bootcamp{
			ID:          key.ID,
			TOPIC:       key.TOPIC,
			INSTRUCTORS: key.INSTRUCTORS,
			ADDRESS:     key.ADDRESS,
			LAT:         key.LAT,
			LON:         key.LON,
			STUDENTS:    key.STUDENTS,
			ENROLLMENTS: key.ENROLLMENTS,
			DESCRIPTION: key.DESCRIPTION,
			AVAILABLE:   true,
			STARTEDAT:   key.STARTEDAT,
			FINISHEDAT:  key.FINISHEDAT,
			AUTHOR:      key.AUTHOR,
			CREATEDAT:   key.CREATEDAT,
			UPDATEDAT:   time.Now(),
		}
		cachebootcamp, err := json.Marshal(bootcamp)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err,
			})
			return
		}
		err = rdbClient.Set("bootcamp"+bootcamp.ID, cachebootcamp, 604800*time.Second).Err()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err,
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"data": bootcampList,
	})
	return
}

func GetBootcamp(c *gin.Context) {
	id := c.Request.URL.Query().Get("id")
	bootcamp := &models.Bootcamp{ID: id}
	val, err := rdbClient.Get("bootcamp" + id).Result()
	if err != nil {
		c.JSON(http.StatusAccepted, gin.H{
			"msg": "failed to get user from cache",
		})
	} else {
		err = json.Unmarshal([]byte(val), &bootcamp)
		if bootcamp != nil {
			c.JSON(http.StatusOK, gin.H{
				"msg":    "succeed",
				"data":   bootcamp,
				"status": "from redis",
			})
			return
		}
	}
	err = dbConnect.Select(bootcamp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "failed to fetch",
		})
		return
	}

	cachebootcamp, err := json.Marshal(bootcamp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	err = rdbClient.Set("bootcamp"+bootcamp.ID, cachebootcamp, 604800*time.Second).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":  "succeed",
		"data": bootcamp,
	})
	return
}

func CreateBootcamp(c *gin.Context) {
	var bootcampBody models.Bootcamp
	c.BindJSON(&bootcampBody)

	bootcamp := models.Bootcamp{
		ID:          uuid.New().String(),
		TOPIC:       bootcampBody.TOPIC,
		INSTRUCTORS: bootcampBody.INSTRUCTORS,
		ADDRESS:     bootcampBody.ADDRESS,
		LAT:         bootcampBody.LAT,
		LON:         bootcampBody.LON,
		STUDENTS:    bootcampBody.STUDENTS,
		ENROLLMENTS: bootcampBody.ENROLLMENTS,
		DESCRIPTION: bootcampBody.DESCRIPTION,
		AVAILABLE:   true,
		STARTEDAT:   bootcampBody.STARTEDAT,
		FINISHEDAT:  bootcampBody.FINISHEDAT,
		AUTHOR:      c.Request.URL.Query().Get("authorid"),
		CREATEDAT:   time.Now(),
		UPDATEDAT:   time.Now(),
	}

	insertError := dbConnect.Insert(&bootcamp)
	if insertError != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": insertError,
		})
		return
	}

	cachebootcamp, err := json.Marshal(bootcamp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	err = rdbClient.Set("bootcamp"+bootcamp.ID, cachebootcamp, 604800*time.Second).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "created",
		"data":    &bootcamp,
	})

	return
}

func UpdateBootcamp(c *gin.Context) {
	id := c.Request.URL.Query().Get("id")
	var bootcampBody models.Bootcamp
	c.BindJSON(&bootcampBody)
	rebootcamp := &models.Bootcamp{ID: id}

	err := dbConnect.Select(rebootcamp)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "failed to fetch",
		})
		return
	}

	bootcamp := models.Bootcamp{
		ID:          id,
		TOPIC:       bootcampBody.TOPIC,
		INSTRUCTORS: bootcampBody.INSTRUCTORS,
		ADDRESS:     bootcampBody.ADDRESS,
		LAT:         bootcampBody.LAT,
		LON:         bootcampBody.LON,
		STUDENTS:    bootcampBody.STUDENTS,
		ENROLLMENTS: bootcampBody.ENROLLMENTS,
		DESCRIPTION: bootcampBody.DESCRIPTION,
		AVAILABLE:   true,
		STARTEDAT:   bootcampBody.STARTEDAT,
		FINISHEDAT:  bootcampBody.FINISHEDAT,
		AUTHOR:      rebootcamp.AUTHOR,
		CREATEDAT:   rebootcamp.CREATEDAT,
		UPDATEDAT:   time.Now(),
	}
	updateError := dbConnect.Update(&bootcamp)

	if updateError != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": updateError,
		})
		return
	}
	cachebootcamp, err := json.Marshal(bootcamp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	err = rdbClient.Set("bootcamp"+bootcamp.ID, cachebootcamp, 604800*time.Second).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "updated",
		"data":    &bootcamp,
	})
	return
}

func SetBootcampAvailability(c *gin.Context) {
	id := c.Request.URL.Query().Get("id")
	availableval := c.Request.URL.Query().Get("available")
	var bootcampBody models.Bootcamp
	c.BindJSON(&bootcampBody)

	bootcamp := models.Bootcamp{}

	_, err := dbConnect.Model(&bootcamp).Set("available = ?", availableval).Where("id = ?", id).Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "Something went wrong",
		})
		return
	}

	cachebootcamp, err := json.Marshal(bootcamp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	err = rdbClient.Set("bootcamp"+bootcamp.ID, cachebootcamp, 604800*time.Second).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Updated Availability",
	})
	return
}

func EnrollBootcamp(c *gin.Context) {
	id := c.Request.URL.Query().Get("id")
	userid := c.Request.URL.Query().Get("userid")
	fmt.Println(userid)
	var bootcampBody models.Bootcamp
	c.BindJSON(&bootcampBody)

	resbootcamp := &models.Bootcamp{ID: id}
	err := dbConnect.Select(resbootcamp)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "Bootcamp not found!",
		})
		return
	}

	user := &models.User{ID: userid}
	err = dbConnect.Select(user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "User not found!",
		})
		return
	}

	enrollers := resbootcamp.ENROLLMENTS

	if enrollers != nil {
		for _, x := range enrollers {
			if x.ID == userid {
			} else {
				enroller := models.Enroller{
					ID:   userid,
					NAME: user.NAME,
				}
				enrollers = append(enrollers, enroller)
			}
		}
	} else {
		enroller := models.Enroller{
			ID:   userid,
			NAME: user.NAME,
		}
		enrollers = append(enrollers, enroller)
	}

	bootcamp := models.Bootcamp{}

	_, err = dbConnect.Model(&bootcamp).Set("enrollments = ?", enrollers).Where("id = ?", id).Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "Something went wrong",
		})
		return
	}

	cachebootcamp, err := json.Marshal(bootcamp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	err = rdbClient.Set("bootcamp"+bootcamp.ID, cachebootcamp, 604800*time.Second).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Updated Availability",
	})
	return
}

func LikeBootcamp(c *gin.Context) {
	id := c.Request.URL.Query().Get("id")
	userid := c.Request.URL.Query().Get("userid")
	fmt.Println(userid)
	var bootcampBody models.Bootcamp
	c.BindJSON(&bootcampBody)

	resbootcamp := &models.Bootcamp{ID: id}
	err := dbConnect.Select(resbootcamp)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "Bootcamp not found!",
		})
		return
	}

	user := &models.User{ID: userid}
	err = dbConnect.Select(user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "User not found!",
		})
		return
	}

	likers := resbootcamp.LIKES
	if likers != nil {
		for i, x := range likers {
			if x.ID == userid && x.LIKE == true {
				likers = RemoveLikeByIndex(likers, i)
				liker := models.Like{
					ID:   userid,
					NAME: user.NAME,
					LIKE: false,
				}
				likers = append(likers, liker)
			} else if x.ID == userid && x.LIKE == false {
				likers = RemoveLikeByIndex(likers, i)
				liker := models.Like{
					ID:   userid,
					NAME: user.NAME,
					LIKE: true,
				}
				likers = append(likers, liker)
			} else {
				liker := models.Like{
					ID:   userid,
					NAME: user.NAME,
					LIKE: true,
				}
				likers = append(likers, liker)
			}
		}
	} else {
		liker := models.Like{
			ID:   userid,
			NAME: user.NAME,
			LIKE: true,
		}
		likers = append(likers, liker)
	}

	bootcamp := models.Bootcamp{}
	_, err = dbConnect.Model(&bootcamp).Set("likes = ?", likers).Where("id = ?", id).Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "Something went wrong",
		})
		return
	}

	cachebootcamp, err := json.Marshal(bootcamp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	err = rdbClient.Set("bootcamp"+bootcamp.ID, cachebootcamp, 604800*time.Second).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "You Liked this Bootcamp post!",
	})
	return
}

func CommentBootcamp(c *gin.Context) {
	id := c.Request.URL.Query().Get("id")
	userid := c.Request.URL.Query().Get("userid")
	fmt.Println(userid)
	var commentBody models.Comment
	c.BindJSON(&commentBody)

	resbootcamp := &models.Bootcamp{ID: id}
	err := dbConnect.Select(resbootcamp)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "Bootcamp not found!",
		})
		return
	}

	user := &models.User{ID: userid}
	err = dbConnect.Select(user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "User not found!",
		})
		return
	}

	commenters := resbootcamp.COMMENTS
	commenter := models.Comment{
		ID:        userid,
		NAME:      user.NAME,
		TEXT:      commentBody.TEXT,
		EDITED:    false,
		UPDATEDAT: time.Now(),
	}

	commenters = append(commenters, commenter)

	bootcamp := models.Bootcamp{}
	_, err = dbConnect.Model(&bootcamp).Set("comments = ?", commenters).Where("id = ?", id).Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "Something went wrong",
		})
		return
	}

	cachebootcamp, err := json.Marshal(bootcamp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	err = rdbClient.Set("bootcamp"+bootcamp.ID, cachebootcamp, 604800*time.Second).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Commented to Bootcamp post!",
	})
	return
}

func DeleteBootcamp(c *gin.Context) {
	id := c.Request.URL.Query().Get("id")
	var bootcampBody models.Bootcamp
	c.BindJSON(&bootcampBody)
	bootcamp := &models.Bootcamp{ID: id}

	err := dbConnect.Delete(bootcamp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}
	err = rdbClient.Del("bootcamp" + id).Err()
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Deleted!",
	})
	return
}

func ResetBootcampCache(c *gin.Context) {
	var bootcampList []models.Bootcamp
	var bootcamp models.Bootcamp

	keys := rdbClient.Keys("bootcamp*")
	keyres := keys.Val()

	for _, key := range keyres {
		val, err := rdbClient.Get(key).Result()
		if err != nil {
			c.JSON(http.StatusAccepted, gin.H{
				"msg": "failed to get user from cache",
			})
			return
		}
		err = json.Unmarshal([]byte(val), &bootcamp)
		if bootcamp.AUTHOR != "" {
			bootcampList = append(bootcampList, bootcamp)
		}
	}

	if len(bootcampList) != 0 {
		for _, key := range bootcampList {
			err := rdbClient.Del("bootcamp" + key.ID).Err()
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
