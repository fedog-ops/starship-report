package handlers

import (
	"api/database"
	"api/models"
	"github.com/gofiber/fiber/v2"
)

func HelloCap(c *fiber.Ctx) error {
	return c.SendString("I am the captain now")
}

func ListCaptain(c *fiber.Ctx) error {
	cap := []models.Captain{}
	if err := database.DB.Db.Find(&cap).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
    }

	return c.Status(200).JSON(cap)
}