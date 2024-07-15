package api

import (
	"net/http"
	"tugas-sb-sanbercode-go-next-2024/Tugas-React/backend-tugas-reactjs/config"
	"tugas-sb-sanbercode-go-next-2024/Tugas-React/backend-tugas-reactjs/docs"
	"tugas-sb-sanbercode-go-next-2024/Tugas-React/backend-tugas-reactjs/routes"

	"github.com/gin-gonic/gin"
)

var (
	app *gin.Engine
)

func init() {
	app = gin.New()

	// Load environment variables
	config.LoadEnv()

	// Set up Swagger documentation
	setupSwagger()

	// Initialize database connection and auto migrate
	config.InitDB()

	// Setup router
	routes.SetupRouter()

	// Entrypoint
	http.HandleFunc("/", Handler)
}

// Handler for serving HTTP requests
func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}

// Function to setup Swagger documentation
func setupSwagger() {
	environment := config.Getenv("ENVIRONMENT", "local")

	docs.SwaggerInfo.Title = "Movie REST API"
	docs.SwaggerInfo.Description = "This is REST API Movie."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = config.Getenv("HOST", "localhost:8080")
	if environment == "local" {
		docs.SwaggerInfo.Schemes = []string{"http", "https"}
	} else {
		docs.SwaggerInfo.Schemes = []string{"https"}
	}
}
