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
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowHeaders = []string{"Content-Type", "Accept", "Origin"}

	// To be able to send tokens to the server.
	corsConfig.AllowCredentials = true
	// OPTIONS method for ReactJS
	corsConfig.AddAllowMethods("OPTIONS")

	app.Use(cors.New(corsConfig))

	// set db to gin context
	app.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

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
