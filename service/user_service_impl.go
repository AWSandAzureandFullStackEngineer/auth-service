package service

import (
	"auth-service/models"
	"auth-service/repository"
	"auth-service/utils"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
)

type userService struct {
	userRepo repository.UserRepository
}

// NewUserService creates a new instance of userService
func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

// RegisterUser handles user registration logic, including password hashing and uniqueness checks
func (s *userService) RegisterUser(user *models.User) error {
	// Check if username or email already exists
	existingUser := &models.User{}
	if err := s.userRepo.FindByUsernameOrEmail(user.Username, user.Email, existingUser); err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	if existingUser.ID != 0 {
		return fmt.Errorf("username or email already in use")
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Error hashing password:", err)
		return err
	}
	user.Password = string(hashedPassword)

	// Save the new user
	return s.userRepo.RegisterUser(user)
}

// LoginUser handles user login logic
func (s *userService) LoginUser(username, password string) (string, error) {
	// Find user by username
	user := &models.User{}
	if err := s.userRepo.FindByUsernameOrEmail(username, "", user); err != nil {
		return "", err
	}

	// Check password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", fmt.Errorf("invalid credentials")
	}

	// Generate JWT token
	token, err := utils.GenerateToken(fmt.Sprintf("%d", user.ID))
	if err != nil {
		return "", err
	}

	return token, nil
}
