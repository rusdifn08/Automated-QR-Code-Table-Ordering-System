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
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "host=localhost user=postgres password=postgres dbname=order_system port=5432 sslmode=disable"
	}

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // IMPORTANT FOR SUPABASE POOLER!
	}), &gorm.Config{})
	
	if err != nil {
		log.Fatal("Failed to connect to database:" + err.Error())
	}

	log.Println("Database connection successfully established")

	// Enable UUID extension
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

	// Seed Tables
	var tableCount int64
	db.Model(&domain.Table{}).Count(&tableCount)
	if tableCount == 0 {
		for i := 1; i <= 5; i++ {
			db.Create(&domain.Table{TableNumber: i, Status: "available"})
		}
	}

	// Seed Menu
	var menuCount int64
	db.Model(&domain.Menu{}).Count(&menuCount)
	if menuCount == 0 {
		menus := []domain.Menu{
			{Name: "Nasi Goreng Spesial", Description: "Authentic Indonesian fried rice with sunny side up egg and chicken satay", Price: 45000, Category: "Main", ImageURL: "/assets/food/nasi_goreng.png"},
			{Name: "Sate Ayam Madura", Description: "Chicken skewers in rich peanut sauce served with shallots", Price: 35000, Category: "Main", ImageURL: "/assets/food/sate_ayam.png"},
			{Name: "Wagyu Beef Burger", Description: "Premium wagyu beef patty with melted cheese and fresh brioche bun", Price: 85000, Category: "Main", ImageURL: "/assets/food/wagyu_burger.png"},
			{Name: "Salmon Sushi Roll", Description: "Fresh salmon rolled with avocado and cucumber, topped with spicy mayo", Price: 65000, Category: "Main", ImageURL: "/assets/food/salmon_sushi.png"},
			{Name: "Iced Caramel Macchiato", Description: "Espresso with cold milk and sweet caramel drizzle", Price: 30000, Category: "Beverage", ImageURL: "/assets/food/caramel_macchiato.png"},
		}
		for _, m := range menus {
			db.Create(&m)
		}
	}

	// Seed Admin
	var adminCount int64
	db.Model(&domain.Admin{}).Count(&adminCount)
	if adminCount == 0 {
		hash, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
		db.Create(&domain.Admin{
			Username:     "admin",
			PasswordHash: string(hash),
		})
	}

	return db
}
