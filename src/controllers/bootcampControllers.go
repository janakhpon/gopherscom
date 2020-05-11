package controllers

import (
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
	err := dbConnect.Model(&bootcampList).Select()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": bootcampList,
	})
	return
}

func GetBootcamp(c *gin.Context) {
	id := c.Request.URL.Query().Get("id")
	bootcamp := &models.Bootcamp{ID: id}
	err := dbConnect.Select(bootcamp)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "failed to fetch",
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

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Deleted!",
	})
	return
}
