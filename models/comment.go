package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	CommentMessage string  `json:"comment"`
	ArticleID      uint    `json:"articleId"`
	Article        Article `gorm:"association_foreignkey:ArticleID" json:"article"`
}
