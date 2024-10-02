package handlers

import (
	"api/database"
	"api/models"
	"fmt"
	"time"

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
	start := time.Now()
	report.LastEditedAt = start.Add(time.Hour * 24 * 30)
	database.DB.Db.Create(&report)

	return c.Status(200).JSON(report)
}

func EditReport(c *fiber.Ctx) error {

	report := new(models.Report)
	id := c.Params("reportId")

	  if err := database.DB.Db.First(&report, "id = ?", id).Error; err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "status":  "error",
            "message": "No report present",
            "data":    nil,
        })
    }

	if report.LastEditedAt.Before(time.Now().Add(-30 * 24 * time.Hour)) {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Cannot edit report after 30 days",
			"data":    nil,
		})
	}

	updatedReportC := new(models.Report)
	if err := c.BodyParser(&updatedReportC); err != nil {
		fmt.Println(updatedReportC)
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "message": err.Error(),
        })
    }

	report.Content = updatedReportC.Content

 	if result := database.DB.Db.Save(&report); result.Error != nil {
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"message": result.Error.Error(),
	})
}

return c.Status(fiber.StatusOK).JSON(fiber.Map{
	"status":  "success",
	"message": "Report updated successfully",
	"data":    report,
})
}