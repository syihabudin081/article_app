package jwt

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
)

// ValidateToken validates the JWT token and returns claims
func ValidateToken(tokenString string) (jwt.MapClaims, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})

	// Jika terjadi error atau token tidak valid, langsung return error
	if err != nil || token == nil || !token.Valid {
		return nil, fmt.Errorf("invalid token: %v", err)
	}

	// Konversi claims ke jwt.MapClaims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("unable to parse claims")
	}

	return claims, nil
}
