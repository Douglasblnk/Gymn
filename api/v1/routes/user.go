package routes

import (
	handlers "gymn/v1/handlers/user"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
	app.Post("/register", handlers.RegisterUser)
	app.Put("/update-user/:id", handlers.UpdateUser)
}
