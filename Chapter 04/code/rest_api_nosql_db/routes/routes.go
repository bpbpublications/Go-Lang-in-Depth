package routes

import (
	"github.com/gofiber/fiber/v2"
	"rest_api_nosql_db/controllers"
)

func SupplierRoute(app *fiber.App) {
	app.Post("/supplier", controllers.CreateSupplier)
	app.Get("/supplier/:supplierId", controllers.GetSupplier)
	app.Put("/supplier/:supplierId", controllers.EditSupplier)
	app.Delete("/supplier/:supplierId", controllers.DeleteSupplier)
	app.Get("/suppliers", controllers.GetAllSuppliers)
}
