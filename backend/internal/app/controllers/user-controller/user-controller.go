package user_controller

import (
	"github.com/AIFuzi/hakuro-shop/internal/app/database"
	"github.com/AIFuzi/hakuro-shop/internal/app/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
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
	d := models.User{}

	err := ctx.BodyParser(&d)
	if err != nil {
		return err
	}

	var user models.User
	res := database.DB.Where("email = ?", d.Email).First(&user)
	if res.Error != nil {
		ctx.Status(fiber.StatusNotFound)

		return ctx.JSON(fiber.Map{"Message": "User not found"})
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, d.Password); err != nil {
		ctx.Status(fiber.StatusBadRequest)

		return ctx.JSON(fiber.Map{"Message": "Incorrect password"})
	}

	uc := jwt.MapClaims{
		"iss":   "issuer",
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
		"Email": user.Email,
		"Role":  user.Role,
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)

	token, err := claims.SignedString([]byte(os.Getenv("jwt_secret_key")))
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError)

		return ctx.JSON(fiber.Map{"Message": err})
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	ctx.Cookie(&cookie)

	return ctx.JSON(token)
}

func Logout(ctx *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	ctx.Cookie(&cookie)

	return nil
}

func Check(ctx *fiber.Ctx) error {
	return nil
}
