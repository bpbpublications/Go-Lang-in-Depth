package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

var template_html = template.Must(template.ParseGlob("templates/*"))

func Home(writer http.ResponseWriter, request *http.Request) {
	var employees []Employee
	employees = GetEmployees()
	log.Println(employees)
	template_html.ExecuteTemplate(writer, "Home", employees)

}

func Create(writer http.ResponseWriter, request *http.Request) {

	template_html.ExecuteTemplate(writer, "Create", nil)

}

func Insert(writer http.ResponseWriter, request *http.Request) {

	var employee Employee
	employee.EmployeeName = request.FormValue("name")
	employee.SSN = request.FormValue("ssn")
	employee.Designation = request.FormValue("designation")
	employee.Department = request.FormValue("department")
	InsertEmployee(employee)
	var employees []Employee
	employees = GetEmployees()
	template_html.ExecuteTemplate(writer, "Home", employees)

}
func Alter(writer http.ResponseWriter, request *http.Request) {

	var employee Employee
	var employeeId int
	var employeeIdStr string
	employeeIdStr = request.FormValue("id")
	fmt.Sscanf(employeeIdStr, "%d", &employeeId)
	employee.EmployeeId = employeeId
	employee.EmployeeName = request.FormValue("name")
	employee.SSN = request.FormValue("ssn")
	employee.Designation = request.FormValue("designation")
	employee.Department = request.FormValue("department")
	UpdateEmployee(employee)
	var employees []Employee
	employees = GetEmployees()
	template_html.ExecuteTemplate(writer, "Home", employees)

}

func Update(writer http.ResponseWriter, request *http.Request) {

	var employeeId int
	var employeeIdStr string
	employeeIdStr = request.FormValue("id")
	fmt.Sscanf(employeeIdStr, "%d", &employeeId)
	var employee Employee
	employee = GetEmployeeById(employeeId)
	template_html.ExecuteTemplate(writer, "Update", employee)

}

func Delete(writer http.ResponseWriter, request *http.Request) {
	var employeeId int
	var employeeIdStr string
	employeeIdStr = request.FormValue("id")
	fmt.Sscanf(employeeIdStr, "%d", &employeeId)
	var employee Employee
	employee = GetEmployeeById(employeeId)
	DeleteEmployee(employee)
	var employees []Employee
	employees = GetEmployees()
	template_html.ExecuteTemplate(writer, "Home", employees)

}

func View(writer http.ResponseWriter, request *http.Request) {

	var employeeId int
	var employeeIdStr string
	employeeIdStr = request.FormValue("id")
	fmt.Sscanf(employeeIdStr, "%d", &employeeId)
	var employee Employee
	employee = GetEmployeeById(employeeId)
	fmt.Println(employee)
	var employees []Employee
	employees = []Employee{employee}
	template_html.ExecuteTemplate(writer, "View", employees)

}

func main() {
	log.Println("Server started on: http://localhost:8000")

	http.HandleFunc("/", Home)
	http.HandleFunc("/alter", Alter)
	http.HandleFunc("/create", Create)
	http.HandleFunc("/update", Update)
	http.HandleFunc("/view", View)
	http.HandleFunc("/insert", Insert)
	http.HandleFunc("/delete", Delete)
	http.ListenAndServe(":8000", nil)
}
