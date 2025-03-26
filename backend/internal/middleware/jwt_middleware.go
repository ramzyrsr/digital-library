package middleware

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func JWTMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return Response(c, fiber.StatusUnauthorized, "Authorization header missing", nil)
		}

		tokenString := authHeader[7:] // Skip "Bearer " part

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fiber.NewError(fiber.StatusUnauthorized, "Invalid signing method")
			}
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil {
			return Response(c, fiber.StatusUnauthorized, "Unauthorized", err.Error())
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return Response(c, fiber.StatusUnauthorized, "Invalid token claims", nil)
		}

		userID, ok := claims["user_id"].(float64)
		if !ok {
			return Response(c, fiber.StatusUnauthorized, "User ID not found in token", nil)
		}

		c.Locals("user_id", uint(userID))
		c.Locals("user", token)

		return c.Next()
	}
}
