package server

import "github.com/gofiber/fiber/v2"

func (app *App) indexPage() fiber.Handler {
	return func(c *fiber.Ctx) error {

		certsDisplayData, err := app.storage.AllForIndexPage()

		if err != nil {
			return c.Status(500).Render("views/error", fiber.Map{})
		}

		return c.Status(200).Render("views/index", fiber.Map{
			"CertsDisplayData": certsDisplayData,
		}, "views/layout")
	}
}
