package main

import (
	"course-go/config"
	"course-go/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.InitDB()
	defer config.CloseDB()

	r := gin.Default()
	r.Static("/uploads", "./uploads")

	uploadDirs := [...]string{"articles", "users"}
	for _, dir := range uploadDirs {
		os.MkdirAll("uploads/"+dir, 0755)
	}

	routes.Serve(r)

	r.Run(":" + os.Getenv("PORT"))
}
