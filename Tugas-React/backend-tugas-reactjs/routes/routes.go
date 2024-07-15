package routes

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres" // Contoh untuk paket gorm
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"tugas-sb-sanbercode-go-next-2024/Tugas-React/backend-tugas-reactjs/config"
	"tugas-sb-sanbercode-go-next-2024/Tugas-React/backend-tugas-reactjs/controllers"
	"tugas-sb-sanbercode-go-next-2024/Tugas-React/backend-tugas-reactjs/repositories"
	"tugas-sb-sanbercode-go-next-2024/Tugas-React/backend-tugas-reactjs/services"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Initialize controllers
	bookService := services.NewBookService(repositories.NewBookRepository(config.DB))
	bookController := controllers.NewBookController(bookService)

	// Middleware untuk Swagger UI
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Routes untuk Books
	books := r.Group("/books")
	{
		books.GET("/", bookController.GetBooks)
		books.POST("/", bookController.CreateBook)
		books.GET("/:id", bookController.GetBookByID)
		books.PATCH("/:id", bookController.UpdateBook)
		books.DELETE("/:id", bookController.DeleteBook)
	}

	return r
}
