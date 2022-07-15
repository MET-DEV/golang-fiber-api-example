package controllers

import (
	"net/http"
	"strconv"

	"github.com/MET-DEV/golang-fiber-api-example/models"
	"github.com/MET-DEV/golang-fiber-api-example/repositories"
	"github.com/gofiber/fiber/v2"
)

func GetAllComment(c *fiber.Ctx) error {
	result, err := repositories.GetAllComments()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}
	return c.Status(http.StatusOK).JSON(result)
}
func GetCommentById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON("Bad Request")
	}
	result, repoErr := repositories.GetCommentById(id)
	if repoErr != nil {
		return c.Status(http.StatusInternalServerError).JSON(repoErr.Error())
	}
	return c.Status(http.StatusOK).JSON(result)
}
func AddComment(c *fiber.Ctx) error {
	var comment models.Comment
	err := c.BodyParser(&comment)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON("Bad Request")
	}
	result, repoErr := repositories.AddComment(comment)
	if repoErr != nil {
		c.Status(http.StatusInternalServerError).JSON(repoErr.Error())
	}
	return c.Status(http.StatusCreated).JSON(result)

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
	result, repoErr := repositories.UpdateComment(comment)
	if repoErr != nil {
		return c.Status(http.StatusInternalServerError).JSON(repoErr.Error())
	}
	return c.Status(http.StatusCreated).JSON(result)

}
func DeleteComment(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON("Bad Request")
	}
	repoErr := repositories.DeleteComment(id)
	if repoErr != nil {
		return c.Status(http.StatusInternalServerError).JSON(repoErr.Error())
	}
	return c.Status(http.StatusOK).JSON("Deleted")
}
