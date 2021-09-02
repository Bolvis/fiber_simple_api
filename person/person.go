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
	response, _ := json.Marshal(array_of_people)
	return c.Send(response)
}
