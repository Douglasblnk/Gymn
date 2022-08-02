package routes

import (
	handlers "gymn/v1/handlers/student"

	"github.com/gofiber/fiber/v2"
)

func StudentRoutes(app *fiber.App) {
	app.Post("/create-student", handlers.CreateStudent)
	app.Put("/update-student/:id", handlers.UpdateStudent)
	app.Get("/get-student-code/:id", handlers.GetStudentCode)
	app.Get("/get-students", handlers.GetStudents)
	app.Get("/get-student/:id", handlers.GetStudent)
}
