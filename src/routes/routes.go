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
	router.POST("/new", controllers.CreateBlog)
	router.GET("/list", controllers.GetBlogList)
	router.GET("/byid", controllers.GetBlog)
	router.PUT("/update", controllers.UpdateBlog)
	router.PUT("/public", controllers.SetBlogPublic)
	router.DELETE("/remove", controllers.DeleteBlog)
	router.GET("/profilelist", controllers.GetProfileList)
	router.GET("/profilebyuser", controllers.GetProfileByUser)
	router.GET("/profilebyid", controllers.GetByID)
	router.POST("/profile", controllers.CreateProfile)
	router.PUT("/updateprofile", controllers.UpdateProfile)
	return router
}
