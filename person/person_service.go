package person

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

func createArrayOfPeople() []Person {
	kinga := Person{Age: 20, Name: "Kinga"}
	filip := Person{Age: 20, Name: "Filip"}
	david := Person{Age: 20, Name: "Dawid"}
	karol := Person{Age: 19, Name: "Karol"}

	return []Person{kinga, david, filip, karol}
}

func getAllPeopleService(c *fiber.Ctx) error {
	array_of_people := createArrayOfPeople()
	response, err := json.MarshalIndent(array_of_people, "\t", "\t")

	if err != nil {
		return c.Status(500).SendString("Internal server error :(")
	}

	c.Context().SetContentType("Application/JSON")

	return c.Send(response)
}
