package rest

import (
	"fmt"
	"gymn/v1/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func InitServer() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	routes.SetUpRoutes(app)

	if err := app.Listen(":3000"); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Listening in port 3000")
}
