package service

import "auth-service/models"

type UserService interface {
	RegisterUser(user *models.User) error
}
