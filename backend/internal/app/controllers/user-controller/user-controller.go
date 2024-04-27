package user_controller

import (
	"github.com/AIFuzi/hakuro-shop/internal/app/database"
	"github.com/AIFuzi/hakuro-shop/internal/app/models"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Register(ctx *fiber.Ctx) error {
	d := models.User{}

	err := ctx.BodyParser(&d)
	if err != nil {
		return err
	}

	pass, _ := bcrypt.GenerateFromPassword(d.Password, 5)

	user := models.User{
		Email:    d.Email,
		Name:     d.Name,
		Password: pass,
		Role:     d.Role,
	}

	res := database.DB.Create(&user)
	if res.Error != nil {
		ctx.Status(fiber.StatusBadRequest)

		return ctx.JSON(fiber.Map{"Message": res.Error})
	}

	return ctx.JSON(user)
}

func Login(ctx *fiber.Ctx) error {
	return nil
}

func Logout(ctx *fiber.Ctx) error {
	return nil
}

func Check(ctx *fiber.Ctx) error {
	return nil
}
