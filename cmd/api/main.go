package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"veloud.com/api/internal/routes"
)

func main() {
	app := fiber.New()

	app.Get("/:id", routes.GetMovie)
	app.Get("/", routes.GetMovies)
	app.Post("/", routes.CreateMovie)
	app.Put("/:id", routes.UpdateMovie)
	app.Delete("/:id", routes.DeleteMovie)

	log.Fatal(app.Listen(":3100"))
}
