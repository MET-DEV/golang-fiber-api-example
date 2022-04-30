package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	CommentMessage string  `json:"comment"`
	ArticleID      uint    `json:"article"`
	Article        Article `gorm:"association_foreignkey:ArticleID" json:"article"`
}
