package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GET /api/v1/articles
// GET /api/v1/articles/:id
// POST /api/v1/articles
// PATCH /api/v1/articles/:id
// DELETE /api/v1/articles/:id

type article struct {
	Title string
	Body  string
}

func Serve(r *gin.Engine) {
	articles := []article{
		{Title: "Title#1", Body: "Body#1"},
		{Title: "Title#2", Body: "Body#2"},
		{Title: "Title#3", Body: "Body#3"},
		{Title: "Title#4", Body: "Body#4"},
		{Title: "Title#5", Body: "Body#5"},
	}
	articlesGroup := r.Group("/api/v1/articles")
	articlesGroup.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"articles": articles})
	})
	articlesGroup.GET("/:id", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Hello, Thailand")
	})
}
