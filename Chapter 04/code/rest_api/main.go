package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	restapi := fiber.New()
	restapi.Use(cors.New())

	customers := restapi.Group("/customers")

	customers.Get("/", func(context *fiber.Ctx) error {
		return context.SendString("Customers are being loaded")
	})

	log.Fatal(restapi.Listen(":5000"))
}
