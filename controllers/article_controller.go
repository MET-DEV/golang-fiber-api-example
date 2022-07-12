package controllers

import (
	"net/http"
	"strconv"

	"github.com/MET-DEV/golang-fiber-api-example/models"
	"github.com/MET-DEV/golang-fiber-api-example/repositories"
	"github.com/gofiber/fiber/v2"
)

func GetAllArticle(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(repositories.GetAllArticle())
}
func GetArticleById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON("Bad Request")
	}
	return c.Status(http.StatusOK).JSON(repositories.GetArticleById(id))
}
func AddArticle(c *fiber.Ctx) error {
	var article models.Article
	err := c.BodyParser(&article)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON("Bad Request")
	}
	return c.Status(http.StatusCreated).JSON(repositories.AddArticle(article))

}
func UpdateArticle(c *fiber.Ctx) error {
	var article models.Article
	err := c.BodyParser(&article)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON("Bad Request")
	}
	if article.ID == 0 {
		return c.Status(http.StatusBadRequest).JSON("Bad Request")
	}
	return c.Status(http.StatusCreated).JSON(repositories.UpdateArticle(article))

}
func DeleteArticle(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON("Bad Request")
	}
	repositories.DeleteArticle(id)
	return c.Status(http.StatusOK).JSON("Deleted")
}
