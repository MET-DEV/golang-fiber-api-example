package migrations

import (
	"github.com/MET-DEV/golang-fiber-api-example/config"
	"github.com/MET-DEV/golang-fiber-api-example/models"
)

func CommentMigration() {
	config.DB.AutoMigrate(&models.Comment{})
}
