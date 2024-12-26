package main

import "fmt"

type Customer struct {
	Name    string
	ID      string
	city    string
	Pincode int
}

func main() {

	var customer Customer
	fmt.Println(customer)

	customer1 := Customer{"Andy Griffith", "CUST2341", "Atlanta", 30033}

	fmt.Println("Customer1: ", customer1)

	customer2 := Customer{Name: "Bill Smith", ID: "CUST6532", city: "Boston", Pincode: 102108}

	fmt.Println("Customer2: ", customer2)

	customer3 := Customer{Name: "Thomas Hardy"}
	fmt.Println("customer3: ", customer3)
}
