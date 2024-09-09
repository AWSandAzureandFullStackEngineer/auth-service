package repository

import (
	"auth-service/models"
	"gorm.io/gorm"
)

// userRepository implements the UserRepository interface
type userRepository struct {
	db *gorm.DB
}

// NewUserRepository creates a new instance of userRepository
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

// RegisterUser saves a new user to the database
func (r *userRepository) RegisterUser(user *models.User) error {
	result := r.db.Create(user)
	return result.Error
}

// FindByUsernameOrEmail checks if a user with the given username or email exists
func (r *userRepository) FindByUsernameOrEmail(username, email string, user *models.User) error {
	return r.db.Where("username = ? OR email = ?", username, email).First(user).Error
}
