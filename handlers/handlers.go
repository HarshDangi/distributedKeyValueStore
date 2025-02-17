package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func AttachHandlers(app *fiber.App) {
	app.Get("/:key", func(c *fiber.Ctx) error {
		return c.SendString("Hello World!")
	})

	app.Put("/", func(c *fiber.Ctx) error {
		var putdata map[string]string
		err := c.BodyParser(&putdata)
		if err != nil {
			return c.Status(400).JSON(&fiber.Map{
				"success": true,
				"message": "Ivalid value",
			})
		}
		for k, v := range putdata {
			fmt.Printf("%T %v %T %v\n", k, k, v, v)
		}
		return c.Status(200).JSON(&fiber.Map{
			"success": true,
			"message": "",
		})
	})

	app.Post("/", func(c *fiber.Ctx) error {
		var putdata map[string]string
		err := c.BodyParser(&putdata)
		if err != nil {
			return c.Status(400).JSON(&fiber.Map{
				"success": true,
				"message": "Ivalid value",
			})
		}
		for k, v := range putdata {
			fmt.Printf("%T %v %T %v\n", k, k, v, v)
		}
		return c.Status(200).JSON(&fiber.Map{
			"success": true,
			"message": "",
		})
	})

	app.Get("/:key", func(c *fiber.Ctx) error {
		return c.SendString("Delete functionality pending!")
	})
}
