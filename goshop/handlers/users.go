package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yourname/goshop/database"
	"github.com/yourname/goshop/models"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(c *fiber.Ctx) error {
	user := new(models.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse JSON"})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to hash password"})
	}
	user.Password = string(hashedPassword)

	if result := database.DB.Create(&user); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": result.Error.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"id":    user.ID,
		"name":  user.Name,
		"email": user.Email,
		"role":  user.Role,
	})
}

func GetUsers(c *fiber.Ctx) error {
	var users []models.User
	if result := database.DB.Find(&users); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": result.Error.Error()})
	}
	return c.JSON(users)
}

func GetUserByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User

	if result := database.DB.First(&user, id); result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "user not found"})
	}
	return c.JSON(user)
}
