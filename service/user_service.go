package service

import "auth-service/models"

type UserService interface {
	RegisterUser(user *models.User) error
	LoginUser(username, password string) (string, error)
}
