package main

import (
	"fiber-prometheus-demo/internal"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

const version = "0.1.0"

func main() {
	fiberConfig := fiber.Config{
		AppName: fmt.Sprintf("Fiber DevOps showcase v%s", version),
	}

	app := fiber.New(fiberConfig)

	fmt.Println("Connecting to database...")
	internal.StartDB()

	SetupRoutes(app)

	app.Listen(":4000")
}
