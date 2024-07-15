package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"tugas-sb-sanbercode-go-next-2024/Tugas-React/backend-tugas-reactjs/controllers"
	"tugas-sb-sanbercode-go-next-2024/Tugas-React/backend-tugas-reactjs/repositories"
	"tugas-sb-sanbercode-go-next-2024/Tugas-React/backend-tugas-reactjs/services"

	"github.com/gin-contrib/cors"
)

func SetupRouter(db *gorm.DB, app *gin.Engine) *gin.Engine {
	// Use CORS middleware
	app.Use(cors.Default())

	// Initialize controllers
	bookService := services.NewBookService(repositories.NewBookRepository(db))
	bookController := controllers.NewBookController(bookService)

	// Middleware untuk Swagger UI
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Routes untuk Books
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
