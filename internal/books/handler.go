package books

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handler holds the book service
type Handler struct {
	service *Service
}

// NewHandler creates a new book handler
func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

// RegisterRoutes sets up book routes
func RegisterRoutes(router *gin.Engine) {
	repo := NewRepository()
	service := NewService(repo)
	handler := NewHandler(service)

	router.GET("/books", handler.getBooks)
	router.POST("/books", handler.addBook)
}

// GetBooks returns all books
func (h *Handler) getBooks(c *gin.Context) {
	books := h.service.GetBooks()
	c.JSON(http.StatusOK, books)
}

// AddBook adds a new book
func (h *Handler) addBook(c *gin.Context) {
	var newBook Book
	if err := c.ShouldBindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	createdBook := h.service.AddBook(newBook.Title, newBook.Author)
	c.JSON(http.StatusCreated, createdBook)
}
