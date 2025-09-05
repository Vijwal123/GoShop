package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yourname/goshop/database"
	"github.com/yourname/goshop/models"
)

func CreateProduct(c *fiber.Ctx) error {
	product := new(models.Product)

	if err := c.BodyParser(product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse JSON"})
	}

	if result := database.DB.Create(&product); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": result.Error.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(product)
}

func GetProducts(c *fiber.Ctx) error {
	var products []models.Product
	if result := database.DB.Find(&products); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": result.Error.Error()})
	}
	return c.JSON(products)
}
