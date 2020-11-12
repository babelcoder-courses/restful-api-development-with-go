package migrations

import (
	"course-go/models"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func m1596977447CreateUsersTable() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "1596977447",
		Migrate: func(tx *gorm.DB) error {

			type user struct {
				gorm.Model
				Email    string `gorm:"uniqueIndex; not null"`
				Password string `gorm:"not null"`
				Name     string `gorm:"not null"`
				Avatar   string
				Role     models.Role `gorm:"default:3; not null"`
			}

			return tx.Migrator().CreateTable(&user{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("users")
		},
	}
}
