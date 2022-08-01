package routes

import (
	"gymn/v1/handlers"

	"github.com/gofiber/fiber/v2"
)

func StudentRoutes(app *fiber.App) {
	app.Post("/create-student", handlers.CreateStudent)
	app.Post("/update-student/:id", handlers.UpdateStudent)
	app.Get("/get-student-code/:id", handlers.GetStudentCode)
}
