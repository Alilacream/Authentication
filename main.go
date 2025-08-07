package main

import (
	"Auth/database"
	"Auth/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.Connection()
	app := fiber.New()
	// this is important, because this allows the frontend to allow
	// the cookie to on credit BUT, in the local serv, because we cannot
	// allow the sensitive cookie to be in another serv, making this cookie secure
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:8080",
		AllowCredentials: true,
	}))

	routes.Setup(app)

	app.Listen(":8080")
}
