package main

import (
	"github.com/gofiber/fiber/v2"
	"rest_api_nosql_db/routes"
)

func main() {
	app := fiber.New()

	routes.SupplierRoute(app)

	app.Listen(":6000")
}
