package handlers

import "github.com/gin-gonic/gin"

func GetProjects(c *gin.Context) {
	c.JSON(200, gin.H{"message": "get projects"})
}

func GetProject(c *gin.Context) {
	c.JSON(200, gin.H{"message": "get project"})
}

func UploadMarkdown(c *gin.Context) {
	c.JSON(200, gin.H{"message": "upload markdown"})
}

func DeleteProject(c *gin.Context) {
	c.JSON(200, gin.H{"message": "delete project"})
}

func GetProjectTasks(c *gin.Context) {
	c.JSON(200, gin.H{"message": "get tasks"})
}
