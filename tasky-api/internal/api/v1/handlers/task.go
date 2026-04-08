package handlers

import (
	"github.com/gofiber/fiber/v3"
	database "github.com/johnyoat/tasky-demo/tasky-api/internal/db"
	"gorm.io/gorm"
)

func GetTasks(db *gorm.DB) fiber.Handler {
	return func(c fiber.Ctx) error {
		tasks := new([]database.Task)

		err := db.Find(&tasks).Error
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to fetch tasks",
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  200,
			"message": "Tasks fetched successfully",
			"data":    tasks,
		})
	}
}

func CreateTask(db *gorm.DB) fiber.Handler {
	return func(c fiber.Ctx) error {
		task := new(database.Task)

		if err := c.Bind().Body(task); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Cannot parse JSON",
			})
		}

		err := db.Create(&task).Error
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

func UpdateTask(db *gorm.DB) fiber.Handler {
	return func(c fiber.Ctx) error {
		id := c.Params("id")
		task := new(database.Task)

		if err := db.First(&task, id).Error; err != nil {
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

		err := db.Save(&task).Error
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

func DeleteTask(db *gorm.DB) fiber.Handler {
	return func(c fiber.Ctx) error {
		id := c.Params("id")
		task := new(database.Task)

		if err := db.First(&task, id).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"status":  404,
				"message": "Task not found",
				"data":    nil,
			})
		}

		err := db.Delete(&task).Error
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
