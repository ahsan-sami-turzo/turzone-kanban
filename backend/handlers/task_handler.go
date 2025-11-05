package handlers

import "github.com/gin-gonic/gin"

func UpdateTask(c *gin.Context) {
	c.JSON(200, gin.H{"message": "update task"})
}

func DeleteTask(c *gin.Context) {
	c.JSON(200, gin.H{"message": "delete task"})
}
