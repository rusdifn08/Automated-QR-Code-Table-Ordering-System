package config

import (
	"log"
	"os"

	"myapp/internal/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	dsn := os.Getenv("DATABASE_URL") // Supabase connection string
	if dsn == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	log.Println("Database connection successfully established")
	
	// Ensure the uuid-ossp extension is created
	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")

	// AutoMigrate models
	err = db.AutoMigrate(&domain.Table{}, &domain.Menu{}, &domain.Order{}, &domain.OrderItem{})
	if err != nil {
		log.Fatal("Failed to automigrate models:", err)
	}

	return db
}
