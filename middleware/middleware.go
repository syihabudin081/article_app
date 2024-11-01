package middleware

import (
	jwt2 "OrdentTest/utils/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"strings"
)

// AuthMiddleware ensures the request is authenticated
func AuthMiddleware(c *fiber.Ctx) error {
	// Extract token from header
	tokenString := c.Get("Authorization")
	if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Missing or malformed JWT",
		})
	}
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	// Validate token
	claims, err := jwt2.ValidateToken(tokenString)
	if err != nil {
		log.Printf("Token validation error: %v\n", err)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Set user claims and safely retrieve author ID
	c.Locals("userClaims", claims)

	userID, ok := claims["userId"].(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "User ID not found in claims",
		})
	}
	c.Locals("userId", userID)
	c.Locals("authorID", userID)

	return c.Next()
}

// AuthorMiddleware ensures the user has an Author role
func AuthorMiddleware(c *fiber.Ctx) error {
	claims, ok := c.Locals("userClaims").(jwt.MapClaims)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "User claims not found",
		})
	}

	role, ok := claims["role"].(string)
	if !ok {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "Role not found in claims",
			"role":    role,
		})
	}
	log.Println(role)

	if role != "Author" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "Access restricted to authors only",
			"role":    role,
		})
	}

	return c.Next()
}

// AdminMiddleware ensures the user has an Admin role
func AdminMiddleware(c *fiber.Ctx) error {
	// Check the user role from the JWT claims
	claims, ok := c.Locals("userClaims").(jwt.MapClaims)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "User claims not found",
		})
	}

	role, ok := claims["role"].(string)
	if !ok {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "Role not found in claims",
		})
	}

	if role != "Admin" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "Access restricted to admins only",
		})
	}

	return c.Next()
}
