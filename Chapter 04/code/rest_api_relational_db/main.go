package main

import (
	"github.com/gin-gonic/gin"
	"rest_api_relational_db/controllers"
	"rest_api_relational_db/models"
)

func main() {

	gin_instance := gin.Default()
	models.ConnectDatabase()

	gin_instance.GET("/employees", controllers.FindEmployees)
	gin_instance.POST("/employees", controllers.CreateEmployee)
	gin_instance.GET("/employees/:id", controllers.FindEmployee)
	gin_instance.PATCH("/employees/:id", controllers.UpdateEmployee)
	gin_instance.DELETE("/employees/:id", controllers.DeleteEmployee)

	err := gin_instance.Run()
	if err != nil {
		return
	}
}
