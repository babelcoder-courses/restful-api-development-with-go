package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name     string `gorm:"unique;not null"`
	Desc     string `gorm:"not null"`
	Articles []Article
}
