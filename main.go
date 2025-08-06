package main

import (
	"Auth/database"
	"Auth/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Connection()
	app := fiber.New()

	routes.Setup(app)

	app.Listen(":8080")
}
