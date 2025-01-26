package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	books.RegisterRoutes(router)

	log.Println("Server running on port 8080")
	http.ListenAndServe(":8080", router)
}
