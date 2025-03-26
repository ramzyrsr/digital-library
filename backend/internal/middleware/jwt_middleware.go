package middleware

import (
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func JWTMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return Response(c, fiber.StatusUnauthorized, "Authorization header missing", nil)
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			return Response(c, fiber.StatusUnauthorized, "Invalid Authorization header format", nil)
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

func GenerateJWT(userID int, role string) (string, error) {
	secret := os.Getenv("JWT_SECRET")
	claims := jwt.MapClaims{
		"user_id": userID,
		"role":    role,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}
