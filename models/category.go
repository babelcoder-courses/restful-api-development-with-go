package models

import "github.com/jinzhu/gorm"

type Category struct {
	gorm.Model
	Name     string `gorm:"unique;not null"`
	Desc     string `gorm:"not null"`
	Articles []Article
}
