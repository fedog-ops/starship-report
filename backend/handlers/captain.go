package handlers

import (
	// "api/database"
	// "api/models"
	"github.com/gofiber/fiber/v2"
)

func HelloCap(c *fiber.Ctx) error {
	return c.SendString("I am the captain now")
}