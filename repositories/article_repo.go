package repositories

import (
	"github.com/MET-DEV/golang-fiber-api-example/config"
	"github.com/MET-DEV/golang-fiber-api-example/models"
)

func GetAllArticles() ([]models.Article, error) {
	var articles []models.Article
	if result := config.DB.Preload("Comments").Find(&articles); result.Error != nil {
		return articles, result.Error
	}
	return articles, nil
}
func GetArticleById(id int) (models.Article, error) {
	var article models.Article
	if result := config.DB.Preload("Comments").First(&article, id); result.Error != nil {
		return article, result.Error
	}
	return article, nil
}

func AddArticle(article models.Article) (models.Article, error) {
	if result := config.DB.Save(&article); result.Error != nil {
		return article, result.Error
	}

	return article, nil
}
func UpdateArticle(article models.Article) (models.Article, error) {
	if result := config.DB.Updates(&article); result.Error != nil {
		return article, result.Error
	}

	return article, nil
}

func DeleteArticle(id int) error {
	var article models.Article
	if result := config.DB.First(&article, id).Delete(&article); result.Error != nil {
		return result.Error
	}
	return nil
}
