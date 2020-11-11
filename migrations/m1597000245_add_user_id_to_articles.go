package migrations

import (
	"course-go/models"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func m1597000245AddUserIDToArticles() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "1597000245",
		Migrate: func(tx *gorm.DB) error {
			return tx.Migrator().AddColumn(&models.Article{}, "user_id")
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropColumn(&models.Article{}, "user_id")
		},
	}
}
