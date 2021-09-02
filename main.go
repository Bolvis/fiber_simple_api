package main

import (
	"fiber_tutorial/person"

	"github.com/gofiber/fiber/v2"
)

func createMapping(app *fiber.App) {
	app.Get("/api/v1/people", person.GetAllPeople)
}

func main() {
	app := fiber.New()

	createMapping(app)

	app.Listen(":3000")
}
