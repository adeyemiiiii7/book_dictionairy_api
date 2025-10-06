package interfaces

import "example/go_api_tutorial/internal/models"

// BookRepository defines the contract for book data operations
type BookRepository interface {
	// Create operations
	Create(book *models.Book) error
	
	// Read operations
	GetAll() ([]models.Book, error)
	GetByID(id uint) (*models.Book, error)
	GetByTitle(title string) ([]models.Book, error)
	GetByAuthor(author string) ([]models.Book, error)
	
	// Update operations
	Update(book *models.Book) error
	UpdateQuantity(id uint, quantity int) error
	
	// Delete operations
	Delete(id uint) error
	
	// Search operations
	Search(query string) ([]models.Book, error)
	
	// Pagination
	GetPaginated(offset, limit int) ([]models.Book, int64, error)
}
