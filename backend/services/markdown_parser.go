// backend/services/markdown_parser.go
package services

import (
	"bufio"
	"errors"
	"strings"
	"turzone-kanban/models"
)

// ParsedMarkdown represents the parsed result
type ParsedMarkdown struct {
	ProjectName string
	Tasks       []TaskData
}

// TaskData represents a single task from markdown
type TaskData struct {
	Title       string
	Description string
}

// ParseMarkdown parses a markdown file content
func ParseMarkdown(content string) (*ParsedMarkdown, error) {
	scanner := bufio.NewScanner(strings.NewReader(content))

	var projectName string
	var tasks []TaskData
	var currentTask *TaskData
	var descriptionLines []string

	for scanner.Scan() {
		line := scanner.Text()
		trimmedLine := strings.TrimSpace(line)

		// Skip empty lines at the start
		if trimmedLine == "" && projectName == "" {
			continue
		}

		// Extract project name from # Header
		if strings.HasPrefix(trimmedLine, "# ") {
			projectName = strings.TrimSpace(strings.TrimPrefix(trimmedLine, "# "))
			continue
		}

		// Extract task title from ## Header
		if strings.HasPrefix(trimmedLine, "## ") {
			// Save previous task if exists
			if currentTask != nil {
				currentTask.Description = strings.TrimSpace(strings.Join(descriptionLines, "\n"))
				tasks = append(tasks, *currentTask)
				descriptionLines = nil
			}

			// Start new task
			taskTitle := strings.TrimSpace(strings.TrimPrefix(trimmedLine, "## "))
			currentTask = &TaskData{
				Title: taskTitle,
			}
			continue
		}

		// Collect description lines for current task
		if currentTask != nil && trimmedLine != "" {
			descriptionLines = append(descriptionLines, line)
		}
	}

	// Save the last task
	if currentTask != nil {
		currentTask.Description = strings.TrimSpace(strings.Join(descriptionLines, "\n"))
		tasks = append(tasks, *currentTask)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	// Validate
	if projectName == "" {
		return nil, errors.New("project name not found (missing # Header)")
	}

	if len(tasks) == 0 {
		return nil, errors.New("no tasks found (missing ## Headers)")
	}

	return &ParsedMarkdown{
		ProjectName: projectName,
		Tasks:       tasks,
	}, nil
}

// ConvertToTaskModels converts TaskData to Task models
func ConvertToTaskModels(tasks []TaskData, projectID uint) []models.Task {
	result := make([]models.Task, len(tasks))

	for i, taskData := range tasks {
		result[i] = models.Task{
			ProjectID:   projectID,
			Title:       taskData.Title,
			Description: taskData.Description,
			Status:      "waiting", // All new tasks start in waiting
			Position:    i,         // Sequential position
		}
	}

	return result
}
