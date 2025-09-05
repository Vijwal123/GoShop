package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yourname/goshop/database"
	"github.com/yourname/goshop/models"
)

func AddToCart(c *fiber.Ctx) error {
	var cart models.Cart
	if err := c.BodyParser(&cart); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid input"})
	}

	if cart.Quantity <= 0 {
		cart.Quantity = 1
	}

	if err := database.DB.Create(&cart).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to add to cart"})
	}

	return c.Status(fiber.StatusCreated).JSON(cart)
}

func GetCart(c *fiber.Ctx) error {
	userID := c.Params("userId")
	var cart []models.Cart
	if err := database.DB.Preload("Product").Where("user_id = ?", userID).Find(&cart).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to fetch cart"})
	}
	return c.JSON(cart)
}

func RemoveFromCart(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := database.DB.Delete(&models.Cart{}, id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to remove item"})
	}
	return c.JSON(fiber.Map{"message": "item removed"})
}
