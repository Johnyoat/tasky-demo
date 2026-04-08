package handlers

import (
	"github.com/gofiber/fiber/v3"
	database "github.com/johnyoat/tasky-demo/tasky-api/internal/db"
	"gorm.io/gorm"
)

// GetTasks returns a fiber handler that fetches all tasks from the database
func GetTasks(db *gorm.DB) fiber.Handler {
	return func(c fiber.Ctx) error {
		var tasks []database.Task

		err := db.Find(&tasks).Error
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status":  500,
				"message": "Failed to fetch tasks",
				"data":    nil,
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  200,
			"message": "Tasks fetched successfully",
			"data":    tasks,
		})
	}
}

// CreateTask returns a fiber handler that creates a new task in the database
func CreateTask(db *gorm.DB) fiber.Handler {
	return func(c fiber.Ctx) error {
		task := new(database.Task)

		if err := c.Bind().Body(task); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Cannot parse JSON",
			})
		}

		err := db.Create(task).Error
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to create task",
			})
		}

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"status":  201,
			"message": "Task created successfully",
			"data":    task,
		})
	}
}

// UpdateTask returns a fiber handler that updates an existing task by its ID
func UpdateTask(db *gorm.DB) fiber.Handler {
	return func(c fiber.Ctx) error {
		id := c.Params("id")
		task := new(database.Task)

		if err := db.Where("id = ?", id).First(task).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"status":  404,
				"message": "Task not found",
				"data":    nil,
			})
		}

		if err := c.Bind().Body(task); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"status":  400,
				"message": "Cannot parse JSON",
				"data":    nil,
			})
		}

		err := db.Save(task).Error
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status":  500,
				"message": "Failed to update task",
				"data":    nil,
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  200,
			"message": "Task updated successfully",
			"data":    task,
		})
	}
}

// DeleteTask returns a fiber handler that deletes a task by its ID
func DeleteTask(db *gorm.DB) fiber.Handler {
	return func(c fiber.Ctx) error {
		id := c.Params("id")
		task := new(database.Task)

		if err := db.Where("id = ?", id).First(task).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"status":  404,
				"message": "Task not found",
				"data":    nil,
			})
		}

		err := db.Delete(task).Error
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status":  500,
				"message": "Failed to delete task",
				"data":    nil,
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  200,
			"message": "Task deleted successfully",
			"data":    nil,
		})
	}
}
