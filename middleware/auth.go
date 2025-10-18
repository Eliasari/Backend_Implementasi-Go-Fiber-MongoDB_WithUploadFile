package middleware

import (
	"fmt"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// AuthRequired middleware untuk verifikasi JWT
func AuthRequired() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(strings.ToLower(authHeader), "bearer ") {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Token tidak ditemukan",
			})
		}

		tokenStr := strings.TrimSpace(authHeader[len("Bearer "):])
		secret := os.Getenv("JWT_SECRET")

		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})
		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Token tidak valid",
			})
		}

		claims := token.Claims.(jwt.MapClaims)

		// ✅ Fix utama: handle user_id bisa string atau number
		userIDRaw := claims["user_id"]
		var userID string

		switch v := userIDRaw.(type) {
		case string:
			userID = v
		case float64:
			userID = fmt.Sprintf("%.0f", v)
		default:
			userID = fmt.Sprint(v)
		}

		// Simpan ke Fiber context
		c.Locals("user_id", userID)

		if username, ok := claims["username"].(string); ok {
			c.Locals("username", username)
		}
		if role, ok := claims["role"].(string); ok {
			c.Locals("role", role)
		}

		return c.Next()
	}
}

// AdminOnly middleware untuk role admin
func AdminOnly() fiber.Handler {
	return func(c *fiber.Ctx) error {
		role, _ := c.Locals("role").(string)
		if role != "admin" {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": "Akses ditolak. Hanya admin yang bisa mengakses",
			})
		}
		return c.Next()
	}
}
