package middleware

import (
	"github.com/gofiber/fiber/v2"
	"strings"
)

const role = "admin"

func CheckIsAdmin(ctx *fiber.Ctx) error {
	val := ctx.Request().Header.Peek("User-Role")
	if !strings.EqualFold(string(val), role) {
		ctx.Status(fiber.StatusForbidden)

		return ctx.JSON(fiber.Map{"Message": "Access is denied"})
	}

	err := ctx.Next()
	if err != nil {
		return err
	}

	return nil
}
