package controllers

import "github.com/gofiber/fiber/v2"

func GetPaths() *fiber.App {
	app := fiber.New()
	app.Get("/", GetAllArticle)
	return app
}
