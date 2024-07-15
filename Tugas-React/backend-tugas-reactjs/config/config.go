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
	DB *gorm.DB
)

func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}
}

func InitDB() {
	// Load .env variables
	LoadEnv()

	var err error
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	dbPassword := os.Getenv("DB_PASSWORD")

	var dbURI string
	if os.Getenv("ENVIRONMENT") == "local" {
		dbURI = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
			dbHost, dbPort, dbUser, dbName, dbPassword)
	} else {
		dbURI = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=require",
			dbHost, dbPort, dbUser, dbName, dbPassword)
	}

	log.Println(dbHost, dbPort, dbUser, dbName, dbPassword)

	DB, err = gorm.Open("postgres", dbURI)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	// Auto Migrate the database
	DB.AutoMigrate(&models.Book{})
}
