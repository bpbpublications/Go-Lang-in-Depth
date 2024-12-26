package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Customer struct {
	Id      int
	Name    string
	Address string
	Company string
}

func main() {

	customer := Customer{}
	customer.Id = 7832
	customer.Name = "George Smith"
	customer.Address = "California"
	customer.Company = "Dell"

	customerData, _ := json.Marshal(customer)
	fmt.Println(string(customerData))

	customer1 := Customer{783, "Thomas Smith", "Boston", "AMD"}
	jsonData, _ := json.Marshal(customer1)
	fmt.Println(string(jsonData))

	custData1 := []byte(`{"Id":354,"Name":"John Smith","Addresss":"Denver","Company":"Jockey"}`)
	var customer2 Customer
	json.Unmarshal(custData1, &customer2)
	fmt.Println(customer2)

	var customer4 Customer
	json.Unmarshal(customerData, &customer4)
	fmt.Println(customer4)

	customer5 := Customer{
		Id:      8934,
		Name:    "Kerry Donner",
		Address: "Charlotte",
		Company: "Best Buy",
	}

	json_file, _ := json.MarshalIndent(customer5, "", " ")

	_ = os.WriteFile("customer.json", json_file, 0644)
}
