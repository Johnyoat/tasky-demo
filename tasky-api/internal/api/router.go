package api

import (
	"github.com/gofiber/fiber/v3"
	"github.com/johnyoat/tasky-demo/tasky-api/internal/api/v1/handlers"
	"gorm.io/gorm"
)

func InitRouter(app *fiber.App, db *gorm.DB) {

	app.Get("/", func(c fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "ok",
			"message": "Tasky API is running",
		})
	},
	)

	v1 := app.Group("/api/v1")

	v1.Get("/tasks", handlers.GetTasks(db))
	v1.Post("/tasks", handlers.CreateTask(db))
	v1.Put("/tasks/:id", handlers.UpdateTask(db))
	v1.Delete("/tasks/:id", handlers.DeleteTask(db))
}
