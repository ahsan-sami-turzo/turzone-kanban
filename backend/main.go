package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
	"turzone-kanban/database"
	"turzone-kanban/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

//go:embed frontend/dist
var frontendFS embed.FS

func main() {
	// Initialize database
	if err := database.InitDB(); err != nil {
		log.Fatal("Failed to initialize database:", err)
	}

	// Set to release mode for production
	gin.SetMode(gin.ReleaseMode)

	// Create Gin router
	router := gin.Default()

	// Enable CORS for API routes
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

	// Serve embedded frontend
	distFS, err := fs.Sub(frontendFS, "frontend/dist")
	if err != nil {
		log.Fatal("Failed to load frontend:", err)
	}

	// Serve static files
	router.StaticFS("/assets", http.FS(distFS))

	// Serve index.html for all other routes (SPA)
	router.NoRoute(func(c *gin.Context) {
		data, err := frontendFS.ReadFile("frontend/dist/index.html")
		if err != nil {
			c.String(http.StatusInternalServerError, "Failed to load page")
			return
		}
		c.Data(http.StatusOK, "text/html; charset=utf-8", data)
	})

	log.Println("🚀 Turzone Kanban started on http://localhost:8080")
	log.Println("📊 Database: ./data/app.db")

	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
