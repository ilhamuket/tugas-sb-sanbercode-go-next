// controllers/bookController.go

package controllers

import (
	"net/http"
	"time"
	"tugas-sb-sanbercode-go-next-2024/Tugas-React/backend-tugas-reactjs/config"
	"tugas-sb-sanbercode-go-next-2024/Tugas-React/backend-tugas-reactjs/models"

	"github.com/gin-gonic/gin"
)

// HTTPError represents an HTTP error response
type HTTPError struct {
	Error string `json:"error"`
}

// CreateBook godoc
// @Summary Create Book
// @Description Create a new book
// @Tags books
// @Accept  json
// @Produce  json
// @Param book body models.BookInput true "Book details"
// @Success 200 {object} models.Book
// @Failure 400 {object} HTTPError
// @Router /books [post]
func CreateBook(c *gin.Context) {
	var input models.BookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, HTTPError{Error: err.Error()})
		return
	}

	book := models.Book{
		Title:       input.Title,
		Description: input.Description,
		ImageURL:    input.ImageURL,
		ReleaseYear: input.ReleaseYear,
		Price:       input.Price,
		TotalPage:   input.TotalPage,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	db := config.GetDB()
	db.Create(&book)

	c.JSON(http.StatusOK, book)
}

// GetBooks godoc
// @Summary Get All Books
// @Description Get all books
// @Tags books
// @Produce  json
// @Success 200 {array} models.Book
// @Router /books [get]
func GetBooks(c *gin.Context) {
	// Fetch database instance
	db := config.GetDB()

	// Initialize slice to store books
	var books []models.Book

	// Query all books from the database
	if err := db.Find(&books).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch books"})
		return
	}

	// Return the list of books in JSON format
	c.JSON(http.StatusOK, books)
}

// GetBook godoc
// @Summary Get Book by ID
// @Description Get a book by its ID
// @Tags books
// @Produce  json
// @Param id path int true "Book ID"
// @Success 200 {object} models.Book
// @Failure 404 {object} HTTPError
// @Router /books/{id} [get]
func GetBook(c *gin.Context) {
	var book models.Book
	db := config.GetDB()
	if err := db.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusNotFound, HTTPError{Error: "Book not found"})
		return
	}
	c.JSON(http.StatusOK, book)
}

// UpdateBook godoc
// @Summary Update Book
// @Description Update an existing book
// @Tags books
// @Accept  json
// @Produce  json
// @Param id path int true "Book ID"
// @Param book body models.BookInput true "Book details"
// @Success 200 {object} models.Book
// @Failure 400 {object} HTTPError
// @Failure 404 {object} HTTPError
// @Router /books/{id} [patch]
func UpdateBook(c *gin.Context) {
	var book models.Book
	db := config.GetDB()
	if err := db.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusNotFound, HTTPError{Error: "Book not found"})
		return
	}

	var input models.BookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, HTTPError{Error: err.Error()})
		return
	}

	db.Model(&book).Updates(models.Book{
		Title:       input.Title,
		Description: input.Description,
		ImageURL:    input.ImageURL,
		ReleaseYear: input.ReleaseYear,
		Price:       input.Price,
		TotalPage:   input.TotalPage,
		UpdatedAt:   time.Now(),
	})

	c.JSON(http.StatusOK, book)
}

// DeleteBook godoc
// @Summary Delete Book
// @Description Delete a book by its ID
// @Tags books
// @Produce  json
// @Param id path int true "Book ID"
// @Success 200 {object} map[string]string
// @Failure 404 {object} HTTPError
// @Router /books/{id} [delete]
func DeleteBook(c *gin.Context) {
	var book models.Book
	db := config.GetDB()
	if err := db.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusNotFound, HTTPError{Error: "Book not found"})
		return
	}

	db.Delete(&book)

	c.JSON(http.StatusOK, map[string]string{"data": "Book deleted"})
}
