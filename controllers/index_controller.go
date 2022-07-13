package controllers

import "github.com/gofiber/fiber/v2"

func GetPaths() *fiber.App {
	app := fiber.New()
	app.Get("api/v1/articles/", GetAllArticle)
	app.Get("api/v1/articles/:id", GetArticleById)
	app.Post("api/v1/articles/", AddArticle)
	app.Put("api/v1/articles/update", UpdateArticle)
	app.Delete("api/v1/articles/:id", DeleteArticle)
	//! <----------------------------Comments------------------------------>
	app.Get("api/v1/comments/", GetAllComment)
	app.Get("api/v1/comments/:id", GetCommentById)
	app.Post("api/v1/comments/", AddComment)
	app.Put("api/v1/comments/update", UpdateComment)
	app.Delete("api/v1/comments/:id", DeleteComment)
	return app
}
