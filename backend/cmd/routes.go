package main

import (
	"api/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App){
	app.Get("/", handlers.Home)
	app.Get("/officers", handlers.ListOfficers)
	app.Post("/officer", handlers.CreateOfficer)

	app.Get("/crew", handlers.ListCrew)
	app.Post("/crew", handlers.CreateCrew)
	
	app.Get("/reports", handlers.ListReports)
	app.Post("/reports", handlers.CreateReport)

	app.Get("/captain", handlers.ListCaptain)
	app.Get("/hellocap", handlers.HelloCap)
}