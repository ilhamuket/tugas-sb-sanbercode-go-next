package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"tugas-sb-sanbercode-go-next-2024/Tugas-React/backend-tugas-reactjs/models"
)

var DB *gorm.DB

func InitDB() *gorm.DB {

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
		dbURI = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=require",
			dbHost, dbPort, dbUser, dbName, dbPassword)
	}

	log.Printf("Connecting to database at %s:%s\n", dbHost, dbPort)

	var err error
	DB, err = gorm.Open(postgres.Open(dbURI), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	// Auto Migrate the database tables
	err = MigrateBooksTable()
	if err != nil {
		log.Fatalf("Error migrating database schema: %v", err)
	}

	log.Println("Database migration completed")

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
