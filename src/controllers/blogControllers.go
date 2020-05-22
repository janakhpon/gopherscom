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

	val, err := rdbClient.Get(id).Result()
	if err != nil {
		c.JSON(http.StatusAccepted, gin.H{
			"msg": "failed to get user from cache",
		})
		return
	}
	err = json.Unmarshal([]byte(val), &blog)
	if blog != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "succeed",
			"data": blog,
		})
		return
	}
	err = dbConnect.Select(blog)
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

	cacheblog, err := json.Marshal(blog)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	err = rdbClient.Set(blog.ID, cacheblog, 604800*time.Second).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
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
	cacheblog, err := json.Marshal(blog)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	err = rdbClient.Set(blog.ID, cacheblog, 604800*time.Second).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
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
	cacheblog, err := json.Marshal(blog)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	err = rdbClient.Set(blog.ID, cacheblog, 604800*time.Second).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
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

	err = rdbClient.Del(id).Err()
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Deleted!",
	})
	return
}

func LikeBlog(c *gin.Context) {
	id := c.Request.URL.Query().Get("id")
	userid := c.Request.URL.Query().Get("userid")
	fmt.Println(userid)
	var blogBody models.Blog
	c.BindJSON(&blogBody)

	resblog := &models.Blog{ID: id}
	err := dbConnect.Select(resblog)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "blog not found!",
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

	likers := resblog.LIKES
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

	blog := models.Blog{}
	_, err = dbConnect.Model(&blog).Set("likes = ?", likers).Where("id = ?", id).Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "Something went wrong",
		})
		return
	}

	cacheblog, err := json.Marshal(blog)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	err = rdbClient.Set(blog.ID, cacheblog, 604800*time.Second).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "You Liked this blog post!",
	})
	return
}

func CommentBLog(c *gin.Context) {
	id := c.Request.URL.Query().Get("id")
	userid := c.Request.URL.Query().Get("userid")
	fmt.Println(userid)
	var commentBody models.Comment
	c.BindJSON(&commentBody)

	resblog := &models.Blog{ID: id}
	err := dbConnect.Select(resblog)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "blog not found!",
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

	commenters := resblog.COMMENTS
	commenter := models.Comment{
		ID:        userid,
		NAME:      user.NAME,
		TEXT:      commentBody.TEXT,
		EDITED:    false,
		UPDATEDAT: time.Now(),
	}

	commenters = append(commenters, commenter)

	blog := models.Blog{}
	_, err = dbConnect.Model(&blog).Set("comments = ?", commenters).Where("id = ?", id).Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "Something went wrong",
		})
		return
	}
	cacheblog, err := json.Marshal(blog)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	err = rdbClient.Set(blog.ID, cacheblog, 604800*time.Second).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Commented to blog post!",
	})
	return
}
