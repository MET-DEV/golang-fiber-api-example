package repositories

import (
	"github.com/MET-DEV/golang-fiber-api-example/config"
	"github.com/MET-DEV/golang-fiber-api-example/models"
)

func GetAllComments() ([]models.Comment, error) {
	var comments []models.Comment
	if result := config.DB.Find(&comments); result.Error != nil {
		return comments, result.Error
	}
	return comments, nil
}

func GetCommentById(id int) (models.Comment, error) {
	var comment models.Comment
	if result := config.DB.First(&comment, id); result.Error != nil {
		return comment, result.Error
	}
	return comment, nil
}

func AddComment(comment models.Comment) (models.Comment, error) {
	if result := config.DB.Save(&comment); result.Error != nil {
		return comment, result.Error
	}
	return comment, nil
}
func UpdateComment(comment models.Comment) (models.Comment, error) {
	if result := config.DB.Updates(&comment); result.Error != nil {
		return comment, result.Error
	}
	return comment, nil
}

func DeleteComment(id int) error {
	var comment models.Comment
	if result := config.DB.First(&comment, id).Delete(&comment); result.Error != nil {
		return result.Error
	}
	return nil
}
