package handlers

import (
	"auth-service/database"
	"auth-service/models"
	"github.com/gofiber/fiber/v2"
)

func Home(c *fiber.Ctx) error {
	return c.SendString("Yo 👋!")
}

func Register(c *fiber.Ctx) error {
	user := new(models.User)

	// Parse the request body into the user model
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// Save the user to the database
	if result := database.DB.Db.Create(user); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": result.Error.Error(),
		})
	}

	// Return the created user with a 201 Created status
	return c.Status(fiber.StatusCreated).JSON(user)
}
