package repository

import "auth-service/models"

type UserRepository interface {
	RegisterUser(user *models.User) error
	FindByUsernameOrEmail(username, email string, user *models.User) error
}
