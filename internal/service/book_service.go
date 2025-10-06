package service

import (
	"errors"
	"example/go_api_tutorial/internal/models"
	"example/go_api_tutorial/internal/repository/interfaces"
	"gorm.io/gorm"
)

// BookService handles business logic for books
type BookService struct {
	bookRepo interfaces.BookRepository
}

// NewBookService creates a new book service
func NewBookService(bookRepo interfaces.BookRepository) *BookService {
	return &BookService{
		bookRepo: bookRepo,
	}
}

// CreateBook creates a new book with validation
func (s *BookService) CreateBook(book *models.Book) error {
	// Business logic validation
	if book.Title == "" {
		return errors.New("book title is required")
	}
	if book.Author == "" {
		return errors.New("book author is required")
	}
	if book.Quantity < 0 {
		return errors.New("book quantity cannot be negative")
	}
	
	return s.bookRepo.Create(book)
}

// GetAllBooks returns all books
func (s *BookService) GetAllBooks() ([]models.Book, error) {
	return s.bookRepo.GetAll()
}

// GetBookByID returns a book by ID
func (s *BookService) GetBookByID(id uint) (*models.Book, error) {
	book, err := s.bookRepo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("book not found")
		}
		return nil, err
	}
	return book, nil
}

// UpdateBook updates a book with validation
func (s *BookService) UpdateBook(id uint, updatedBook *models.Book) (*models.Book, error) {
	// Check if book exists
	existingBook, err := s.bookRepo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("book not found")
		}
		return nil, err
	}
	
	// Business logic validation
	if updatedBook.Title == "" {
		return nil, errors.New("book title is required")
	}
	if updatedBook.Author == "" {
		return nil, errors.New("book author is required")
	}
	if updatedBook.Quantity < 0 {
		return nil, errors.New("book quantity cannot be negative")
	}
	
	// Update fields
	existingBook.Title = updatedBook.Title
	existingBook.Author = updatedBook.Author
	existingBook.Quantity = updatedBook.Quantity
	
	err = s.bookRepo.Update(existingBook)
	if err != nil {
		return nil, err
	}
	
	return existingBook, nil
}

// DeleteBook deletes a book
func (s *BookService) DeleteBook(id uint) error {
	// Check if book exists
	_, err := s.bookRepo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("book not found")
		}
		return err
	}
	
	return s.bookRepo.Delete(id)
}

// SearchBooks searches for books by title or author
func (s *BookService) SearchBooks(query string) ([]models.Book, error) {
	if query == "" {
		return s.bookRepo.GetAll()
	}
	return s.bookRepo.Search(query)
}

// GetBooksPaginated returns paginated books
func (s *BookService) GetBooksPaginated(page, pageSize int) ([]models.Book, int64, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}
	
	offset := (page - 1) * pageSize
	return s.bookRepo.GetPaginated(offset, pageSize)
}

// UpdateBookQuantity updates only the quantity of a book
func (s *BookService) UpdateBookQuantity(id uint, quantity int) error {
	if quantity < 0 {
		return errors.New("quantity cannot be negative")
	}
	
	// Check if book exists
	_, err := s.bookRepo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("book not found")
		}
		return err
	}
	
	return s.bookRepo.UpdateQuantity(id, quantity)
}
