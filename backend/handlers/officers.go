package handlers

import (
	"api/database"
	"api/models"
	"github.com/gofiber/fiber/v2"
)

func Home(c *fiber.Ctx) error {
	return c.SendString("Welcome Aboard! -> Available endpoints /officers /crew /reports")
}

func ListOfficers(c *fiber.Ctx) error {

	facts := []models.Officer{}
	if err := database.DB.Db.Preload("Reports").Find(&facts).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
    }

	return c.Status(200).JSON(facts)
}

func CreateOfficer(c *fiber.Ctx) error {
	fact := new(models.Officer)
	if err := c.BodyParser(fact); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	database.DB.Db.Create(&fact)

	return c.Status(200).JSON(fact)
}