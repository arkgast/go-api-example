package routes

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"veloud.com/api/internal/data"
)

func GetMovie(c *fiber.Ctx) error {
	id := c.Params("id")
	movieStore := c.Locals("movieStore").(*data.MovieStore)

	movie, err := movieStore.GetMovie(id)
	if err != nil {
		log.Printf("Movie with id %s not found", id)
		return nil
	}

	return c.JSON(movie)
}

func GetMovies(c *fiber.Ctx) error {
	movieStore := c.Locals("movieStore").(*data.MovieStore)
	return c.JSON(movieStore.GetMovies())
}

func CreateMovie(c *fiber.Ctx) error {
	movieStore := c.Locals("movieStore").(*data.MovieStore)
	movie := movieStore.CreateMovie()
	return c.JSON(movie)
}

func UpdateMovie(c *fiber.Ctx) error {
	id := c.Params("id")
	movieStore := c.Locals("movieStore").(*data.MovieStore)

	movie, err := movieStore.UpdateMovie(id)
	if err != nil {
		log.Printf("Movie with id %s not found", id)
		return nil
	}

	return c.JSON(movie)
}

func DeleteMovie(c *fiber.Ctx) error {
	id := c.Params("id")
	movieStore := c.Locals("movieStore").(*data.MovieStore)

	err := movieStore.DeleteMovie(id)
	if err != nil {
		log.Panicf("Movie with id %s not found", id)
		return nil
	}

	return c.SendStatus(fiber.StatusNoContent)
}
