package handler

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/ramzyrsr/digital-library/internal/entity"
	"github.com/ramzyrsr/digital-library/internal/middleware"
	"github.com/ramzyrsr/digital-library/internal/repository"
	"golang.org/x/crypto/bcrypt"
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

	existingUser, err := h.UserRepo.GetUserByEmail(req.Email)
	if err == nil && existingUser != nil {
		return middleware.Response(c, fiber.StatusConflict, "Email is already registered", nil)
	}

	user := &entity.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
		Role:     req.Role,
	}

	err = h.UserRepo.CreateUser(user)
	if err != nil {
		log.Println("Error creating user:", err)
		return middleware.Response(c, fiber.StatusConflict, "Failed to register user", nil)
	}

	return middleware.Response(c, fiber.StatusCreated, "User registered successfully", nil)
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&req); err != nil {
		return middleware.Response(c, fiber.StatusBadRequest, "Invalid input", nil)
	}

	user, err := h.UserRepo.GetUserByEmail(req.Email)
	if err != nil {
		return middleware.Response(c, fiber.StatusUnauthorized, "Invalid email or password", nil)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return middleware.Response(c, fiber.StatusUnauthorized, "Invalid email or password", nil)
	}

	token, err := middleware.GenerateJWT(user.ID, user.Role)
	if err != nil {
		return middleware.Response(c, fiber.StatusInternalServerError, "Failed to generate token", nil)
	}

	return middleware.Response(c, fiber.StatusOK, token, nil)
}

func (h *AuthHandler) CreateMember(c *fiber.Ctx) error {
	var req struct {
		UserID uint   `json:"user_id"`
		Name   string `json:"name"`
		Email  string `json:"email"`
		Phone  string `json:"phone"`
	}

	if err := c.BodyParser(&req); err != nil {
		return middleware.Response(c, fiber.StatusBadRequest, "Invalid input", nil)
	}

	userID := c.Locals("user_id").(uint)

	user := &entity.Member{
		UserID: userID,
		Name:   req.Name,
		Email:  req.Email,
		Phone:  req.Phone,
	}

	err := h.UserRepo.CreateMember(user)
	if err != nil {
		return middleware.Response(c, fiber.StatusConflict, "Failed to register member", err.Error())
	}

	return middleware.Response(c, fiber.StatusCreated, "Member registered successfully", nil)
}
