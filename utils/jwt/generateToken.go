package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

// JWT secret key
var jwtSecret = []byte("your_secret_key") // Replace with a more secure key in production

// GenerateToken generates a new JWT token
func GenerateToken(email string, role string, userID string) (string, error) {
	claims := jwt.MapClaims{
		"email":  email,
		"role":   role,
		"userId": userID,
		"exp":    time.Now().Add(time.Hour * 72).Unix(), // Token expiration time
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
