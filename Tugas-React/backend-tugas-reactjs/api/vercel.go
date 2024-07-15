package api

import (
	"log"
	"net/http"
	"tugas-sb-sanbercode-go-next-2024/Tugas-React/backend-tugas-reactjs/config"
	"tugas-sb-sanbercode-go-next-2024/Tugas-React/backend-tugas-reactjs/docs"
	"tugas-sb-sanbercode-go-next-2024/Tugas-React/backend-tugas-reactjs/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	app *gin.Engine
)

func init() {
	app = gin.Default()

	environment := config.Getenv("ENVIRONMENT", "development")

	if environment == "development" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	docs.SwaggerInfo.Title = "Movie REST API"
	docs.SwaggerInfo.Description = "This is REST API Movie."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = config.Getenv("VERCEL_URL", "localhost")
	if environment == "development" {
		docs.SwaggerInfo.Schemes = []string{"http", "https"}
	} else {
		docs.SwaggerInfo.Schemes = []string{"https"}
	}

	// Initialize database connection and auto migrate
	db := config.InitDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	// Setup router
	routes.SetupRouter(config.DB, app)
}

// Handler untuk menangani permintaan HTTP
func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}
