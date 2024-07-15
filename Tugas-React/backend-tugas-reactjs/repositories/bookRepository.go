package repositories

import (
	"gorm.io/gorm"
	"tugas-sb-sanbercode-go-next-2024/Tugas-React/backend-tugas-reactjs/models"
)

type BookRepository interface {
	CreateBook(book *models.Book) error
	GetBooks() ([]models.Book, error)
	GetBookByID(id int) (*models.Book, error)
	UpdateBook(book *models.Book) error
	DeleteBook(id int) error
}

type bookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &bookRepository{db}
}

func (r *bookRepository) CreateBook(book *models.Book) error {
	err := r.db.Create(book).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *bookRepository) GetBooks() ([]models.Book, error) {
	var books []models.Book
	err := r.db.Find(&books).Error
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (r *bookRepository) GetBookByID(id int) (*models.Book, error) {
	var book models.Book
	err := r.db.First(&book, id).Error
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func (r *bookRepository) UpdateBook(book *models.Book) error {
	err := r.db.Save(book).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *bookRepository) DeleteBook(id int) error {
	err := r.db.Delete(&models.Book{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
