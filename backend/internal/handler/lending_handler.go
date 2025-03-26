package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/ramzyrsr/digital-library/internal/entity"
	"github.com/ramzyrsr/digital-library/internal/middleware"
	"github.com/ramzyrsr/digital-library/internal/repository"
)

type LendingHandler struct {
	LendingRepo *repository.LendingRepository
}

func (h *LendingHandler) BorrowBook(c *fiber.Ctx) error {
	var req entity.Lending
	if err := c.BodyParser(&req); err != nil {
		return middleware.Response(c, fiber.StatusBadRequest, "Invalid input", err.Error())
	}

	userID := c.Locals("user_id").(uint)
	req.CreatedBy = userID

	if err := h.LendingRepo.BorrowBook(&req); err != nil {
		return middleware.Response(c, fiber.StatusBadRequest, "Failed to borrow book", err.Error())
	}

	return middleware.Response(c, fiber.StatusCreated, "Success to borrow book", map[string]interface{}{
		"id": req.ID,
	})
}

func (h *LendingHandler) ReturnBook(c *fiber.Ctx) error {
	lendingID, _ := strconv.Atoi(c.Params("id"))

	if err := h.LendingRepo.ReturnBook(lendingID); err != nil {
		return middleware.Response(c, fiber.StatusBadRequest, "Failed to return book", err.Error())
	}

	return middleware.Response(c, fiber.StatusOK, "Success to returned book", nil)
}
