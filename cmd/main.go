// add code hello world
package main

import (
	"github.com/gofiber/fiber/v2"
)

// setupApp configures and returns a Fiber app instance
func setupApp() *fiber.App {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	})

	return app
}

func main() {
	app := setupApp()
	app.Listen(":3000")
}
