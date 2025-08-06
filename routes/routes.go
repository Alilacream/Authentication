package routes

import (
	"Auth/controller"

	"github.com/gofiber/fiber/v2"
)

func Setup(router *fiber.App) {
	router.Get("/", controller.Greetings)
	router.Post("/api/register", controller.Register)
	router.Post("/api/login", controller.Login)
}
