package handler

import (
	"log"
	"strconv"

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

func (h *BookHandler) GetBooks(c *fiber.Ctx) error {
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	offset, _ := strconv.Atoi(c.Query("offset", "0"))

	books, totalBookFiltered, totalBook, err := h.BookRepo.GetBooks(limit, offset)
	if err != nil {
		return middleware.Response(c, fiber.StatusBadRequest, "Failed to retrieve books", nil)
	}

	return middleware.Response(c, fiber.StatusOK, "Success to retrieve book", map[string]interface{}{
		"data":                books,
		"total_data":          totalBookFiltered,
		"total_data_filtered": totalBook,
	})
}

func (h *BookHandler) GetBooksByTitle(c *fiber.Ctx) error {
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	offset, _ := strconv.Atoi(c.Query("offset", "0"))
	filter := c.Query("title", "")

	books, totalBookFiltered, totalBook, err := h.BookRepo.GetBooksByTitle(limit, offset, filter)
	if err != nil {
		return middleware.Response(c, fiber.StatusBadRequest, "Failed to retrieve books", nil)
	}

	return middleware.Response(c, fiber.StatusOK, "Success to retrieve book", map[string]interface{}{
		"data":                books,
		"total_data":          totalBookFiltered,
		"total_data_filtered": totalBook,
	})
}

