package routes

import (
	handlers "gymn/v1/handlers/auth"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app *fiber.App) {
	app.Post("/email-validation", handlers.EmailValidation)
	app.Post("/sign-in", handlers.SignInHandler)
}
