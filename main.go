package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/jvicrosario1106/jwt-golang/database"
	"github.com/jvicrosario1106/jwt-golang/routes"
)

func main() {
	database.ConnectionDB()
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
	app.Use(logger.New())
	routes.SetupRoutes(app)

	log.Print(app.Listen(":8000"))

}
