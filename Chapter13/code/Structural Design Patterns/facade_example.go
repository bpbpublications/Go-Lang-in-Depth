package main

import "fmt"

type EmployeeFacade struct {
	employee *Employee
	address  *Address
	profile  *Profile
}

func newEmployeeFacade(empName string, line1 string, line2 string, ssn string) *EmployeeFacade {
	fmt.Println("Creating the employee Facade components")
	employeeFacade := &EmployeeFacade{
		employee: newEmployee(empName),
		address:  newAddress(line1, line2),
		profile:  newProfile(ssn),
	}
	fmt.Println("Employee Facade created")
	return employeeFacade
}

func (empFacade *EmployeeFacade) saveEmployeeBO() {
	empFacade.employee.saveEmployee()
	empFacade.address.saveAddress()
	empFacade.profile.saveProfile()

}

type Employee struct {
	name string
}

func newEmployee(empName string) *Employee {
	return &Employee{
		name: empName,
	}
}

func (emp *Employee) saveEmployee() {

	fmt.Println("Persisting the employee in the datastore")

}

type Address struct {
	line1 string
	line2 string
}

func newAddress(line1 string, line2 string) *Address {
	return &Address{
		line1: line1,
		line2: line2,
	}
}

func (address *Address) saveAddress() {

	fmt.Println("Persisting the address in the datastore")

}

type Profile struct {
	ssn string
}

func newProfile(ssn string) *Profile {
	return &Profile{ssn: ssn}
}

func (profile *Profile) saveProfile() {

	fmt.Println("Persisting theprofile in the datastore")

}

func main() {

	empFacade := newEmployeeFacade("Jay Smith", "200 South Blvd", "Boston, MA 01234", "23211111")

	empFacade.saveEmployeeBO()

}
