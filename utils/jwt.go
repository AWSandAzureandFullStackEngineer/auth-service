// utils/jwt.go
package utils

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var secretKey = []byte("your_secret_key")

// GenerateToken generates a JWT token for a user
func GenerateToken(userID string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 2).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

// VerifyToken verifies the JWT token
func VerifyToken(tokenStr string) (*jwt.Token, error) {
	return jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidToken
		}
		return secretKey, nil
	})
}

// ErrInvalidToken indicates an invalid token
var ErrInvalidToken = errors.New("invalid token")
