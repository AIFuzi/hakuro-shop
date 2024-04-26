package routes

import "github.com/gofiber/fiber/v2"

func Setup(app *fiber.App) {
	app.Get("/healthcheck", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Status ok")
	})
}
