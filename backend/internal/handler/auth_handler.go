package handler

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/ramzyrsr/digital-library/internal/entity"
	"github.com/ramzyrsr/digital-library/internal/middleware"
	"github.com/ramzyrsr/digital-library/internal/repository"
)

type AuthHandler struct {
	UserRepo *repository.UserRepository
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var req struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
		Role     string `json:"role"`
	}

	if err := c.BodyParser(&req); err != nil {
		return middleware.Response(c, fiber.StatusBadRequest, "Invalid input", nil)
	}

	if req.Name == "" || req.Email == "" || req.Password == "" || req.Role == "" {
		return middleware.Response(c, fiber.StatusBadRequest, "All fields are required", nil)
	}

	user := &entity.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
		Role:     req.Role,
	}

	err := h.UserRepo.CreateUser(user)
	if err != nil {
		log.Println("Error creating user:", err)
		return middleware.Response(c, fiber.StatusConflict, "Failed to register user", nil)
	}

	return middleware.Response(c, fiber.StatusCreated, "User registered successfully", nil)
}
