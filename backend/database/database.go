package database

import (
	"turzone-kanban/models"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() error {
	var err error
	DB, err = gorm.Open(sqlite.Open("../data/app.db"), &gorm.Config{})
	if err != nil {
		return err
	}

	// Auto-migrate the schema
	err = DB.AutoMigrate(&models.Project{}, &models.Task{})
	if err != nil {
		return err
	}

	return nil
}
