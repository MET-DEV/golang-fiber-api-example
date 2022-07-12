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
func GetArticleById(id int) models.Article {
	var article models.Article
	config.DB.Preload("Comments").First(&article, id)
	return article
}

func AddArticle(article models.Article) models.Article {
	config.DB.Save(&article)
	return article
}
func UpdateArticle(article models.Article) models.Article {
	config.DB.Save(&article)
	return article
}

func DeleteArticle(id int) {
	var article models.Article
	config.DB.First(&article, id).Delete(&article)
}
