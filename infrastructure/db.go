package infrastructure

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/jt00721/habit-tracker/internal/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file:", err)
	}

	dbName := os.Getenv("POSTGRES_DB")
	if os.Getenv("TEST_MODE") == "true" {
		dbName = os.Getenv("POSTGRES_TEST_DB")
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		dbName,
		os.Getenv("POSTGRES_PORT"),
	)

	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	err = DB.AutoMigrate(&domain.Habit{})
	if err != nil {
		log.Fatal("Migration failed:", err)
	}

	log.Println("Database intialised & migrated successfully!")
}
