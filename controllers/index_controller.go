package controllers

import "github.com/gofiber/fiber/v2"

func GetPaths() *fiber.App {
	app := fiber.New()
	app.Get("/", GetAllArticle)
	app.Get("/:id", GetArticleById)
	app.Post("/", AddArticle)
	app.Put("/update", UpdateArticle)
	app.Delete("/:id", DeleteArticle)
	return app
}
