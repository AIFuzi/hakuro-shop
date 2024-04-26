package routes

import (
	typeControllers "github.com/AIFuzi/hakuro-shop/internal/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/healthcheck", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Status ok")
	})

	app.Get("/api/types", typeControllers.GetAll)
	app.Post("/api/types/create", typeControllers.Create)
}
