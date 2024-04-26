package typeControllers

import (
	"github.com/AIFuzi/hakuro-shop/internal/app/database"
	"github.com/AIFuzi/hakuro-shop/internal/app/models"
	"github.com/gofiber/fiber/v2"
)

const queryFilter = "id"

func GetAll(ctx *fiber.Ctx) error {
	var productTypes []models.Types

	q := ctx.Queries()
	if len(q[queryFilter]) > 0 {
		res := database.DB.Where("id = ?", q[queryFilter]).First(&productTypes)
		if res.Error != nil {
			ctx.Status(fiber.StatusNotFound)

			return ctx.JSON(fiber.Map{"Message": "RecordNotFound"})
		}
	} else {
		database.DB.Find(&productTypes)
	}

	return ctx.JSON(productTypes)
}

func Create(ctx *fiber.Ctx) error {
	productType := models.Types{}
	err := ctx.BodyParser(&productType)
	if err != nil {
		return err
	}

	res := database.DB.Create(&productType)
	if res.Error != nil {
		ctx.Status(fiber.StatusBadRequest)

		return ctx.JSON(fiber.Map{"Message": res.Error})
	}

	return ctx.JSON(productType)
}

func Delete(ctx *fiber.Ctx) error {
	id, _ := ctx.ParamsInt(queryFilter)
	productTypes := models.Types{ID: uint(id)}

	res := database.DB.Delete(&productTypes)
	if res.RowsAffected < 1 {
		ctx.Status(fiber.StatusNotFound)

		return ctx.JSON(fiber.Map{"Message": "RecordNotFound"})
	}

	return ctx.JSON(fiber.Map{"Message": "Type deleted"})
}
