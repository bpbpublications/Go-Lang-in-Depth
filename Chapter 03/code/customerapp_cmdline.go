package main

import (
	"bufio"
	"fmt"
	"os"
)

type Customer struct {
	id        int
	name      string
	address   string
	contactno string
}

var customers []Customer

func authenticate() bool {
	var userid, password string
	fmt.Print("Enter Your userid: ")
	fmt.Scan(&userid)
	fmt.Print("Enter Your password: ")
	fmt.Scan(&password)

	if userid == "guest" && password == "guest" {
		fmt.Println("You are authenticated")
		return true
	} else {
		fmt.Println("Your input is wrong. Please enter again")
		return false
	}
}

func showCommandMenu() {
	fmt.Println("1. Display Customers")
	fmt.Println("2. Display Customer by ID")
	fmt.Println("3. Enter Customer Information")
	fmt.Println("4. Update Customer Information")
	fmt.Println("5. Delete Customer")
	fmt.Println("0. Exit")
	fmt.Print("Enter the menu choice id: ")
}

func CommandMenu() {
	var choice int
	showCommandMenu()
	for {
		fmt.Scan(&choice)
		if choice == 0 {
			fmt.Println("You entered 0 for Exiting the Program")
			break
		} else {
			switch choice {
			case 1:
				displayCustomerList(customers)
				fmt.Println("Input 'Enter' to go back to command menu....")
				bufio.NewReader(os.Stdin).ReadBytes('\n')
				showCommandMenu()
			case 2:
				var customerId int
				fmt.Println()
				fmt.Print("Input Customer ID :")
				fmt.Scan(&customerId)
				fmt.Println("\nDisplaying  Customer with ID ", customerId)
				displayCustomerById(customerId, customers)
				fmt.Println("Input 'Enter' to back to command menu....")
				bufio.NewReader(os.Stdin).ReadBytes('\n')
				showCommandMenu()
			case 3:
				fmt.Println("Add a new Customer")
				addNewCustomer(&customers)
				fmt.Println()
				fmt.Println("Input 'Enter' to back to menu....")
				bufio.NewReader(os.Stdin).ReadBytes('\n')
				showCommandMenu()
			case 4:
				var customerId int
				fmt.Print("Input Customer Id: ")
				fmt.Scan(&customerId)
				updateCustomerById(customerId, &customers)
				fmt.Println()
				fmt.Println("Imput 'Enter' to back to menu....")
				bufio.NewReader(os.Stdin).ReadBytes('\n')
				showCommandMenu()
			case 5:
				var customerId int
				fmt.Println()
				fmt.Print("Input Customer Id: ")
				fmt.Scan(&customerId)
				deleteCustomerById(customerId, &customers)
				fmt.Println()
				fmt.Println("input 'Enter' to back to menu....")
				bufio.NewReader(os.Stdin).ReadBytes('\n')
				showCommandMenu()
			default:
				fmt.Println("Input your choice again!")
				showCommandMenu()
			}
		}
	}
}

func displayCustomerList(customers []Customer) {

	for _, customer := range customers {
		fmt.Println("ID	:", customer.id)
		fmt.Println("Name	:", customer.name)
		fmt.Println("Address	:", customer.address)
		fmt.Println("Contact No	:", customer.contactno)
		fmt.Println()
	}
}

func displayCustomerById(id int, Customers []Customer) {
	for _, customer := range Customers {
		if customer.id == id {
			fmt.Println("ID	:", customer.id)
			fmt.Println("Name	:", customer.name)
			fmt.Println("Address	:", customer.address)
			fmt.Println("Contact No	:", customer.contactno)
			fmt.Println()
		}
	}
}

func addNewCustomer(customers *[]Customer) {
	var customerId int
	consoleReader := bufio.NewReader(os.Stdin)
	fmt.Println()
	fmt.Print("Enter Customer Id : ")
	fmt.Scan(&customerId)
	fmt.Print("Enter Customer Name : ")
	name, _ := consoleReader.ReadString('\n')

	fmt.Print("Enter Customer Address : ")
	address, _ := consoleReader.ReadString('\n')

	fmt.Print("Enter Customer Contact No : ")
	contactno, _ := consoleReader.ReadString('\n')

	Customer_ := Customer{customerId, name, address, contactno}

	*customers = append(*customers, Customer_)
}

func updateCustomerById(id int, customers *[]Customer) {
	var newContactno string
	fmt.Println()
	fmt.Print("Enter the New Contact No of the Customer: ")
	fmt.Scan(&newContactno)

	for i, customer := range *customers {
		if customer.id == id {
			(*customers)[i] = Customer{id, customer.name, customer.address, newContactno}
		}
	}
}

func deleteCustomerById(id int, customers *[]Customer) {
	for i, customer := range *customers {
		if customer.id == id {
			*customers = append((*customers)[:i], (*customers)[i+1:]...)
		}
	}
}

func CRMConsoleApp() {
	for !authenticate() {
		authenticate()
		break
	}

	customers = append(customers, Customer{id: 1, name: "John Smith", address: "California", contactno: "231456"})
	customers = append(customers, Customer{id: 2, name: "Brad Smith", address: "Boston", contactno: "43212344"})

	CommandMenu()
}

func main() {
	CRMConsoleApp()
}
