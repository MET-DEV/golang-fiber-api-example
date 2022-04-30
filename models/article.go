package models

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Comments    []Comment `gorm:"foreignkey:ArticleID"`
}
