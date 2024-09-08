package main

import (
	"auth-service/database"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	database.ConnectDb()
	// Initialize a new Fiber app
	app := fiber.New()

	setupRoutes(app)

	log.Fatal(app.Listen(":8080"))
}
