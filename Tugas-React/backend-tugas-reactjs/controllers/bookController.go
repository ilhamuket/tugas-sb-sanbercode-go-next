package controllers

import (
	"errors"
	"net/http"
	"regexp"
	"time"
	"tugas-sb-sanbercode-go-next-2024/Tugas-React/backend-tugas-reactjs/config"
	"tugas-sb-sanbercode-go-next-2024/Tugas-React/backend-tugas-reactjs/models"

	"github.com/gin-gonic/gin"
)

// HTTPError defines a simple error structure for JSON responses.
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
		c.JSON(http.StatusBadRequest, HTTPError{Error: "Data tidak valid"})
		return
	}

	// Validate image URL
	if err := validateImageURL(input.ImageURL); err != nil {
		c.JSON(http.StatusBadRequest, HTTPError{Error: "Format URL gambar tidak valid"})
		return
	}

	// Validate release year
	if err := validateReleaseYear(input.ReleaseYear); err != nil {
		c.JSON(http.StatusBadRequest, HTTPError{Error: "Tahun rilis harus berada di antara 1980 dan 2021"})
		return
	}

	// Determine thickness based on total page
	thickness := determineThickness(input.TotalPage)

	book := models.Book{
		Title:       input.Title,
		Description: input.Description,
		ImageURL:    input.ImageURL,
		ReleaseYear: input.ReleaseYear,
		Price:       input.Price,
		TotalPage:   input.TotalPage,
		Thickness:   thickness,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	db := config.GetDB()
	if err := db.Create(&book).Error; err != nil {
		c.JSON(http.StatusInternalServerError, HTTPError{Error: "Gagal menyimpan buku"})
		return
	}

	c.JSON(http.StatusOK, book)
}

func validateImageURL(url string) error {
	if !isValidURL(url) {
		return errors.New("format URL gambar tidak valid")
	}
	return nil
}

func isValidURL(url string) bool {
	return regexp.MustCompile(`^https?://`).MatchString(url)
}

func validateReleaseYear(releaseYear int) error {
	if releaseYear < 1980 || releaseYear > 2021 {
		return errors.New("tahun rilis harus berada di antara 1980 dan 2021")
	}
	return nil
}

func determineThickness(totalPage int) string {
	switch {
	case totalPage <= 100:
		return "tipis"
	case totalPage <= 200:
		return "sedang"
	default:
		return "tebal"
	}
}

// GetBooks godoc
// @Summary Get All Books
// @Description Get all books
// @Tags books
// @Produce  json
// @Success 200 {array} models.Book
// @Router /books [get]
func GetBooks(c *gin.Context) {
	db := config.GetDB()

	var books []models.Book

	if err := db.Find(&books).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch books"})
		return
	}

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
		c.JSON(http.StatusNotFound, HTTPError{Error: "Buku tidak ditemukan"})
		return
	}

	var input models.BookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, HTTPError{Error: "Data tidak valid"})
		return
	}

	// Validate image URL
	if err := validateImageURL(input.ImageURL); err != nil {
		c.JSON(http.StatusBadRequest, HTTPError{Error: "Format URL gambar tidak valid"})
		return
	}

	// Validate release year
	if err := validateReleaseYear(input.ReleaseYear); err != nil {
		c.JSON(http.StatusBadRequest, HTTPError{Error: "Tahun rilis harus berada di antara 1980 dan 2021"})
		return
	}

	// Determine thickness based on total page
	thickness := determineThickness(input.TotalPage)

	// Update book details
	db.Model(&book).Updates(models.Book{
		Title:       input.Title,
		Description: input.Description,
		ImageURL:    input.ImageURL,
		ReleaseYear: input.ReleaseYear,
		Price:       input.Price,
		TotalPage:   input.TotalPage,
		Thickness:   thickness,
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
