package typeControllers

import (
	"github.com/AIFuzi/hakuro-shop/internal/database"
	"github.com/AIFuzi/hakuro-shop/internal/models"
	"github.com/gofiber/fiber/v2"
)

const queryFilter = "id"

func GetAll(ctx *fiber.Ctx) error {
	var productTypes []models.Types

	query := ctx.Queries()
	if len(query[queryFilter]) > 0 {
		res := database.DB.Where("id = ?", query[queryFilter]).First(&productTypes)
		if res.Error != nil {
			ctx.Status(fiber.StatusBadRequest)

			return ctx.JSON(fiber.Map{"Message": "RecordNotFound"})
		}
	} else {
		database.DB.Find(&productTypes)
	}

	return ctx.JSON(productTypes)
}

func Create(ctx *fiber.Ctx) error {
	var data map[string]string
	err := ctx.BodyParser(&data)
	if err != nil {
		return err
	}

	productType := models.Types{
		TypeName: data["typeName"],
	}

	res := database.DB.Create(&productType)
	if res.Error != nil {
		ctx.Status(fiber.StatusBadRequest)

		return ctx.JSON(fiber.Map{"Message": res.Error})
	}
	return ctx.JSON(productType)
}
