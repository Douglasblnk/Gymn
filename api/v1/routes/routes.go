package routes

import "github.com/gofiber/fiber/v2"

func SetUpRoutes(app *fiber.App) {
	PublicRoutes(app)
	PrivateRoutes(app)
}
