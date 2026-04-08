package db

import (
	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Task represents the database model for a task
type Task struct {
	ID          string `gorm:"primaryKey" json:"id"`
	Title       string `gorm:"not null" json:"title"`
	Description string `json:"description"`
	Completed   bool   `gorm:"default:false" json:"completed"`
}

// BeforeCreate generates a new UUID for the task ID before it is created in the database
func (t *Task) BeforeCreate(tx *gorm.DB) (err error) {
	t.ID = uuid.New().String()
	return
}

// InitDB initializes the GORM database connection and auto-migrates the Task table
func InitDB(dbPath string) (*gorm.DB, error) {
	var err error
	DB, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = DB.AutoMigrate(&Task{})
	if err != nil {
		return nil, err
	}

	return DB, nil
}
