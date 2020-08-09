package models

import "github.com/jinzhu/gorm"

type Article struct {
	gorm.Model
	Title      string `gorm:"unique;not null"`
	Excerpt    string `gorm:"not null"`
	Body       string `gorm:"not null"`
	Image      string `gorm:"not null"`
	CategoryID uint
	Category   Category
}
