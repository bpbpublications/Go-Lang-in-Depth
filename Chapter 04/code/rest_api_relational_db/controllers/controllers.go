package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rest_api_relational_db/models"
)

type CreateEmployeeInput struct {
	Name       string `json:"name" binding:"required"`
	Department string `json:"department" binding:"required"`
}

type UpdateEmployeeInput struct {
	Name       string `json:"name"`
	Department string `json:"department"`
}

func FindEmployees(context *gin.Context) {
	var employees []models.Employee
	models.DB.Find(&employees)

	context.JSON(http.StatusOK, gin.H{"data": employees})
}

func CreateEmployee(context *gin.Context) {
	var input CreateEmployeeInput
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	employee := models.Employee{Name: input.Name, Department: input.Department}
	models.DB.Create(&employee)

	context.JSON(http.StatusOK, gin.H{"data": employee})
}

func FindEmployee(context *gin.Context) {
	var employee models.Employee

	if err := models.DB.Where("id = ?", context.Param("id")).First(&employee).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": employee})
}

func UpdateEmployee(context *gin.Context) {

	var employee models.Employee
	if err := models.DB.Where("id = ?", context.Param("id")).First(&employee).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input UpdateEmployeeInput
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&employee).Updates(input)

	context.JSON(http.StatusOK, gin.H{"data": employee})
}

func DeleteEmployee(context *gin.Context) {

	var employee models.Employee
	if err := models.DB.Where("id = ?", context.Param("id")).First(&employee).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&employee)

	context.JSON(http.StatusOK, gin.H{"data": true})
}
