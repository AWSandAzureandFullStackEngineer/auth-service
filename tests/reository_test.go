package tests

import (
	"auth-service/database"
	"auth-service/models"
	"auth-service/repository"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRegisterUser(t *testing.T) {
	// Initialize the database connection for tests
	database.ConnectDb()
	db := database.DB.Db
	userRepo := repository.NewUserRepository(db)

	// Define a user with all required fields
	user := &models.User{
		FirstName:   "John",
		LastName:    "Doe",
		PhoneNumber: "1234567890",
		Username:    "johndoe",
		Email:       "john@example.com",
		Password:    "securepassword",
	}

	// Attempt to register the user
	err := userRepo.RegisterUser(user)

	// Assert that no error occurred
	assert.NoError(t, err)
}
