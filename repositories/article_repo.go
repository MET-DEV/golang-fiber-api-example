package repositories

import (
	"github.com/MET-DEV/golang-fiber-api-example/config"
	"github.com/MET-DEV/golang-fiber-api-example/models"
)

func GetAllArticle() []models.Article {
	var articles []models.Article
	config.DB.Preload("Comments").Find(&articles)
	return articles
}
