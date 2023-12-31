package routes

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"veloud.com/api/pkg/data"
)

func GetMovie(c *fiber.Ctx) error {
	id := c.Params("id")
	movieStore := c.Locals("movieStore").(*data.MovieStore)

	movie, err := movieStore.GetMovie(id)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	log.Printf("Movie with ID: %s was retrieved successfully", movie.ID)
	return c.JSON(movie)
}

func GetMovies(c *fiber.Ctx) error {
	movieStore := c.Locals("movieStore").(*data.MovieStore)
	log.Println("Get movies executed successfully")
	return c.Status(fiber.StatusOK).JSON(movieStore.GetMovies())
}

func CreateMovie(c *fiber.Ctx) error {
	newMovie := new(data.MovieInput)
	if err := c.BodyParser(newMovie); err != nil {
		return err
	}

	movieStore := c.Locals("movieStore").(*data.MovieStore)
	movie := movieStore.CreateMovie(*newMovie)

	log.Printf("Movie with ID: %s was created successfully", movie.ID)
	return c.Status(fiber.StatusCreated).JSON(movie)
}

func UpdateMovie(c *fiber.Ctx) error {
	id := c.Params("id")
	movieStore := c.Locals("movieStore").(*data.MovieStore)

	movieData := new(data.MovieInput)
	if err := c.BodyParser(movieData); err != nil {
		return err
	}

	movie, err := movieStore.UpdateMovie(id, *movieData)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	log.Printf("Movie with ID: %s was updated successfully", id)
	return c.Status(fiber.StatusOK).JSON(movie)
}

func DeleteMovie(c *fiber.Ctx) error {
	id := c.Params("id")
	movieStore := c.Locals("movieStore").(*data.MovieStore)

	err := movieStore.DeleteMovie(id)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	log.Printf("Movie with ID: %s was deleted successfully", id)
	return c.SendStatus(fiber.StatusNoContent)
}
