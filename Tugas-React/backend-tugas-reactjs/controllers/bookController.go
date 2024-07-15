package controllers

import (
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
	"tugas-sb-sanbercode-go-next-2024/Tugas-React/backend-tugas-reactjs/models"
	"tugas-sb-sanbercode-go-next-2024/Tugas-React/backend-tugas-reactjs/services"

	"github.com/gin-gonic/gin"
)

type BookController struct {
	service services.BookService
}

func NewBookController(service services.BookService) *BookController {
	return &BookController{service}
}

func validateBook(book *models.Book) []string {
	var errors []string

	// Validasi image_url
	match, _ := regexp.MatchString(`^(http|https):\/\/[^\s]+$`, book.ImageURL)
	if !match {
		errors = append(errors, "Image URL is not valid")
	}

	// Validasi release_year
	if book.ReleaseYear < 1980 || book.ReleaseYear > 2021 {
		errors = append(errors, "Release year must be between 1980 and 2021")
	}

	return errors
}

func setThickness(book *models.Book) {
	if book.TotalPage <= 100 {
		book.Thickness = "tipis"
	} else if book.TotalPage <= 200 {
		book.Thickness = "sedang"
	} else {
		book.Thickness = "tebal"
	}
}

// CreateBook godoc
// @Summary Create a new book
// @Description Create a new book
// @Tags books
// @Accept json
// @Produce json
// @Param book body models.BookInput true "Book"
// @Success 201 {object} models.Book
// @Failure 400 {object} ResponseError
// @Failure 500 {object} ResponseError
// @Router /books [post]
func (ctrl *BookController) CreateBook(c *gin.Context) {
	var input models.BookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, ResponseError{Error: err.Error()})
		return
	}

	// Validasi input
	errors := validateBookInput(&input)
	if len(errors) > 0 {
		c.JSON(http.StatusBadRequest, ResponseError{Error: strings.Join(errors, ", ")})
		return
	}

	// Buat objek Book dari input
	book := models.Book{
		Title:       input.Title,
		Description: input.Description,
		ImageURL:    input.ImageURL,
		ReleaseYear: input.ReleaseYear,
		Price:       input.Price,
		TotalPage:   input.TotalPage,
	}

	// Validasi tambahan seperti validasi custom
	errors = validateBook(&book)
	if len(errors) > 0 {
		c.JSON(http.StatusBadRequest, ResponseError{Error: strings.Join(errors, ", ")})
		return
	}

	// Set thickness
	setThickness(&book)

	// Set CreatedAt and UpdatedAt
	book.CreatedAt = time.Now()
	book.UpdatedAt = time.Now()

	if err := ctrl.service.CreateBook(&book); err != nil {
		c.JSON(http.StatusInternalServerError, ResponseError{Error: "Failed to create book"})
		return
	}

	c.JSON(http.StatusCreated, book)
}

func validateBookInput(input *models.BookInput) []string {
	var errors []string

	// Validasi required fields
	if input.Title == "" {
		errors = append(errors, "Title is required")
	}
	if input.Description == "" {
		errors = append(errors, "Description is required")
	}
	if input.ImageURL == "" {
		errors = append(errors, "Image URL is required")
	} else {
		// Validasi URL format
		match, _ := regexp.MatchString(`^(http|https):\/\/[^\s]+$`, input.ImageURL)
		if !match {
			errors = append(errors, "Image URL is not valid")
		}
	}

	// Validasi ReleaseYear range
	if input.ReleaseYear < 1980 || input.ReleaseYear > 2021 {
		errors = append(errors, "Release year must be between 1980 and 2021")
	}

	// Validasi TotalPage required
	if input.TotalPage <= 0 {
		errors = append(errors, "Total page must be greater than 0")
	}

	// Validasi Price required
	if input.Price == "" {
		errors = append(errors, "Price is required")
	}

	return errors
}

// GetBooks godoc
// @Summary Get all books
// @Description Get all books
// @Tags books
// @Produce json
// @Success 200 {array} models.Book
// @Failure 500 {object} ResponseError
// @Router /books [get]
func (ctrl *BookController) GetBooks(c *gin.Context) {
	books, err := ctrl.service.GetBooks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, ResponseError{Error: "Failed to fetch books"})
		return
	}

	c.JSON(http.StatusOK, books)
}

// GetBookByID godoc
// @Summary Get a book by ID
// @Description Get a book by ID
// @Tags books
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} models.Book
// @Failure 404 {object} ResponseError
// @Router /books/{id} [get]
func (ctrl *BookController) GetBookByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	book, err := ctrl.service.GetBookByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, ResponseError{Error: "Book not found"})
		return
	}

	c.JSON(http.StatusOK, book)
}

// UpdateBook godoc
// @Summary Update a book by ID
// @Description Update a book by ID
// @Tags books
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Param book body models.BookInput true "Book"
// @Success 200 {object} models.Book
// @Failure 400 {object} ResponseError
// @Failure 500 {object} ResponseError
// @Router /books/{id} [patch]
func (ctrl *BookController) UpdateBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	// Fetch existing book
	existingBook, err := ctrl.service.GetBookByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, ResponseError{Error: "Book not found"})
		return
	}

	// Bind input to BookInput
	var input models.BookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, ResponseError{Error: err.Error()})
		return
	}

	// Update existingBook fields
	existingBook.Title = input.Title
	existingBook.Description = input.Description
	existingBook.ImageURL = input.ImageURL
	existingBook.ReleaseYear = input.ReleaseYear
	existingBook.Price = input.Price
	existingBook.TotalPage = input.TotalPage

	// Validasi input
	errors := validateBook(existingBook)
	if len(errors) > 0 {
		c.JSON(http.StatusBadRequest, ResponseError{Error: strings.Join(errors, ", ")})
		return
	}

	// Set thickness
	setThickness(existingBook)

	// Set UpdatedAt
	existingBook.UpdatedAt = time.Now()

	if err := ctrl.service.UpdateBook(id, existingBook); err != nil {
		c.JSON(http.StatusInternalServerError, ResponseError{Error: "Failed to update book"})
		return
	}

	c.JSON(http.StatusOK, existingBook)
}

// DeleteBook godoc
// @Summary Delete a book by ID
// @Description Delete a book by ID
// @Tags books
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} ResponseMessage
// @Failure 500 {object} ResponseError
// @Router /books/{id} [delete]
func (ctrl *BookController) DeleteBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := ctrl.service.DeleteBook(id); err != nil {
		c.JSON(http.StatusInternalServerError, ResponseError{Error: "Failed to delete book"})
		return
	}

	c.JSON(http.StatusOK, ResponseMessage{Message: "Book deleted successfully"})
}
