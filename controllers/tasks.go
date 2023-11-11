package controllers

import (
	"go-project/models" // Only if you need to reference models
	"net/http"

	"github.com/gin-gonic/gin"
)

var tasks = []models.Task{}

func CreateTask(c *gin.Context) {
	var newTask models.Task
	if err := c.BindJSON(&newTask); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	tasks = append(tasks, newTask) // Add the new task to the slice
	c.JSON(http.StatusCreated, newTask)
}

// GetTasks handles GET requests to retrieve all tasks
func GetTasks(c *gin.Context) {
	c.JSON(http.StatusOK, tasks)
}

func UpdateTask(c *gin.Context) {
	id := c.Param("id") // Get the ID from the URL parameter
	var updatedTask models.Task

	if err := c.BindJSON(&updatedTask); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	for i, t := range tasks {
		if t.ID == id {
			tasks[i] = updatedTask // Update the task
			c.JSON(http.StatusOK, updatedTask)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
}

func DeleteTask(c *gin.Context) {
	id := c.Param("id") // Get the ID from the URL parameter

	for i, t := range tasks {
		if t.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...) // Remove the task
			c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
}
