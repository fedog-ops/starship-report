package handlers

import (
	"api/database"
	"api/models"
	"github.com/gofiber/fiber/v2"
)

func ListCrew(c *fiber.Ctx) error {

	crew := []models.CrewMember{}
	if err := database.DB.Db.Preload("Reports").Find(&crew).Error; err!= nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
    }

	return c.Status(200).JSON(crew)
}

func CreateCrew(c *fiber.Ctx) error {
	crew := new(models.CrewMember)
	if err := c.BodyParser(crew); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	database.DB.Db.Create(&crew)

	return c.Status(200).JSON(crew)
}

// without gorm

// func CreateCrew(c *fiber.Ctx) error {
// 	crew := new(CrewMember)
// 	if err := c.BodyParser(crew); err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"message": err.Error(),
// 		})
// 	}

// 	query := `INSERT INTO crew_members (name) VALUES ($1) RETURNING id`
// 	err := db.QueryRow(query, crew.Name).Scan(&crew.ID)
// 	if err != nil {
// 		log.Printf("Error inserting crew member: %v", err)
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"message": "Failed to create crew member",
// 		})
// 	}

// 	return c.Status(fiber.StatusOK).JSON(crew)
// }