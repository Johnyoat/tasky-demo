package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Task struct {
	gorm.Model
	Title       string `gorm:"not null"`
	Description string
	Completed   bool `gorm:"default:false"`
}

// InitDB initializes the GORM database connection and auto-migrates the Task table
func InitDB(dbPath string) error {
	var err error
	DB, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return err
	}

	return DB.AutoMigrate(&Task{})
}

// CreateTask inserts a new task into the database using GORM
func CreateTask(task *Task) error {
	return DB.Create(task).Error
}

// GetTask retrieves a single task by ID using GORM
func GetTask(id uint) (*Task, error) {
	var task Task
	err := DB.First(&task, id).Error
	if err != nil {
		return nil, err
	}
	return &task, nil
}

// GetAllTasks retrieves all tasks from the database using GORM
func GetAllTasks() ([]Task, error) {
	var tasks []Task
	err := DB.Find(&tasks).Error
	return tasks, err
}

// UpdateTask updates an existing task in the database using GORM
func UpdateTask(task *Task) error {
	return DB.Save(task).Error
}

// DeleteTask deletes a task from the database by ID using GORM
func DeleteTask(id uint) error {
	return DB.Delete(&Task{}, id).Error
}

// CloseDB closes the GORM database connection
func CloseDB() error {
	sqlDB, err := DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
