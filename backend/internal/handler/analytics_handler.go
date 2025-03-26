package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/ramzyrsr/digital-library/internal/middleware"
	"github.com/ramzyrsr/digital-library/internal/repository"
)

type AnalyticsHandler struct {
	AnalyticsRepo *repository.AnalyticsRepository
}

func (h *AnalyticsHandler) MostBorrowedBooks(c *fiber.Ctx) error {
	limit, _ := strconv.Atoi(c.Query("limit", "5"))

	books, err := h.AnalyticsRepo.GetMostBorrowedBooks(limit)
	if err != nil {
		return middleware.Response(c, fiber.StatusBadRequest, "Failed to retrieve data", err.Error())
	}

	return middleware.Response(c, fiber.StatusOK, "Success to retrieve data", books)
}

func (h *AnalyticsHandler) MonthlyBorrowingTrends(c *fiber.Ctx) error {
	trends, err := h.AnalyticsRepo.GetMonthlyBorrowingTrends()
	if err != nil {
		return middleware.Response(c, fiber.StatusBadRequest, "Failed to retrieve data", err.Error())
	}

	return middleware.Response(c, fiber.StatusOK, "Success to retrieve data", trends)
}

