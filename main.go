package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	// Define a route for the root path
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello, man!"})
	})

	// Start the server on port 8080
	fmt.Println("Server listening on port 8080")
	router.Run(":8080")
}
