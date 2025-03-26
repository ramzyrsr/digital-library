package handler

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/ramzyrsr/digital-library/internal/entity"
	"github.com/ramzyrsr/digital-library/internal/middleware"
	"github.com/ramzyrsr/digital-library/internal/repository"
)

type BookHandler struct {
	BookRepo *repository.BookRepository
}

func (h *BookHandler) CreateBook(c *fiber.Ctx) error {
	var req entity.Book
	if err := c.BodyParser(&req); err != nil {
		return middleware.Response(c, fiber.StatusBadRequest, "Invalid input", nil)
	}

	userID := c.Locals("user_id").(uint)
	req.CreatedBy = userID

	if err := h.BookRepo.CreateBook(&req); err != nil {
		log.Fatal(req)
		return middleware.Response(c, fiber.StatusBadRequest, "Failed to create book", nil)
	}

	return middleware.Response(c, fiber.StatusCreated, "Success to create book", nil)
}
