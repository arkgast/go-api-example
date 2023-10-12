package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"veloud.com/api/internal/data"
	"veloud.com/api/internal/routes"
)

func main() {
	app := fiber.New()

	movieStore := data.NewMovieStore()

	app.Use(func(c *fiber.Ctx) error {
		c.Locals("movieStore", movieStore)
		return c.Next()
	})

	app.Use(func(c *fiber.Ctx) error {
		movieStore.Lock()
		defer movieStore.Unlock()
		return c.Next()
	})

	app.Get("/:id", routes.GetMovie)
	app.Get("/", routes.GetMovies)
	app.Post("/", routes.CreateMovie)
	app.Put("/:id", routes.UpdateMovie)
	app.Delete("/:id", routes.DeleteMovie)

	log.Fatal(app.Listen(":3100"))
}
