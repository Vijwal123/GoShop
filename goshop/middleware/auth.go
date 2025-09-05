package middleware

import (
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	jwt "github.com/golang-jwt/jwt/v5"
)

func JWTProtected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "missing token"})
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		secret := os.Getenv("JWT_SECRET")
		if secret == "" {
			secret = "defaultsecret"
		}

		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})

		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid token"})
		}

		c.Locals("user", token.Claims.(jwt.MapClaims))
		return c.Next()
	}
}

func AdminOnly() fiber.Handler {
	return func(c *fiber.Ctx) error {
		claims := c.Locals("user").(jwt.MapClaims)
		if claims["role"] != "admin" {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "admin only"})
		}
		return c.Next()
	}
}
