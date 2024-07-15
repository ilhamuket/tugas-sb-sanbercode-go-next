package api

import (
	"github.com/jinzhu/gorm"
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
	docs.SwaggerInfo.Host = config.Getenv("HOST", "localhost:8080")
	if environment == "development" {
		docs.SwaggerInfo.Schemes = []string{"http", "https"}
	} else {
		docs.SwaggerInfo.Schemes = []string{"https"}
	}

	// Initialize database connection and auto migrate
	db := config.InitDB()

	// Menutup koneksi database menggunakan defer
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {
			// Handle error saat menutup koneksi
			log.Fatalf("Error closing database connection: %v", err)
		}
	}(db)
	// Tutup koneksi database setelah selesai

	// Setup router
	routes.SetupRouter(config.DB, app)

	// Entrypoint
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		app.ServeHTTP(w, r)
	})
}
