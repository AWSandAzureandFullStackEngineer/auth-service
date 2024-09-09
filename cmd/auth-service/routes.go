package main

import (
	"auth-service/handlers"
	"auth-service/service"
	"github.com/gofiber/fiber/v2"
)

// SetupRoutes sets up routes for the application
func SetupRoutes(app *fiber.App, userService service.UserService) {
	app.Post("/register", func(c *fiber.Ctx) error {
		return handlers.Register(c, userService)
	})
}
