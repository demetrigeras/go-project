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
