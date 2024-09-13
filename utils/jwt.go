package utils

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

var secretKey = []byte(os.Getenv("SECRET_KEY"))

// GenerateToken generates a JWT token for a user
func GenerateToken(userID string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(), // Set expiration to 1 day
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
