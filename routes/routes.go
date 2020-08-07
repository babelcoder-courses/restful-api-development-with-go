package routes

import (
	"course-go/config"
	"course-go/controllers"

	"github.com/gin-gonic/gin"
)

func Serve(r *gin.Engine) {
	db := config.GetDB()
	articlesGroup := r.Group("/api/v1/articles")
	articleController := controllers.Articles{DB: db}
	{
		articlesGroup.GET("", articleController.FindAll)
		articlesGroup.GET("/:id", articleController.FindOne)
		articlesGroup.POST("", articleController.Create)
	}
}
