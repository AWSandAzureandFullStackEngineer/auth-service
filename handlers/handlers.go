package handlers

import (
	"auth-service/models"
	"auth-service/service"
	"auth-service/utils"
	"github.com/gofiber/fiber/v2"
	"log"
)

func Register(c *fiber.Ctx, userService service.UserService) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// Hash the password before saving
	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		log.Println("Error hashing password:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to hash password",
		})
	}
	user.Password = hashedPassword

	if err := userService.RegisterUser(user); err != nil {
		log.Println("Error registering user:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to register user",
		})
	}

	// Generate a JWT token
	token, err := utils.GenerateToken(user.Username)
	if err != nil {
		log.Println("Error generating token:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to generate token",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"token": token,
	})
}

// HashPassword hashes the user's password
func HashPassword(password string) (string, error) {
	// Implement password hashing here
	return password, nil // Placeholder: Replace with actual hash
}
