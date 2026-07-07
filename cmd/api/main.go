package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"myapp/internal/config"
	httpHandler "myapp/internal/delivery/http"
	"myapp/internal/repository"
	"myapp/internal/usecase"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Initialize Database Connection
	db := config.ConnectDB()

	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New())

	// Dependency Injection
	menuRepo := repository.NewMenuRepository(db)
	orderRepo := repository.NewOrderRepository(db)
	tableRepo := repository.NewTableRepository(db)

	menuUC := usecase.NewMenuUsecase(menuRepo)
	orderUC := usecase.NewOrderUsecase(orderRepo, menuRepo)
	tableUC := usecase.NewTableUsecase(tableRepo)

	// Setup Routes
	httpHandler.NewHttpHandler(app, menuUC, orderUC, tableUC)

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	log.Fatal(app.Listen(":" + port))
}
