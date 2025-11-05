// backend/models/models.go
package models

import (
	"time"
)

// Project represents a project/board
type Project struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"unique;not null" json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Tasks     []Task    `gorm:"foreignKey:ProjectID;constraint:OnDelete:CASCADE" json:"tasks,omitempty"`
}

// Task represents a task within a project
type Task struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	ProjectID   uint      `gorm:"not null;index" json:"project_id"`
	Title       string    `gorm:"not null" json:"title"`
	Description string    `gorm:"type:text" json:"description"`
	Status      string    `gorm:"not null;default:'waiting'" json:"status"` // waiting, in_progress, testing, completed
	Position    int       `gorm:"default:0" json:"position"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// TableName overrides (optional, GORM will pluralize by default)
func (Project) TableName() string {
	return "projects"
}

func (Task) TableName() string {
	return "tasks"
}
