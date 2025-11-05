// backend/handlers/task_handler.go (UPDATED)
package handlers

import (
	"net/http"
	"turzone-kanban/database"
	"turzone-kanban/models"

	"github.com/gin-gonic/gin"
)

// UpdateTaskRequest represents the request body for updating a task
type UpdateTaskRequest struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Status      *string `json:"status"`
	Position    *int    `json:"position"`
}

// UpdateTask updates a task's fields
func UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var task models.Task

	// Find task
	if err := database.DB.First(&task, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	// Parse request body
	var req UpdateTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Update fields if provided
	if req.Title != nil {
		task.Title = *req.Title
	}
	if req.Description != nil {
		task.Description = *req.Description
	}
	if req.Status != nil {
		// Validate status
		validStatuses := map[string]bool{
			"waiting":     true,
			"in_progress": true,
			"testing":     true,
			"completed":   true,
		}
		if !validStatuses[*req.Status] {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid status"})
			return
		}
		task.Status = *req.Status
	}
	if req.Position != nil {
		task.Position = *req.Position
	}

	// Save changes
	if err := database.DB.Save(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update task"})
		return
	}

	c.JSON(http.StatusOK, task)
}

// DeleteTask deletes a task
func DeleteTask(c *gin.Context) {
	id := c.Param("id")

	result := database.DB.Delete(&models.Task{}, id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}
