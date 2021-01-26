package routes

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", func(context *fiber.Ctx) error {
		return context.Status(fiber.StatusOK).JSON(fiber.Map{
			"value": "Bitwise rules them all 😎",
		})
	})

	api := app.Group("/api")

	api.Get("/", func(context *fiber.Ctx) error {
		return context.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "You are on api root",
		})
	})

	FactRoute(api.Group("/fact"))
}