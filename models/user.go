package models

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	Email    string `gorm:"unique_index; not null"`
	Password string `gorm:"not null"`
	Name     string `gorm:"not null"`
	Avatar   string
	Role     string `gorm:"default:'Member'; not null"`
}

func (u *User) BeforeSave(scope *gorm.Scope) {
	if u.Password == "" {
		return
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
	scope.SetColumn("password", string(hash))
}
