package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
	"tugas-sb-sanbercode-go-next-2024/Tugas-React/backend-tugas-reactjs/controllers"
	"tugas-sb-sanbercode-go-next-2024/Tugas-React/backend-tugas-reactjs/repositories"
	"tugas-sb-sanbercode-go-next-2024/Tugas-React/backend-tugas-reactjs/services"

	"github.com/gin-contrib/cors"
)

func SetupRouter(db *gorm.DB, app *gin.Engine) *gin.Engine {
	// Configure CORS
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowHeaders = []string{"Content-Type", "Accept", "Origin"}

	// Enable CORS with the configured settings
	app.Use(cors.New(corsConfig))

	// Set the db object into the gin context
	app.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	// Initialize controllers and services
	bookService := services.NewBookService(repositories.NewBookRepository(db))
	bookController := controllers.NewBookController(bookService)

	// Middleware for Swagger UI
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Routes for Books
	books := app.Group("/books")
	{
		books.GET("/", bookController.GetBooks)
		books.POST("/", bookController.CreateBook)
		books.GET("/:id", bookController.GetBookByID)
		books.PATCH("/:id", bookController.UpdateBook)
		books.DELETE("/:id", bookController.DeleteBook)
	}

	return app
}
