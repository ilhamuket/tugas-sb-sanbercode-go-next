package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"tugas-sb-sanbercode-go-next-2024/Tugas-React/backend-tugas-reactjs/config"
	"tugas-sb-sanbercode-go-next-2024/Tugas-React/backend-tugas-reactjs/docs"
	"tugas-sb-sanbercode-go-next-2024/Tugas-React/backend-tugas-reactjs/routes"
)

var app *gin.Engine

// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func init() {

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
}

// Handler function to handle HTTP requests
func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}
