package person

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

type Person struct {
	Age  int    `json:"age"`
	Name string `json:"name"`
}

func GetAllPeople(c *fiber.Ctx) error {
	array_of_people := CreateArrayOfPeople()
	response, err := json.MarshalIndent(array_of_people, "\t", "\t")

	if err != nil {
		return c.Status(500).SendString("Internal server error :(")
	}

	return c.Send(response)
}
