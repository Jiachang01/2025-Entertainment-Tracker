package main

import (
	"log"

	"github.com/Jiachang01/2025-Entertainment-Tracker/internal/books"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	books.RegisterRoutes(router)

	log.Println("Server running on port 8080")
	err := router.Run(":8080")
	if err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
