package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName   string `json:"first_name" gorm:"text;not null;default:null"`
	LastName    string `json:"last_name" gorm:"text;not null;default:null"`
	PhoneNumber string `json:"phone_number"`
	Username    string `json:"username" gorm:"text;not null;default:null"`
	Email       string `json:"email" gorm:"text;not null;default:null"`
	Password    string `json:"password" gorm:"text;not null;default:null"`
}
