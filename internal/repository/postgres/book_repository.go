package postgres

import (
	"example/go_api_tutorial/internal/models"
	"example/go_api_tutorial/internal/repository/interfaces"
	"gorm.io/gorm"
)

// bookRepository implements the BookRepository interface
type bookRepository struct {
	db *gorm.DB
}

// NewBookRepository creates a new book repository
func NewBookRepository(db *gorm.DB) interfaces.BookRepository {
	return &bookRepository{db: db}
}

// Create creates a new book
func (r *bookRepository) Create(book *models.Book) error {
	return r.db.Create(book).Error
}

// GetAll returns all books
func (r *bookRepository) GetAll() ([]models.Book, error) {
	var books []models.Book
	err := r.db.Find(&books).Error
	return books, err
}

// GetByID returns a book by ID
func (r *bookRepository) GetByID(id uint) (*models.Book, error) {
	var book models.Book
	err := r.db.First(&book, id).Error
	if err != nil {
		return nil, err
	}
	return &book, nil
}

// GetByTitle returns books by title (partial match)
func (r *bookRepository) GetByTitle(title string) ([]models.Book, error) {
	var books []models.Book
	err := r.db.Where("title ILIKE ?", "%"+title+"%").Find(&books).Error
	return books, err
}

// GetByAuthor returns books by author (partial match)
func (r *bookRepository) GetByAuthor(author string) ([]models.Book, error) {
	var books []models.Book
	err := r.db.Where("author ILIKE ?", "%"+author+"%").Find(&books).Error
	return books, err
}

// Update updates a book
func (r *bookRepository) Update(book *models.Book) error {
	return r.db.Save(book).Error
}

// UpdateQuantity updates only the quantity of a book
func (r *bookRepository) UpdateQuantity(id uint, quantity int) error {
	return r.db.Model(&models.Book{}).Where("id = ?", id).Update("quantity", quantity).Error
}

// Delete soft deletes a book
func (r *bookRepository) Delete(id uint) error {
	return r.db.Delete(&models.Book{}, id).Error
}

// Search searches books by title or author
func (r *bookRepository) Search(query string) ([]models.Book, error) {
	var books []models.Book
	searchPattern := "%" + query + "%"
	err := r.db.Where("title ILIKE ? OR author ILIKE ?", searchPattern, searchPattern).Find(&books).Error
	return books, err
}

// GetPaginated returns paginated books with total count
func (r *bookRepository) GetPaginated(offset, limit int) ([]models.Book, int64, error) {
	var books []models.Book
	var total int64
	
	// Get total count
	if err := r.db.Model(&models.Book{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	
	// Get paginated results
	err := r.db.Offset(offset).Limit(limit).Find(&books).Error
	return books, total, err
}
