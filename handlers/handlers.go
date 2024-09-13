package handlers

import (
	"auth-service/models"
	"auth-service/service"
	_ "auth-service/utils"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

// Register handles user registration
func Register(c *fiber.Ctx, userService service.UserService) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}
	if err := userService.RegisterUser(&user); err != nil {
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}
	return c.SendStatus(http.StatusCreated)
}

// Login handles user login
func Login(c *fiber.Ctx, userService service.UserService) error {
	var request struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&request); err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	token, err := userService.LoginUser(request.Username, request.Password)
	if err != nil {
		return c.Status(http.StatusUnauthorized).SendString(err.Error())
	}

	return c.JSON(fiber.Map{
		"token": token,
	})
}
