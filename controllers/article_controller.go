package controllers

import (
	"net/http"

	"github.com/MET-DEV/golang-fiber-api-example/repositories"
	"github.com/gofiber/fiber/v2"
)

func GetAllArticle(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(repositories.GetAllArticle())
}
