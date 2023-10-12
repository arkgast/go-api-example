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

	log.Printf("Movie with ID: %s was retrieved successfully", movie.ID)
	return c.JSON(movie)
}

func GetMovies(c *fiber.Ctx) error {
	movieStore := c.Locals("movieStore").(*data.MovieStore)
	log.Println("Get movies executed successfully")
	return c.JSON(movieStore.GetMovies())
}

func CreateMovie(c *fiber.Ctx) error {
	movieStore := c.Locals("movieStore").(*data.MovieStore)
	movie := movieStore.CreateMovie()

	log.Printf("Movie with ID: %s was created successfully", movie.ID)
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

	log.Printf("Movie with ID: %s was updated successfully", id)
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

	log.Printf("Movie with ID: %s was deleted successfully", id)
	return c.SendStatus(fiber.StatusNoContent)
}
