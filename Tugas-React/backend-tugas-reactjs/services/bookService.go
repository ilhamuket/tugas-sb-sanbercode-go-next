package services

import "tugas-sb-sanbercode-go-next-2024/Tugas-React/backend-tugas-reactjs/models"
import "tugas-sb-sanbercode-go-next-2024/Tugas-React/backend-tugas-reactjs/repositories"

type BookService interface {
	CreateBook(book *models.Book) error
	GetBooks() ([]models.Book, error)
	GetBookByID(id int) (*models.Book, error)
	UpdateBook(id int, book *models.Book) error
	DeleteBook(id int) error
}

type bookService struct {
	repo repositories.BookRepository
}

func NewBookService(repo repositories.BookRepository) BookService {
	return &bookService{repo}
}

func (s *bookService) CreateBook(book *models.Book) error {
	return s.repo.CreateBook(book)
}

func (s *bookService) GetBooks() ([]models.Book, error) {
	return s.repo.GetBooks()
}

func (s *bookService) GetBookByID(id int) (*models.Book, error) {
	return s.repo.GetBookByID(id)
}

func (s *bookService) UpdateBook(id int, book *models.Book) error {
	book.ID = id
	return s.repo.UpdateBook(book)
}

func (s *bookService) DeleteBook(id int) error {
	return s.repo.DeleteBook(id)
}
