package api

import (
	"github.com/gofiber/fiber/v3"
)

func InitRouter(app *fiber.App) {

	app.Get("/", func(c fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "ok",
			"message": "Tasky API is running",
		})
	},
	)
}
