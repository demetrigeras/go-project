package main

import (
	"go-project/controllers"

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
	r.POST("/tasks", controllers.CreateTask)       // Create a new task
	r.GET("/tasks", controllers.GetTasks)          // Retrieve all tasks
	r.PUT("/tasks/:id", controllers.UpdateTask)    // Update a specific task
	r.DELETE("/tasks/:id", controllers.DeleteTask) // Delete a specific task

	// Run the server on port 8080
	r.Run(":8080")
}
