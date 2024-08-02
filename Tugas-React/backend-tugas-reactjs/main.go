package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"tugas-sb-sanbercode-go-next-2024/Tugas-React/backend-tugas-reactjs/config"
	"tugas-sb-sanbercode-go-next-2024/Tugas-React/backend-tugas-reactjs/docs"
	"tugas-sb-sanbercode-go-next-2024/Tugas-React/backend-tugas-reactjs/routes"
)

// Define `app` as a global variable
var app *gin.Engine

// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	// Initialize database
	config.InitDB()

	// Configure Swagger documentation
	docs.SwaggerInfo.Title = "Book REST API"
	docs.SwaggerInfo.Description = "This is a REST API for managing books."
	docs.SwaggerInfo.Version = "1.0"

	environment := os.Getenv("ENVIRONMENT")
	if environment == "development" {
		docs.SwaggerInfo.Host = "localhost:8080"
		docs.SwaggerInfo.Schemes = []string{"http"}
	} else {
		docs.SwaggerInfo.Host = os.Getenv("VERCEL_URL")
		docs.SwaggerInfo.Schemes = []string{"https"}
	}

	// Initialize the Gin engine
	app = gin.Default()

	// Pass the initialized `app` to `SetupRouter`
	db := config.GetDB()
	routes.SetupRouter(db, app)

	// Run the server
	app.Run(":8080")
}

// Ensure the `Handler` function is needed and has access to `app`
func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}
