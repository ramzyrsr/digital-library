package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/ramzyrsr/digital-library/config"
	"github.com/ramzyrsr/digital-library/internal/handler"
	"github.com/ramzyrsr/digital-library/internal/middleware"
	"github.com/ramzyrsr/digital-library/internal/repository"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Printf("error loading .env file: %v", err)
		return
	}

	db, err := config.ConnectDB()
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}
	defer db.Close()

	app := fiber.New()

	// Define routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Welcome to the Digital Library API!"})
	})

	// Initialize Repositories & Handlers
	userRepo := &repository.UserRepository{DB: db}
	authHandler := &handler.AuthHandler{UserRepo: userRepo}
	bookRepo := &repository.BookRepository{DB: db}
	bookHandler := &handler.BookHandler{BookRepo: bookRepo}
	lendingRepo := &repository.LendingRepository{DB: db}
	lendingHandler := &handler.LendingHandler{LendingRepo: lendingRepo}
	analyticsRepo := &repository.AnalyticsRepository{DB: db}
	analyticsHandler := &handler.AnalyticsHandler{AnalyticsRepo: analyticsRepo}

	// Auth Routes
	app.Post("/register", authHandler.Register)
	app.Post("/login", authHandler.Login)

	app.Use(middleware.JWTMiddleware())
	app.Post("/book", middleware.StaffOnlyMiddleware(), bookHandler.CreateBook)
	app.Get("/books", bookHandler.GetBooks)
	app.Get("/books/search", bookHandler.GetBooksByTitle)
	app.Delete("/book/:id", middleware.StaffOnlyMiddleware(), bookHandler.DeleteBook)
	app.Post("/lending/book", middleware.StaffOnlyMiddleware(), lendingHandler.BorrowBook)
	app.Put("/lending/return/:id", middleware.StaffOnlyMiddleware(), lendingHandler.ReturnBook)
	app.Get("/analytics/most-borrowed", analyticsHandler.MostBorrowedBooks)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Fatal(app.Listen(":" + port))
}
