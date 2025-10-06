package handler

import (
	"net/http"
	"strconv"

	"example/go_api_tutorial/internal/models"
	"example/go_api_tutorial/internal/service"
	"github.com/gin-gonic/gin"
)

// BookHandler handles HTTP requests for books
type BookHandler struct {
	bookService *service.BookService
}

// NewBookHandler creates a new book handler
func NewBookHandler(bookService *service.BookService) *BookHandler {
	return &BookHandler{
		bookService: bookService,
	}
}

// GetBooks handles GET /books
func (h *BookHandler) GetBooks(c *gin.Context) {
	// Check for search query
	search := c.Query("search")
	if search != "" {
		books, err := h.bookService.SearchBooks(search)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"books": books, "count": len(books)})
		return
	}
	
	// Check for pagination
	pageStr := c.Query("page")
	pageSizeStr := c.Query("page_size")
	
	if pageStr != "" || pageSizeStr != "" {
		page, _ := strconv.Atoi(pageStr)
		pageSize, _ := strconv.Atoi(pageSizeStr)
		
		books, total, err := h.bookService.GetBooksPaginated(page, pageSize)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		
		c.JSON(http.StatusOK, gin.H{
			"books":      books,
			"total":      total,
			"page":       page,
			"page_size":  pageSize,
			"total_pages": (total + int64(pageSize) - 1) / int64(pageSize),
		})
		return
	}
	
	// Default: get all books
	books, err := h.bookService.GetAllBooks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, books)
}

// GetBookByID handles GET /books/:id
func (h *BookHandler) GetBookByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}
	
	book, err := h.bookService.GetBookByID(uint(id))
	if err != nil {
		if err.Error() == "book not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, book)
}

// CreateBook handles POST /books
func (h *BookHandler) CreateBook(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	if err := h.bookService.CreateBook(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusCreated, book)
}

// UpdateBook handles PUT /books/:id
func (h *BookHandler) UpdateBook(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}
	
	var updatedBook models.Book
	if err := c.ShouldBindJSON(&updatedBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	book, err := h.bookService.UpdateBook(uint(id), &updatedBook)
	if err != nil {
		if err.Error() == "book not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, book)
}

// DeleteBook handles DELETE /books/:id
func (h *BookHandler) DeleteBook(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}
	
	if err := h.bookService.DeleteBook(uint(id)); err != nil {
		if err.Error() == "book not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}

// UpdateBookQuantity handles PATCH /books/:id/quantity
func (h *BookHandler) UpdateBookQuantity(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}
	
	var req struct {
		Quantity int `json:"quantity" binding:"required,min=0"`
	}
	
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	if err := h.bookService.UpdateBookQuantity(uint(id), req.Quantity); err != nil {
		if err.Error() == "book not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"message": "Book quantity updated successfully"})
}
