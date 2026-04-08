package api

import (
	"github.com/gofiber/fiber/v3"
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
}
