package routes

import (
	"github.com/jvicrosario1106/jwt-golang/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	app.Post("/register", controllers.Register)
	app.Post("/login", controllers.Login)

}
