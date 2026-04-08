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
