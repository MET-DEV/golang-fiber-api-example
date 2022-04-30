package main

import (
	"fmt"
	"net/http"

	"github.com/MET-DEV/golang-fiber-api-example/config"
	"github.com/MET-DEV/golang-fiber-api-example/migrations"
	"github.com/MET-DEV/golang-fiber-api-example/repositories"
	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println("Server started...")
	config.DbConfiguration()
	migrations.IndexMigration()

	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).JSON(repositories.GetAllArticle())
	})

	app.Listen(":3000")
}
