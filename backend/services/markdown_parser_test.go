package services

import (
	"testing"
)

func TestParseMarkdown(t *testing.T) {
	markdown := `# My Project

## Task 1
This is the first task
- Detail 1
- Detail 2

## Task 2
Second task description

## Task 3
Third task
`

	result, err := ParseMarkdown(markdown)
	if err != nil {
		t.Fatalf("ParseMarkdown failed: %v", err)
	}

	if result.ProjectName != "My Project" {
		t.Errorf("Expected project name 'My Project', got '%s'", result.ProjectName)
	}

	if len(result.Tasks) != 3 {
		t.Errorf("Expected 3 tasks, got %d", len(result.Tasks))
	}

	if result.Tasks[0].Title != "Task 1" {
		t.Errorf("Expected first task title 'Task 1', got '%s'", result.Tasks[0].Title)
	}
}
