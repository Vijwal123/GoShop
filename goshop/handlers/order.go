package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/yourname/goshop/database"
	"github.com/yourname/goshop/models"
)

func Checkout(c *fiber.Ctx) error {
	userID := c.Params("userId")

	var cart []models.Cart
	if err := database.DB.Preload("Product").Where("user_id = ?", userID).Find(&cart).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "cannot fetch cart"})
	}

	if len(cart) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cart is empty"})
	}

	var total float64
	var orderItems []models.OrderItem
	for _, item := range cart {
		total += item.Product.Price * float64(item.Quantity)
		orderItems = append(orderItems, models.OrderItem{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     item.Product.Price,
		})
	}

	order := models.Order{
		UserID: parseID(userID),
		Total:  total,
		Items:  orderItems,
	}
	if err := database.DB.Create(&order).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to create order"})
	}

	database.DB.Where("user_id = ?", userID).Delete(&models.Cart{})

	return c.JSON(order)
}

func GetOrders(c *fiber.Ctx) error {
	userID := c.Params("userId")
	var orders []models.Order
	if err := database.DB.Preload("Items.Product").Where("user_id = ?", userID).Find(&orders).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to fetch orders"})
	}
	return c.JSON(orders)
}

func parseID(id string) uint {
	var result uint
	fmt.Sscanf(id, "%d", &result)
	return result
}
