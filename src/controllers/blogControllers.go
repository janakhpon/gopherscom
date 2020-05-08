package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/janakhpon/gopherscom/src/models"
)

func GetBlogList(c *gin.Context) {
	var blogList []models.Blog
	err := dbConnect.Model(&blogList).Select()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": blogList,
	})
	return
}

func GetBlog(c *gin.Context) {
	id := c.Request.URL.Query().Get("id")
	blog := &models.Blog{ID: id}
	err := dbConnect.Select(blog)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "failed to fetch",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "succeed",
		"data": blog,
	})
	return
}

func CreateBlog(c *gin.Context) {
	var blogBody models.Blog
	c.BindJSON(&blogBody)

	blog := models.Blog{
		ID:        uuid.New().String(),
		TITLE:     blogBody.TITLE,
		BODY:      blogBody.BODY,
		PUBLIC:    blogBody.PUBLIC,
		APPTYPE:   blogBody.APPTYPE,
		LANGUAGES: blogBody.LANGUAGES,
		TAGS:      blogBody.TAGS,
		LIBRARIES: blogBody.LIBRARIES,
		AUTHOR:    blogBody.AUTHOR,
		CREATEDAT: time.Now(),
		UPDATEDAT: time.Now(),
	}

	insertError := dbConnect.Insert(&blog)
	if insertError != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": insertError,
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "created",
		"data":    &blog,
	})

	return
}

func UpdateBlog(c *gin.Context) {
	id := c.Request.URL.Query().Get("id")
	var blogBody models.Blog
	c.BindJSON(&blogBody)
	reblog := &models.Blog{ID: id}

	err := dbConnect.Select(reblog)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "failed to fetch",
		})
		return
	}

	blog := models.Blog{
		ID:        id,
		TITLE:     blogBody.TITLE,
		BODY:      blogBody.BODY,
		PUBLIC:    blogBody.PUBLIC,
		APPTYPE:   blogBody.APPTYPE,
		LANGUAGES: blogBody.LANGUAGES,
		TAGS:      blogBody.TAGS,
		LIBRARIES: blogBody.LIBRARIES,
		AUTHOR:    blogBody.AUTHOR,
		CREATEDAT: reblog.CREATEDAT,
		UPDATEDAT: time.Now(),
	}
	updateError := dbConnect.Update(&blog)

	if updateError != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": updateError,
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "updated",
		"data":    &blog,
	})
	return
}

func SetBlogPublic(c *gin.Context) {
	id := c.Request.URL.Query().Get("id")
	publicval := c.Request.URL.Query().Get("public")
	var blogBody models.Blog
	c.BindJSON(&blogBody)

	blog := models.Blog{}

	_, err := dbConnect.Model(&blog).Set("public = ?", publicval).Where("id = ?", id).Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Public Value Change",
	})
	return
}

func DeleteBlog(c *gin.Context) {
	id := c.Request.URL.Query().Get("id")
	var blogBody models.Blog
	c.BindJSON(&blogBody)
	blog := &models.Blog{ID: id}

	err := dbConnect.Delete(blog)
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
