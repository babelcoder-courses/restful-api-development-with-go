package routes

import (
	"course-go/config"
	"course-go/controllers"
	"course-go/middleware"

	"github.com/gin-gonic/gin"
)

func Serve(r *gin.Engine) {
	db := config.GetDB()
	v1 := r.Group("/api/v1")
	authenticate := middleware.Authenticate().MiddlewareFunc()

	authGroup := v1.Group("auth")
	authController := controllers.Auth{DB: db}
	{
		authGroup.POST("/sign-up", authController.Signup)
		authGroup.POST("/sign-in", middleware.Authenticate().LoginHandler)
		authGroup.GET("/profile", authenticate, authController.GetProfile)
		authGroup.PATCH("/profile", authenticate, authController.UpdateProfile)
	}

	articlesGroup := v1.Group("articles")
	articleController := controllers.Articles{DB: db}
	{
		articlesGroup.GET("", articleController.FindAll)
		articlesGroup.GET("/:id", articleController.FindOne)
		articlesGroup.PATCH("/:id", articleController.Update)
		articlesGroup.DELETE("/:id", articleController.Delete)
		articlesGroup.POST("", authenticate, articleController.Create)
	}

	categoriesGroup := v1.Group("categories")
	categoryController := controllers.Categories{DB: db}
	{
		categoriesGroup.GET("", categoryController.FindAll)
		categoriesGroup.GET("/:id", categoryController.FindOne)
		categoriesGroup.PATCH("/:id", categoryController.Update)
		categoriesGroup.DELETE("/:id", categoryController.Delete)
		categoriesGroup.POST("", categoryController.Create)
	}
}
