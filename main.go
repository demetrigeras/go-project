package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// Create a new Gin router
	r := gin.Default()

	// Define a handler for the GET request on the root URL ("/")
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	// Run the server on port 8080
	r.Run(":8080")
}
