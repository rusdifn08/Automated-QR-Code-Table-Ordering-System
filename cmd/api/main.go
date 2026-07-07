package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
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

	// Metrics dashboard
	app.Get("/metrics", monitor.New(monitor.Config{Title: "MyMetrics Dashboard"}))

	// Dependency Injection
	menuRepo := repository.NewMenuRepository(db)
	orderRepo := repository.NewOrderRepository(db)
	tableRepo := repository.NewTableRepository(db)
	adminRepo := repository.NewAdminRepository(db)

	menuUsecase := usecase.NewMenuUsecase(menuRepo)
	orderUsecase := usecase.NewOrderUsecase(orderRepo, menuRepo)
	tableUsecase := usecase.NewTableUsecase(tableRepo)
	adminUsecase := usecase.NewAdminUsecase(adminRepo)

	handler := httpHandler.NewHttpHandler(menuUsecase, orderUsecase, tableUsecase, adminUsecase)
	httpHandler.RegisterRoutes(app, handler)

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
