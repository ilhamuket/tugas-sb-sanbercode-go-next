package config

import (
	"fmt"
	"log"
	"os"
	"tugas-sb-sanbercode-go-next-2024/Tugas-React/backend-tugas-reactjs/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var (
	DB *gorm.DB // Memindahkan deklarasi DB ke level package
)

func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading .env file")
	}
}

func InitDB() *gorm.DB { // Mengembalikan *gorm.DB
	// Load environment variables
	LoadEnv()

	dbHost := Getenv("DB_HOST", "localhost")
	dbPort := Getenv("DB_PORT", "5432")
	dbUser := Getenv("DB_USER", "postgres")
	dbName := Getenv("DB_NAME", "db_books")
	dbPassword := Getenv("DB_PASSWORD", "postgres")
	environment := Getenv("ENVIRONMENT", "local")

	var dbURI string
	if environment == "local" {
		dbURI = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
			dbHost, dbPort, dbUser, dbName, dbPassword)
	} else {
		dbURI = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=require",
			dbHost, dbPort, dbUser, dbName, dbPassword)
	}

	log.Printf("Connecting to database at %s:%s\n", dbHost, dbPort)

	var err error
	DB, err = gorm.Open("postgres", dbURI)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	// Auto Migrate the database
	DB.AutoMigrate(&models.Book{})

	return DB // Mengembalikan DB yang sudah diconnect
}

// Getenv Function to retrieve environment variable with default value
func Getenv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
