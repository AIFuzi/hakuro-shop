package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"os"
)

const role = "ADMIN"

func CheckIsAdmin(ctx *fiber.Ctx) error {
	token, err := jwt.Parse(ctx.Cookies(os.Getenv("USER_COOKIE_TOKEN_NAME")), func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})

	if err != nil {
		ctx.Status(fiber.StatusUnauthorized)

		return ctx.JSON(fiber.Map{"Message": "Unauthorized"})
	}

	claims := token.Claims.(jwt.MapClaims)

	if claims["Role"] != role {
		ctx.Status(fiber.StatusForbidden)

		return ctx.JSON(fiber.Map{"Message": "Access is denied"})
	}

	err = ctx.Next()
	if err != nil {
		return err
	}

	return nil
}
