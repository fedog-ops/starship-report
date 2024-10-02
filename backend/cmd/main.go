package main

import (
	"api/database"
	"github.com/gofiber/fiber/v2"
	"api/handlers"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func main() {
	database.ConnectDb() 
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3001, http://example.com", // Allowed origins
		AllowMethods:     "GET, POST, PUT, DELETE",                   // Allowed methods
		AllowHeaders:     "Origin, Content-Type, Authorization",      // Allowed headers
		AllowCredentials: true,                                       // Whether credentials (cookies, auth headers) are allowed
	}))

	// Apply Basic Authentication middleware to protected routes
	app.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			"admin": "captainvader", // Replace with actual username and password
			"captain": "darkside",
		},
	}))

	SetupRoutes(app)

	app.Listen(":3000")
}

func SetupRoutes(app *fiber.App){
	app.Get("/", handlers.Home)
	app.Get("/officers", handlers.ListOfficers)
	app.Post("/officer", handlers.CreateOfficer)

	app.Get("/crew", handlers.ListCrew)
	app.Post("/crew", handlers.CreateCrew)
	
	app.Get("/reports", handlers.ListReports)
	app.Post("/reports", handlers.CreateReport)
	app.Put("/reports/:reportId", handlers.EditReport)

	app.Get("/captain", handlers.ListCaptain)
	app.Get("/hellocap", handlers.HelloCap)
}