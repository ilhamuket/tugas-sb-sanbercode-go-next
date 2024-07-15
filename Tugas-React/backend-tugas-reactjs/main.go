package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"tugas-sb-sanbercode-go-next-2024/Tugas-React/backend-tugas-reactjs/config"
	"tugas-sb-sanbercode-go-next-2024/Tugas-React/backend-tugas-reactjs/docs"
	"tugas-sb-sanbercode-go-next-2024/Tugas-React/backend-tugas-reactjs/routes"
)

var (
	app *gin.Engine
)

func init() {
	app = gin.Default()

	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	docs.SwaggerInfo.Title = "Movie REST API"
	docs.SwaggerInfo.Description = "This is REST API Movie."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = config.Getenv("VERCEL_URL", "localhost:8080")
	// In development, allow both HTTP and HTTPS
	environment := config.Getenv("ENVIRONMENT", "development")
	log.Print(environment)
	if environment == "development" {
		docs.SwaggerInfo.Schemes = []string{"http", "https"}
	} else {
		docs.SwaggerInfo.Schemes = []string{"https"}
	}

	// Initialize database connection and auto migrate
	config.InitDB()
}

func main() {
	// Setup router
	r := gin.Default()
	r = routes.SetupRouter(config.DB, r)

	// Run the server
	port := config.Getenv("PORT", "8080")
	err := r.Run(":" + port)
	if err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
