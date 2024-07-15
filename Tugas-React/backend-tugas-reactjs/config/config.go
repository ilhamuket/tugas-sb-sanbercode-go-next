package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"tugas-sb-sanbercode-go-next-2024/Tugas-React/backend-tugas-reactjs/models"
)

var DB *gorm.DB

func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading .env file")
	}
}

func InitDB() *gorm.DB {
	// Load environment variables
	LoadEnv()

	// Set default values for environment variables
	dbHost := Getenv("DB_HOST", "localhost")
	dbPort := Getenv("DB_PORT", "5432")
	dbUser := Getenv("DB_USER", "postgres")
	dbName := Getenv("DB_NAME", "db_books")
	dbPassword := Getenv("DB_PASSWORD", "postgres")
	environment := Getenv("ENVIRONMENT", "local")

	// Build database URI based on environment
	var dbURI string
	if environment == "development" {
		dbURI = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
			dbHost, dbPort, dbUser, dbName, dbPassword)
	} else {
		dbURI = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=require",
			dbHost, dbUser, dbPassword, dbName, dbPort)
	}

	log.Printf("Connecting to database at %s:%s\n", dbHost, dbPort)

	// Open a new database connection if DB is not already set
	if DB == nil {
		var err error
		DB, err = gorm.Open(postgres.Open(dbURI), &gorm.Config{})
		if err != nil {
			log.Fatalf("Error connecting to database: %v", err)
		}

		log.Println("Database connection established")
	}

	return DB
}

func MigrateBooksTable() error {
	if err := DB.AutoMigrate(&models.Book{}); err != nil {
		// Check if the error is due to table already existing
		if !DB.Migrator().HasTable(&models.Book{}) {
			return fmt.Errorf("failed to migrate books table: %v", err)
		}
		log.Println("Books table already exists")
	}
	return nil
}

// Getenv Function to retrieve environment variable with default value
func Getenv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
