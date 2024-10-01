package handlers

import (
	"api/database"
	"api/models"
	"github.com/gofiber/fiber/v2"
)

func ListReports(c *fiber.Ctx) error {

	reports := []models.Report{}
	database.DB.Db.Find(&reports)

	return c.Status(200).JSON(reports)
}

func CreateReport(c *fiber.Ctx) error {
	report := new(models.Report)
	if err := c.BodyParser(report); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	database.DB.Db.Create(&report)

	return c.Status(200).JSON(report)
}