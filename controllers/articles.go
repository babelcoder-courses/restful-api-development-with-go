package controllers

import (
	"mime/multipart"

	"github.com/gin-gonic/gin"
)

type Articles struct {
}

type createArticleForm struct {
	Title string                `form:"title" binding:"required"`
	Body  string                `form:"body" binding:"required"`
	Image *multipart.FileHeader `form:"image" binding:"required"`
}

func (a *Articles) FindAll(ctx *gin.Context) {

}

func (a *Articles) FindOne(ctx *gin.Context) {

}

func (a *Articles) Create(ctx *gin.Context) {

}
