package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"size:255;unique;not null"`
	Password string `gorm:"size:255;not null"`
	Email    string `gorm:"size:255;unique;not null"`
}
