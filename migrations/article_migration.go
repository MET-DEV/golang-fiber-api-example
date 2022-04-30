package migrations

import (
	"github.com/MET-DEV/golang-fiber-api-example/config"
	"github.com/MET-DEV/golang-fiber-api-example/models"
)

func ArticleMigration() {
	config.DB.AutoMigrate(&models.Article{})
}
