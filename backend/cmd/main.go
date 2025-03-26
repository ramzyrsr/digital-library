package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE",
		AllowHeaders: "Origin,Content-Type,Accept,Authorization",
	}))

	// Define routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Welcome to the Digital Library API!"})
	})

	api := app.Group("/api/v1")

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
	api.Post("/register", authHandler.Register)
	api.Post("/login", authHandler.Login)

	api.Use(middleware.JWTMiddleware())
	api.Post("/member", authHandler.CreateMember)
	api.Post("/book", middleware.StaffOnlyMiddleware(), bookHandler.CreateBook)
	api.Get("/books", bookHandler.GetBooks)
	api.Get("/books/search", bookHandler.GetBooksByTitle)
	api.Delete("/book/:id", middleware.StaffOnlyMiddleware(), bookHandler.DeleteBook)
	api.Post("/lending/book", middleware.StaffOnlyMiddleware(), lendingHandler.BorrowBook)
	api.Put("/lending/return/:id", middleware.StaffOnlyMiddleware(), lendingHandler.ReturnBook)
	api.Get("/analytics/most-borrowed", analyticsHandler.MostBorrowedBooks)
	api.Get("/analytics/borrowing-trends", analyticsHandler.MonthlyBorrowingTrends)
	api.Get("/analytics/books-by-category", analyticsHandler.GetBooksByCategory)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Fatal(app.Listen(":" + port))
}
