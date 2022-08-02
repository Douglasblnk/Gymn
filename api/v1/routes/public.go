package routes

import (
	authHandlers "gymn/v1/handlers/auth"
	userHandlers "gymn/v1/handlers/user"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app *fiber.App) {
	app.Post("/email-validation", authHandlers.EmailValidation)
	app.Post("/sign-in", authHandlers.SignInHandler)
}

func UserPublicRoutes(app *fiber.App) {
	app.Post("/register", userHandlers.RegisterUser)
}

func PublicRoutes(app *fiber.App) {
	AuthRoutes(app)
	UserPublicRoutes(app)
}
