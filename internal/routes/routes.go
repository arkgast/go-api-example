package routes

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"veloud.com/api/internal/data"
)

func IndexHandler(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func GetMovie(c *fiber.Ctx) error {
	params := c.AllParams()
	id := params["id"]

	movie, err := data.GetMovie(id)
	if err != nil {
		log.Printf("Movie with id %s not found", id)
		return nil
	}

	return c.JSON(movie)
}

func GetMovies(c *fiber.Ctx) error {
	movies := data.GetMovies()
	return c.JSON(movies)
}

func CreateMovie(c *fiber.Ctx) error {
	movie := data.CreateMovie()
	return c.JSON(movie)
}

func UpdateMovie(c *fiber.Ctx) error {
	id := c.Params("id")
	movie, err := data.UpdateMovie(id)
	if err != nil {
		log.Printf("Movie with id %s not found", id)
		return nil
	}

	return c.JSON(movie)
}

func DeleteMovie(c *fiber.Ctx) error {
	id := c.Params("id")
	err := data.DeleteMovie(id)
	if err != nil {
		log.Panicf("Movie with id %s not found", id)
		return nil
	}
	return c.SendStatus(fiber.StatusNoContent)
}
