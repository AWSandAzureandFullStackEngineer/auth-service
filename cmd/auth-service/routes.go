package main

import (
	"auth-service/handlers"
	"auth-service/service"
	"auth-service/utils"
	"github.com/gofiber/fiber/v2"
	_ "github.com/swaggo/fiber-swagger"
	"strings"
)

// SetupRoutes sets up routes for the application
func SetupRoutes(app *fiber.App, userService service.UserService) {
	app.Post("/register", func(c *fiber.Ctx) error {
		return handlers.Register(c, userService)
	})

	app.Post("/login", func(c *fiber.Ctx) error {
		return handlers.Login(c, userService)
	})

	// Middleware to protect the Swagger route
	app.Use("/swagger.json", func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		_, err := utils.VerifyToken(token)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
		}

		return c.SendFile("./swagger.json")
	})
}
