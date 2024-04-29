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

	// Start the server on port 5000
	fmt.Println("Server listening on port 5000")
	router.Run("0.0.0.0:5000")
}
