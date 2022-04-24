package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/handlebars"
)

func main() {
	engine := handlebars.New("./views", ".hbs")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static(
		"/",        // mount address
		"./public", // path to the file folder
	)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("pages/index", fiber.Map{
			"message": "Hello World! The world is flat.",
		}, "layouts/mgmt")
	})

	app.Get("/main", func(c *fiber.Ctx) error {
		return c.Render("pages/main", fiber.Map{
			"message": "Hello World! The world is flat.",
		}, "layouts/mgmt")
	})

	app.Listen(":8080")
}
