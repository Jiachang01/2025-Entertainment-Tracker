package books

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	bookRoutes := router.Group("/books")
	bookRoutes.POST("/", AddBook)
	bookRoutes.DELETE("/:id", DeleteBook)
}

func AddBook(c *gin.Context) {
	var book Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Save to database (use repository function)
	c.JSON(http.StatusCreated, gin.H{"message": "Book added successfully"})
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")

	// Delete from database (use repository function)
	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}
