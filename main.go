package main

import (
	"fiber_tutorial/person"

	"github.com/gofiber/fiber/v2"
)

func create_mapping(app *fiber.App) {
	app.Get("/api/v1/people", person.Get_all_people)
}

func main() {
	app := fiber.New()

	create_mapping(app)

	app.Listen(":3000")
}
