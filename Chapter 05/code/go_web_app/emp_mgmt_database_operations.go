package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Employee struct {
	EmployeeId   int
	EmployeeName string
	SSN          string
	Designation  string
	Department   string
}

func GetConnection() (database *sql.DB) {
	databaseDriver := "mysql"
	databaseUser := "newuser"
	databasePass := "newuser"
	databaseName := "empMgmt"
	database, error := sql.Open(databaseDriver, databaseUser+":"+databasePass+"@/"+databaseName)
	if error != nil {
		panic(error.Error())
	}
	return database
}

func GetEmployeeById(employeeId int) Employee {
	var database *sql.DB
	database = GetConnection()

	var error error
	var rows *sql.Rows
	rows, error = database.Query("SELECT * FROM Employee WHERE EmployeeId=?", employeeId)
	if error != nil {
		panic(error.Error())
	}

	var employee Employee
	employee = Employee{}

	for rows.Next() {
		var employeeId int
		var employeeName string
		var SSN string
		var Designation string
		var Department string
		error = rows.Scan(&employeeId, &employeeName, &SSN, &Designation, &Department)
		if error != nil {
			panic(error.Error())
		}
		employee.EmployeeId = employeeId
		employee.EmployeeName = employeeName
		employee.SSN = SSN
		employee.Designation = Designation
		employee.Department = Department
	}
	defer database.Close()
	return employee
}

func GetEmployees() []Employee {
	var database *sql.DB
	database = GetConnection()

	var error error
	var rows *sql.Rows
	rows, error = database.Query("SELECT * FROM Employee ORDER BY Employeeid DESC")
	if error != nil {
		panic(error.Error())
	}
	var employee Employee
	employee = Employee{}

	var employees []Employee
	employees = []Employee{}
	for rows.Next() {
		var employeeId int
		var employeeName string
		var ssn string
		var designation string
		var department string
		error = rows.Scan(&employeeId, &employeeName, &ssn, &designation, &department)
		if error != nil {
			panic(error.Error())
		}
		employee.EmployeeId = employeeId
		employee.EmployeeName = employeeName
		employee.SSN = ssn
		employee.Designation = designation
		employee.Department = department
		employees = append(employees, employee)
	}

	defer database.Close()

	return employees
}

func InsertEmployee(employee Employee) {
	var database *sql.DB
	database = GetConnection()

	var error error
	var insert *sql.Stmt
	insert, error = database.Prepare("INSERT INTO EMPLOYEE(EmployeeName,SSN,Designation,Department) VALUES(?,?,?,?)")
	if error != nil {
		panic(error.Error())
	}
	insert.Exec(employee.EmployeeName, employee.SSN, employee.Designation, employee.Department)

	defer database.Close()

	//return Employee{}
}

func UpdateEmployee(employee Employee) {
	var database *sql.DB
	database = GetConnection()

	var error error
	var update *sql.Stmt
	update, error = database.Prepare("UPDATE EMPLOYEE SET EmployeeName=?, SSN=?,DESIGNATION=?,DEPARTMENT=? WHERE EmployeeId=?")
	if error != nil {
		panic(error.Error())
	}
	update.Exec(employee.EmployeeName, employee.SSN, employee.Designation, employee.Department, employee.EmployeeId)

	defer database.Close()

	//return Employee{}
}
func DeleteEmployee(employee Employee) {
	var database *sql.DB
	database = GetConnection()

	var error error
	var delete *sql.Stmt
	delete, error = database.Prepare("DELETE FROM Employee WHERE Employeeid=?")
	if error != nil {
		panic(error.Error())
	}
	delete.Exec(employee.EmployeeId)

	defer database.Close()

}
