package main

import (
	"log"
	"net/http"

	"Quiz-1/api/handlers"
	"Quiz-1/api/middleware"
	"Quiz-1/config"
	"Quiz-1/db"
	"github.com/julienschmidt/httprouter"
)

func main() {
	db.InitDB(config.GetDBConnectionString())

	router := httprouter.New()

	// Bangun Datar dan Bangun Ruang
	router.GET("/bangun-datar/:shape", handlers.CalculateShape)
	router.GET("/bangun-ruang/:shape", handlers.CalculateShape)

	// Categories
	router.GET("/categories", handlers.GetCategories)
	router.POST("/categories", middleware.BasicAuth(handlers.CreateCategory))
	router.PUT("/categories/:id", middleware.BasicAuth(handlers.UpdateCategory))
	router.DELETE("/categories/:id", middleware.BasicAuth(handlers.DeleteCategory))
	router.GET("/categories/:id/articles", handlers.GetArticlesByCategoryID)

	// Articles
	router.GET("/articles", handlers.GetArticles)
	router.POST("/articles", middleware.BasicAuth(handlers.CreateArticle))
	router.PUT("/articles/:id", middleware.BasicAuth(handlers.UpdateArticle))
	router.DELETE("/articles/:id", middleware.BasicAuth(handlers.DeleteArticle))

	log.Println("Server running on port 8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}
