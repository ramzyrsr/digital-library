package middleware

import (
	"github.com/gofiber/fiber/v2"
)

type ResponseFormat struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func Response(c *fiber.Ctx, statusCode int, message string, data interface{}) error {
	var status string
	if statusCode >= 200 && statusCode < 300 {
		status = "success" // 2xx success codes
	} else {
		status = "error" // 4xx and 5xx error codes
	}

	response := ResponseFormat{
		Status:  status,
		Message: message,
		Data:    data,
	}

	return c.Status(statusCode).JSON(response)
}
