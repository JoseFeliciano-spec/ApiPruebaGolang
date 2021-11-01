package routes

import "github.com/gofiber/fiber/v2"

func Routes(app *fiber.App) {
	app.Post("/", func(c *fiber.Ctx) error {
		return c.JSON("Holisis")
	})
}
