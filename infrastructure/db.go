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

func InitDB() error {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, using system environment variables")
	}

	dsn := os.Getenv("DATABASE_URL")

	if env := os.Getenv("ENV"); env == "Dev" || env == "development" {
		dsn = fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			os.Getenv("POSTGRES_HOST"),
			os.Getenv("POSTGRES_USER"),
			os.Getenv("POSTGRES_PASSWORD"),
			os.Getenv("POSTGRES_DB"),
			os.Getenv("POSTGRES_PORT"),
		)
	}

	if dsn == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	err = db.AutoMigrate(&domain.Habit{})
	if err != nil {
		log.Fatal("Migration failed:", err)
		return fmt.Errorf("failed to auto-migrate database model: %w", err)
	}

	DB = db
	log.Println("Database intialised & migrated successfully!")
	return nil
}
