package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func m1596813596CreateArticlesTable() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "1596813596",
		Migrate: func(tx *gorm.DB) error {
			type article struct {
				gorm.Model
				Title   string `gorm:"unique;not null"`
				Excerpt string `gorm:"not null"`
				Body    string `gorm:"not null"`
				Image   string `gorm:"not null"`
			}

			return tx.Migrator().CreateTable(&article{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("articles")
		},
	}
}
