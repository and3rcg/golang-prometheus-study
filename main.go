package main

import (
	"fiber-prometheus-demo/internal"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const version = "0.1.0"

func main() {
	fiberConfig := fiber.Config{
		AppName: fmt.Sprintf("Fiber DevOps showcase v%s", version),
	}

	app := fiber.New(fiberConfig)

	log.Println("Connecting to database...")
	internal.StartDB()

	log.Println("Starting Prometheus...")
	prometheus_instance := internal.StartPrometheus()
	app.Get("/metrics", adaptor.HTTPHandler(promhttp.Handler()))
	app.Use(prometheus_instance.Middleware)

	SetupRoutes(app)

	app.Listen(":4000")
}
