package main

import (
	"github.com/gofiber/fiber/v2"
	"veloud.com/api/internal/routes"
)

func main() {
	app := fiber.New()

	app.Get("/", routes.IndexHandler)

	app.Listen(":3100")
}
