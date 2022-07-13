package controllers

import (
	"net/http"
	"strconv"

	"github.com/MET-DEV/golang-fiber-api-example/models"
	"github.com/MET-DEV/golang-fiber-api-example/repositories"
	"github.com/gofiber/fiber/v2"
)

func GetAllComment(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(repositories.GetAllComments())
}
func GetCommentById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON("Bad Request")
	}
	return c.Status(http.StatusOK).JSON(repositories.GetCommentById(id))
}
func AddComment(c *fiber.Ctx) error {
	var comment models.Comment
	err := c.BodyParser(&comment)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON("Bad Request")
	}
	return c.Status(http.StatusCreated).JSON(repositories.AddComment(comment))

}
func UpdateComment(c *fiber.Ctx) error {
	var comment models.Comment
	err := c.BodyParser(&comment)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON("Bad Request")
	}
	if comment.ID == 0 {
		return c.Status(http.StatusBadRequest).JSON("Bad Request")
	}
	return c.Status(http.StatusCreated).JSON(repositories.UpdateComment(comment))

}
func DeleteComment(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON("Bad Request")
	}
	repositories.DeleteComment(id)
	return c.Status(http.StatusOK).JSON("Deleted")
}
