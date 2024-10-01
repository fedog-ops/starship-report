package handlers

import (
	"api/database"
	"api/models"
	"github.com/gofiber/fiber/v2"
)

func Home(c *fiber.Ctx) error {
	return c.SendString("Welcome Aboard!    ->    Available /officers")
}

func ListOfficers(c *fiber.Ctx) error {

	facts := []models.Officer{}
	database.DB.Db.Find(&facts)

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