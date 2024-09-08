package main

import (
	"auth-service/handlers"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.Home)

	app.Post("/users", handlers.Register)

}
