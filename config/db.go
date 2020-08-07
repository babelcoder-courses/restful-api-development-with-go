package config

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

func InitDB() {
	var err error
	db, err = gorm.Open("postgres", os.Getenv("DATABASE_CONNECTION"))
	if err != nil {
		log.Fatal(err)
	}

	db.LogMode(gin.Mode() == gin.DebugMode)
}

func GetDB() *gorm.DB {
	return db
}

func CloseDB() {
	db.Close()
}
