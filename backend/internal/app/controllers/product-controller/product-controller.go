package product_controller

import (
	"github.com/AIFuzi/hakuro-shop/internal/app/database"
	"github.com/AIFuzi/hakuro-shop/internal/app/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"os"
)

const queryIdFilter = "id"
const queryTypeIdFilter = "typeID"

func GetAll(ctx *fiber.Ctx) error {
	var product []models.Product

	id := ctx.Query(queryTypeIdFilter)
	res := database.DB

	if len(id) > 0 {
		res = database.DB.Where("type_id = ?", id).Find(&product)
	} else {
		res = database.DB.Find(&product)
	}

	if res.Error != nil {
		ctx.Status(fiber.StatusBadRequest)

		return ctx.JSON(fiber.Map{"Message": "Bad request"})
	}

	return ctx.JSON(product)
}

func GetOne(ctx *fiber.Ctx) error {
	id, _ := ctx.ParamsInt(queryIdFilter)
	product := models.Product{}

	res := database.DB.Where("id = ?", id).First(&product)
	if res.Error != nil {
		ctx.Status(fiber.StatusNotFound)

		return ctx.JSON(fiber.Map{"Message": "RecordNotFound"})
	}

	return ctx.JSON(product)
}

func Create(ctx *fiber.Ctx) error {
	p := models.Product{}
	err := ctx.BodyParser(&p)
	if err != nil {
		return err
	}

	file, err := ctx.FormFile("Image")
	if err != nil {
		ctx.Status(fiber.StatusBadRequest)

		return ctx.JSON(fiber.Map{"Message": "Failed to read file"})
	}

	fn := uuid.New().String() + ".png"
	if err = ctx.SaveFile(file, "./static/"+fn); err != nil {
		ctx.Status(fiber.StatusBadRequest)

		return ctx.JSON(fiber.Map{"Message": "Failed to upload file"})
	}

	product := models.Product{
		TypeID:      p.TypeID,
		Image:       fn,
		ProductName: p.ProductName,
		Price:       p.Price,
	}

	res := database.DB.Create(&product)
	if res.Error != nil {
		ctx.Status(fiber.StatusBadRequest)

		return ctx.JSON(fiber.Map{"Message": res.Error})
	}

	return ctx.JSON(product)
}

func Delete(ctx *fiber.Ctx) error {
	id, _ := ctx.ParamsInt(queryIdFilter)
	product := models.Product{ID: uint(id)}

	database.DB.Where("id = ?", id).Find(&product)

	err := os.Remove("./static/" + product.Image)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest)

		return ctx.JSON(fiber.Map{"Message": err})
	}

	res := database.DB.Delete(&product)
	if res.RowsAffected < 1 {
		ctx.Status(fiber.StatusNotFound)

		return ctx.JSON(fiber.Map{"Message": "RecordNotFound"})
	}

	return ctx.JSON(fiber.Map{"Message": "Product deleted"})
}
