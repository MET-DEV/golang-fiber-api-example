package controllers

import (
	"net/http"
	"strconv"

	"github.com/MET-DEV/golang-fiber-api-example/models"
	"github.com/MET-DEV/golang-fiber-api-example/repositories"
	"github.com/gofiber/fiber/v2"
)

func GetAllArticle(c *fiber.Ctx) error {
	result, repoErr := repositories.GetAllArticles()
	if repoErr != nil {
		return c.Status(http.StatusInternalServerError).JSON(repoErr.Error())
	}
	return c.Status(http.StatusOK).JSON(result)

}
func GetArticleById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON("Bad Request")
	}
	result, repoErr := repositories.GetArticleById(id)
	if repoErr != nil {
		return c.Status(http.StatusInternalServerError).JSON(repoErr.Error())
	}
	return c.Status(http.StatusOK).JSON(result)
}
func AddArticle(c *fiber.Ctx) error {
	var article models.Article
	err := c.BodyParser(&article)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON("Bad Request")
	}
	result, repoErr := repositories.AddArticle(article)
	if repoErr != nil {
		return c.Status(http.StatusInternalServerError).JSON(repoErr.Error())
	}
	return c.Status(http.StatusOK).JSON(result)

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
	result, repoErr := repositories.UpdateArticle(article)
	if repoErr != nil {
		return c.Status(http.StatusInternalServerError).JSON(repoErr.Error())
	}
	return c.Status(http.StatusOK).JSON(result)

}
func DeleteArticle(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON("Bad Request")
	}
	repoErr := repositories.DeleteArticle(id)
	if repoErr != nil {
		return c.Status(http.StatusInternalServerError).JSON(repoErr.Error())
	}
	return c.Status(http.StatusOK).JSON("Deleted")
}
