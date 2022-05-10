package main

import (
	"fmt"

	"github.com/MET-DEV/golang-fiber-api-example/config"
	"github.com/MET-DEV/golang-fiber-api-example/controllers"
	"github.com/MET-DEV/golang-fiber-api-example/migrations"
)

func main() {
	fmt.Println("Server started...")
	config.DbConfiguration()
	migrations.IndexMigration()

	app := controllers.GetPaths()

	app.Listen(":3000")
}
