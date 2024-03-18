package routes

import (
	"jwt-token/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)
	app.Get("/api/validate", controllers.Validate)
	app.Post("/api/logout", controllers.Logout)
}
