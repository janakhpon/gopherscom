package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/janakhpon/gopherscom/src/controllers"
)

func ExtRouter(mode string) *gin.Engine {
	gin.SetMode(mode)
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "HELLO FROM GOPHERS COMMUNITY !!!",
		})
	})

	router.POST("/singup", controllers.UserSignup)
	router.POST("/signin", controllers.UserSignin)
	router.GET("/refreshToken", controllers.RefreshToken)

	authedUserOnly := router.Group("/protected/user")
	authedUserOnly.Use(controllers.TokenVerifyMiddleWare())
	{
		authedUserOnly.GET("/list", controllers.GetUserList)
		authedUserOnly.GET("/byid", controllers.GetUser)
	}

	authedBlogOnly := router.Group("/protected/blog")
	authedBlogOnly.Use(controllers.TokenVerifyMiddleWare())
	{
		authedBlogOnly.POST("/new", controllers.CreateBlog)
		authedBlogOnly.GET("/list", controllers.GetBlogList)
		authedBlogOnly.GET("/byid", controllers.GetBlog)
		authedBlogOnly.PUT("/update", controllers.UpdateBlog)
		authedBlogOnly.PUT("/public", controllers.SetBlogPublic)
		authedBlogOnly.DELETE("/remove", controllers.DeleteBlog)
	}

	authedProfileOnly := router.Group("/protected/profile")
	authedProfileOnly.Use(controllers.TokenVerifyMiddleWare())
	{
		authedProfileOnly.GET("/list", controllers.GetProfileList)
		authedProfileOnly.GET("/byuser", controllers.GetProfileByUser)
		authedProfileOnly.GET("/byid", controllers.GetByID)
		authedProfileOnly.POST("/new", controllers.CreateProfile)
		authedProfileOnly.PUT("/update", controllers.UpdateProfile)
	}

	return router
}
