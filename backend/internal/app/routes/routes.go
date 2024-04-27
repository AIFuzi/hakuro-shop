package routes

import (
	"github.com/AIFuzi/hakuro-shop/internal/app/controllers/product-controller"
	"github.com/AIFuzi/hakuro-shop/internal/app/controllers/type-controller"
	user_controller "github.com/AIFuzi/hakuro-shop/internal/app/controllers/user-controller"
	"github.com/AIFuzi/hakuro-shop/internal/app/middleware"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/healthcheck", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Status ok")
	})

	app.Get("/api/types", type_controller.GetAll)
	app.Post("/api/types/create", middleware.CheckIsAdmin, type_controller.Create)
	app.Delete("/api/types/delete/:id", middleware.CheckIsAdmin, type_controller.Delete)

	app.Get("/api/product", product_controller.GetAll)
	app.Get("/api/product:id", product_controller.GetOne)
	app.Post("/api/product/create", middleware.CheckIsAdmin, product_controller.Create)
	app.Delete("/api/product/delete/:id", middleware.CheckIsAdmin, product_controller.Delete)

	app.Post("api/user/signup", user_controller.Register)
	app.Post("api/user/login", user_controller.Login)
	app.Post("api/user/logout", user_controller.Logout)
	app.Get("api/user/check", user_controller.Check)
}
