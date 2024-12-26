package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	var employees []Employee
	employees = GetEmployees()
	fmt.Println(employees)

	var employee Employee
	employee.EmployeeName = "Thomas Smith"
	employee.SSN = "2323343"

	InsertEmployee(employee)

	var employee Employee
	employee.EmployeeName = "George Thompson"
	employee.SSN = "23233432"
	employee.EmployeeId = 2
	UpdateEmployee(employee)

	var employee Employee
	employee.EmployeeName = "George Thompson"
	employee.SSN = "23233432"
	employee.EmployeeId = 2

	deleteEmployee(employee)
	employees = GetEmployees()
	fmt.Println(employees)

}
