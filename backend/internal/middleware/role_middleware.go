package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func StaffOnlyMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := c.Locals("user").(*jwt.Token)
		claims, ok := user.Claims.(jwt.MapClaims)
		if !ok {
			return Response(c, fiber.StatusUnauthorized, "Invalid token", nil)
		}

		// Check role
		role, exists := claims["role"].(string)
		if !exists || role != "staff" {
			return Response(c, fiber.StatusForbidden, "Access denied. Staff only.", nil)
		}

		return c.Next()
	}
}
