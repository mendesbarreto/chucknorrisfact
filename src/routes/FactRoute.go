package routes

import (
	"chucknorrisfact/src/controllers"
	"github.com/gofiber/fiber/v2"
)

func FactRoute(router fiber.Router) {
	router.Get("", controllers.GetFact)
}