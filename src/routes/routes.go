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

	router.POST("/signup", controllers.UserSignup)
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

	authedApptypeOnly := router.Group("/protected/apptype")
	authedApptypeOnly.Use(controllers.TokenVerifyMiddleWare())
	{
		authedApptypeOnly.GET("/list", controllers.GetApptypeList)
		authedApptypeOnly.GET("/byid", controllers.GetApptype)
		authedApptypeOnly.POST("/new", controllers.CreateApptype)
		authedApptypeOnly.PUT("/update", controllers.UpdateApptype)
		authedApptypeOnly.DELETE("/remove", controllers.DeleteApptype)
	}

	authedLibraryOnly := router.Group("/protected/library")
	authedLibraryOnly.Use(controllers.TokenVerifyMiddleWare())
	{
		authedLibraryOnly.GET("/list", controllers.GetLibraryList)
		authedLibraryOnly.GET("/byid", controllers.GetLibrary)
		authedLibraryOnly.POST("/new", controllers.CreateLibrary)
		authedLibraryOnly.PUT("/update", controllers.UpdateLibrary)
		authedLibraryOnly.DELETE("/remove", controllers.DeleteLibrary)
	}

	authedOtherOnly := router.Group("/protected/other")
	authedOtherOnly.Use(controllers.TokenVerifyMiddleWare())
	{
		authedOtherOnly.GET("/list", controllers.GetOtherList)
		authedOtherOnly.GET("/byid", controllers.GetOther)
		authedOtherOnly.POST("/new", controllers.CreateOther)
		authedOtherOnly.PUT("/update", controllers.UpdateOther)
		authedOtherOnly.DELETE("/remove", controllers.DeleteOther)
	}

	authedTagOnly := router.Group("/protected/tag")
	authedTagOnly.Use(controllers.TokenVerifyMiddleWare())
	{
		authedTagOnly.GET("/list", controllers.GetTagList)
		authedTagOnly.GET("/byid", controllers.GetTag)
		authedTagOnly.POST("/new", controllers.CreateTag)
		authedTagOnly.PUT("/update", controllers.UpdateTag)
		authedTagOnly.DELETE("/remove", controllers.DeleteTag)
	}

	CompanyOnly := router.Group("/public/company")
	CompanyOnly.Use()
	{
		CompanyOnly.GET("/list", controllers.GetCompanyList)
		CompanyOnly.POST("/new", controllers.AddCompany)
	}

	BranchOnly := router.Group("/public/branch")
	BranchOnly.Use()
	{
		BranchOnly.POST("/new", controllers.AddCompanyBranch)
		BranchOnly.GET("/branches", controllers.GetCompanyBranches)
		BranchOnly.PUT("/update", controllers.UpdateCompanyBranch)
		BranchOnly.DELETE("/delete", controllers.DeleteCompanyBranch)

	}

	return router
}
