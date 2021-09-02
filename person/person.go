package person

import (
	"github.com/gofiber/fiber/v2"
)

type Person struct {
	Age  int    `json:"age"`
	Name string `json:"name"`
}

func GetAllPeople(c *fiber.Ctx) error {
	return getAllPeopleService(c)
}
