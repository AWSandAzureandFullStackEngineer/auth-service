package main

import (
	"auth-service/database"
	"auth-service/repository"
	"auth-service/service"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	// Initialize database connection
	database.ConnectDb()

	// Create repository and service instances
	userRepo := repository.NewUserRepository(database.DB.Db)
	userService := service.NewUserService(userRepo)

	// Initialize Fiber app
	app := fiber.New()

	// Setup routes with userService
	SetupRoutes(app, userService)

	// Start the application
	log.Fatal(app.Listen(":8080"))
}
