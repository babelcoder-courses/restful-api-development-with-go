package main

import (
	"course-go/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	routes.Serve(r)

	r.Run()
}
