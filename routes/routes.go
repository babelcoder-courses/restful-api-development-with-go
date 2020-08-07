package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GET /api/v1/articles
// GET /api/v1/articles/:id
// POST /api/v1/articles
// PATCH /api/v1/articles/:id
// DELETE /api/v1/articles/:id

type article struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

type createArticleForm struct {
	Title string `json:"title" binding:"required"`
	Body  string `json:"body" binding:"required"`
}

func Serve(r *gin.Engine) {
	articles := []article{
		{ID: 1, Title: "Title#1", Body: "Body#1"},
		{ID: 2, Title: "Title#2", Body: "Body#2"},
		{ID: 3, Title: "Title#3", Body: "Body#3"},
		{ID: 4, Title: "Title#4", Body: "Body#4"},
		{ID: 5, Title: "Title#5", Body: "Body#5"},
	}

	articlesGroup := r.Group("/api/v1/articles")
	articlesGroup.GET("", func(ctx *gin.Context) {
		result := articles
		if limit := ctx.Query("limit"); limit != "" {
			n, _ := strconv.Atoi(limit)

			result = result[:n]
		}

		ctx.JSON(http.StatusOK, gin.H{"articles": result})
	})
	articlesGroup.GET("/:id", func(ctx *gin.Context) {
		id, _ := strconv.Atoi(ctx.Param("id"))

		for _, item := range articles {
			if item.ID == uint(id) {
				ctx.JSON(http.StatusOK, gin.H{"article": item})
				return
			}
		}

		ctx.JSON(http.StatusNotFound, gin.H{"error": "Article not found"})
	})
	articlesGroup.POST("", func(ctx *gin.Context) {
		var form createArticleForm
		if err := ctx.ShouldBindJSON(&form); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		a := article{
			ID:    uint(len(articles) + 1),
			Title: form.Title,
			Body:  form.Body,
		}

		articles = append(articles, a)

		ctx.JSON(http.StatusCreated, gin.H{"article": a})
	})
}
