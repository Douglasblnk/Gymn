package routes

import (
	studentHandlers "gymn/v1/handlers/student"
	userHandlers "gymn/v1/handlers/user"
	"gymn/v1/middleware"

	"github.com/gofiber/fiber/v2"
)

func StudentRoutes(app *fiber.App) {
	app.Post("/create-student", middleware.JWTAuth, studentHandlers.CreateStudent)
	app.Put("/update-student/:id", middleware.JWTAuth, studentHandlers.UpdateStudent)
	app.Get("/get-student-code/:id", middleware.JWTAuth, studentHandlers.GetStudentCode)
	app.Get("/get-students", middleware.JWTAuth, studentHandlers.GetStudents)
	app.Get("/get-student/:id", middleware.JWTAuth, studentHandlers.GetStudent)
}

func UserPrivateRoutes(app *fiber.App) {
	app.Put("/update-user/:id", middleware.JWTAuth, userHandlers.UpdateUser)
}

func PrivateRoutes(app *fiber.App) {
	StudentRoutes(app)
	UserPrivateRoutes(app)
}
