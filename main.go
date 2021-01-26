package main

import (
	"chucknorrisfact/src/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)


func main() {
	app := fiber.New()
	app.Use(logger.New())

	routes.SetupRoutes(app)

	err := app.Listen(":8000")

	if err != nil {
		panic(err)
	}
}
