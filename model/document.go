package model

import "github.com/jinzhu/gorm"

type Document struct {
	gorm.Model
	UserID      int    `gorm:"not null"`
	Title       string `gorm:"size:255;nut null"`
	Description string `gorm:"not null"`
	FilePath    string `gorm:"size:255;not null"`
}
