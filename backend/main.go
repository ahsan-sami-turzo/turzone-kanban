package main

import (
    "log"
    "turzone-kanban/database"
    "turzone-kanban/handlers"
    
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
)

func main() {
    // Initialize database
    if err := database.InitDB(); err != nil {
        log.Fatal("Failed to initialize database:", err)
    }
    
    // Create Gin router
    router := gin.Default()
    
    // Enable CORS for frontend development
    router.Use(cors.Default())
    
    // API routes
    api := router.Group("/api")
    {
        // Project routes
        api.GET("/projects", handlers.GetProjects)
        api.GET("/projects/:id", handlers.GetProject)
        api.POST("/projects/upload", handlers.UploadMarkdown)
        api.DELETE("/projects/:id", handlers.DeleteProject)
        
        // Task routes
        api.GET("/projects/:id/tasks", handlers.GetProjectTasks)
        api.PUT("/tasks/:id", handlers.UpdateTask)
        api.DELETE("/tasks/:id", handlers.DeleteTask)
    }
    
    // Health check
    router.GET("/health", func(c *gin.Context) {
        c.JSON(200, gin.H{"status": "ok"})
    })
    
    log.Println("Server starting on :8080")
    if err := router.Run(":8080"); err != nil {
        log.Fatal("Failed to start server:", err)
    }
}