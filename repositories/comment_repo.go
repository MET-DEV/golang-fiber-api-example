package repositories

import (
	"github.com/MET-DEV/golang-fiber-api-example/config"
	"github.com/MET-DEV/golang-fiber-api-example/models"
)

func GetAllComments() []models.Comment {
	var comments []models.Comment
	config.DB.Find(&comments)
	return comments
}

func GetCommentById(id int) models.Comment {
	var comment models.Comment
	config.DB.First(&comment, id)
	return comment
}

func AddComment(comment models.Comment) models.Comment {
	config.DB.Save(&comment)
	return comment
}
func UpdateComment(comment models.Comment) models.Comment {
	config.DB.Updates(&comment)
	return comment
}

func DeleteComment(id int) {
	var comment models.Comment
	config.DB.First(&comment, id).Delete(&comment)
}
