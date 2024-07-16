package api

import (
	"github.com/gin-contrib/cors"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"os"
	"tugas-sb-sanbercode-go-next-2024/Tugas-React/backend-tugas-reactjs/config"
	"tugas-sb-sanbercode-go-next-2024/Tugas-React/backend-tugas-reactjs/controllers"
	"tugas-sb-sanbercode-go-next-2024/Tugas-React/backend-tugas-reactjs/docs"

	"github.com/gin-gonic/gin"
)

var app *gin.Engine

func init() {
	// Initialize Gin engine
	app = gin.Default()

	// Setup CORS
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowHeaders = []string{"Content-Type", "Accept", "Origin"}

	// Enable CORS with the configured settings
	app.Use(cors.New(corsConfig))

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

	// Setup routes
	setupRouter()
}

func setupRouter() {
	// Middleware to set db in context
	app.Use(func(c *gin.Context) {
		c.Set("db", config.GetDB())
		c.Next()
	})

	// Routes
	app.POST("/books", controllers.CreateBook)
	app.GET("/books", controllers.GetBooks)
	app.GET("/books/:id", controllers.GetBook)
	app.PATCH("/books/:id", controllers.UpdateBook)
	app.DELETE("/books/:id", controllers.DeleteBook)

	// Swagger endpoint
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

// Handler function to handle HTTP requests
func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}
