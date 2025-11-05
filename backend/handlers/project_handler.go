// backend/handlers/project_handler.go (UPDATED)
package handlers

import (
	"io"
	"net/http"
	"turzone-kanban/database"
	"turzone-kanban/models"
	"turzone-kanban/services"

	"github.com/gin-gonic/gin"
)

// GetProjects returns all projects
func GetProjects(c *gin.Context) {
	var projects []models.Project

	if err := database.DB.Find(&projects).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, projects)
}

// GetProject returns a single project by ID
func GetProject(c *gin.Context) {
	id := c.Param("id")
	var project models.Project

	if err := database.DB.First(&project, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
		return
	}

	c.JSON(http.StatusOK, project)
}

// UploadMarkdown handles markdown file upload
func UploadMarkdown(c *gin.Context) {
	// Get uploaded file
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file uploaded"})
		return
	}

	// Check file extension
	if file.Header.Get("Content-Type") != "text/markdown" && !isMarkdownFile(file.Filename) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Only markdown files (.md) are allowed"})
		return
	}

	// Read file content
	openedFile, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read file"})
		return
	}
	defer openedFile.Close()

	content, err := io.ReadAll(openedFile)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read file content"})
		return
	}

	// Parse markdown
	parsed, err := services.ParseMarkdown(string(content))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid markdown format: " + err.Error()})
		return
	}

	// Check if project exists
	var project models.Project
	result := database.DB.Where("name = ?", parsed.ProjectName).First(&project)

	if result.Error != nil {
		// Project doesn't exist, create new
		project = models.Project{
			Name: parsed.ProjectName,
		}
		if err := database.DB.Create(&project).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create project"})
			return
		}
	}

	// Convert and create tasks
	taskModels := services.ConvertToTaskModels(parsed.Tasks, project.ID)

	// Get current max position for existing tasks
	var maxPosition int
	database.DB.Model(&models.Task{}).
		Where("project_id = ?", project.ID).
		Select("COALESCE(MAX(position), -1)").
		Scan(&maxPosition)

	// Adjust positions for new tasks
	for i := range taskModels {
		taskModels[i].Position = maxPosition + 1 + i
	}

	// Batch insert tasks
	if err := database.DB.Create(&taskModels).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create tasks"})
		return
	}

	// Load project with tasks
	database.DB.Preload("Tasks").First(&project, project.ID)

	c.JSON(http.StatusOK, gin.H{
		"message":       "Project uploaded successfully",
		"project":       project,
		"tasks_created": len(taskModels),
	})
}

// DeleteProject deletes a project and all its tasks
func DeleteProject(c *gin.Context) {
	id := c.Param("id")

	result := database.DB.Delete(&models.Project{}, id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Project deleted successfully"})
}

// GetProjectTasks returns all tasks for a project
func GetProjectTasks(c *gin.Context) {
	projectID := c.Param("id")
	var tasks []models.Task

	if err := database.DB.Where("project_id = ?", projectID).Order("position").Find(&tasks).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

// Helper function to check markdown file extension
func isMarkdownFile(filename string) bool {
	return len(filename) > 3 && filename[len(filename)-3:] == ".md"
}
