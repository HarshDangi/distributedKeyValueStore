package handlers

import "github.com/gofiber/fiber/v2"

func AttachHandlers(app *fiber.App) {
	app.Get("/:id", func(c *fiber.Ctx) error {
		return c.SendString("Hello World!")
	})
}
