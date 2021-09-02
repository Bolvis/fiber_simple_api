package person

import (
	"github.com/gofiber/fiber/v2"
)

type Person struct {
	Age  int    `json:"age"`
	Name string `json:"name"`
}

func Get_all_people(c *fiber.Ctx) error {
	return get_all_people_service(c)
}
