package routes

import (
	equipmentHandlers "gymn/v1/handlers/equipment"
	studentHandlers "gymn/v1/handlers/student"
	trainingSheetHandlers "gymn/v1/handlers/training-sheet"
	userHandlers "gymn/v1/handlers/user"
	workoutHandlers "gymn/v1/handlers/workout"
	"gymn/v1/middleware"

	"github.com/gofiber/fiber/v2"
)

func EquipmentRoutes(app *fiber.App) {
	app.Post("/equipment", middleware.JWTAuth, equipmentHandlers.CreateEquipment)
	app.Get("/equipment", middleware.JWTAuth, equipmentHandlers.GetAllEquipments)
	app.Get("/equipment/:id", middleware.JWTAuth, equipmentHandlers.GetEquipment)
	app.Put("/equipment/:id", middleware.JWTAuth, equipmentHandlers.UpdateEquipment)
	app.Delete("/equipment/:id", middleware.JWTAuth, equipmentHandlers.DeleteEquipment)
}

func WorkoutRoutes(app *fiber.App) {
	app.Post("/workout", middleware.JWTAuth, workoutHandlers.CreateWorkout)
	app.Post("/workout-training-sheet", middleware.JWTAuth, workoutHandlers.AssociateWorkoutTrainingSheet)
	app.Get("/workout", middleware.JWTAuth, workoutHandlers.GetAllWorkouts)
	app.Get("/training-sheet-workout/:id", middleware.JWTAuth, workoutHandlers.GetTrainingSheetWorkouts)
	app.Put("/workout/:id", middleware.JWTAuth, workoutHandlers.UpdateWorkout)
	app.Get("/workout/:id", middleware.JWTAuth, workoutHandlers.GetWorkout)
	app.Delete("/workout/:id", middleware.JWTAuth, workoutHandlers.DeleteWorkout)
}

func TrainingSheetRoutes(app *fiber.App) {
	app.Post("/training-sheet", middleware.JWTAuth, trainingSheetHandlers.CreateTrainingSheet)
	app.Post("/training-sheet-student", middleware.JWTAuth, trainingSheetHandlers.AssociateTrainingSheetStudent)
	app.Get("/student-training-sheet/:id", middleware.JWTAuth, trainingSheetHandlers.GetStudentTrainingSheets)
	app.Get("/training-sheet", middleware.JWTAuth, trainingSheetHandlers.GetAllTrainingSheets)
	app.Get("/training-sheet/:id", middleware.JWTAuth, trainingSheetHandlers.GetTrainingSheet)
	app.Put("/training-sheet/:id", middleware.JWTAuth, trainingSheetHandlers.UpdateTrainingSheet)
	app.Delete("/training-sheet/:id", middleware.JWTAuth, trainingSheetHandlers.DeleteTrainingSheet)
}

func StudentRoutes(app *fiber.App) {
	app.Post("/student", middleware.JWTAuth, studentHandlers.CreateStudent)
	app.Get("/student-code/:id", middleware.JWTAuth, studentHandlers.GetStudentCode)
	app.Get("/student", middleware.JWTAuth, studentHandlers.GetAllStudents)
	app.Get("/student/:id", middleware.JWTAuth, studentHandlers.GetStudent)
	app.Put("/student/:id", middleware.JWTAuth, studentHandlers.UpdateStudent)
	app.Delete("/student/:id", middleware.JWTAuth, studentHandlers.DeleteStudent)
}

func UserPrivateRoutes(app *fiber.App) {
	app.Put("/user/:id", middleware.JWTAuth, userHandlers.UpdateUser)
}

func PrivateRoutes(app *fiber.App) {
	EquipmentRoutes(app)
	WorkoutRoutes(app)
	TrainingSheetRoutes(app)
	StudentRoutes(app)
	UserPrivateRoutes(app)
}
