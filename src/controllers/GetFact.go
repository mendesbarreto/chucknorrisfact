package controllers

import (
	"chucknorrisfact/src/models"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func GetFact(context *fiber.Ctx) error {
	response, err := http.Get("https://api.chucknorris.io/jokes/random")

	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err,
		})
	}

	if response.StatusCode != http.StatusOK {
		return context.Status(response.StatusCode).JSON(fiber.Map{
			"message": response,
		})
	}

	var joke models.Joke

	err = json.NewDecoder(response.Body).Decode(&joke)

	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err,
		})
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"fact": models.Fact{
			Id:       joke.Id,
			ImageUrl: joke.IconUrl,
			Value:    joke.Value,
			ValueUrl: joke.Url,
			Provider: "https://api.chucknorris.io",
		},
	})
}
