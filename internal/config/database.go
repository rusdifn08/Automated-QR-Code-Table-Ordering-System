package config

import (
	"log"
	"os"

	"myapp/internal/domain"
	"golang.org/x/crypto/bcrypt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	dsn := os.Getenv("DATABASE_URL") // Supabase connection string
	if dsn == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	}

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	log.Println("Database connection successfully established")
	
	// Ensure the uuid-ossp extension is created
	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")

	// AutoMigrate models
	err = db.AutoMigrate(
		&domain.Admin{},
		&domain.Table{},
		&domain.Menu{},
		&domain.Order{},
		&domain.OrderItem{},
	)
	if err != nil {
		log.Fatal("Failed to automigrate models:" + err.Error())
	}

	// Seed Admin
	var count int64
	db.Model(&domain.Admin{}).Count(&count)
	if count == 0 {
		hash, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
		db.Create(&domain.Admin{
			Username:     "admin",
			PasswordHash: string(hash),
		})
	}

	return db
}
