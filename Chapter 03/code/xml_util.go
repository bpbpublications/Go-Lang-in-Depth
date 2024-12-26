package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

type Customers struct {
	XMLName xml.Name `xml:"customers"`

	Customers []Customer `xml:"customer"`
}

type Customer struct {
	XMLName xml.Name `xml:"customer"`
	//Type    string   `xml:"type,attr"`
	Address string  `xml:"address"`
	Name    string  `xml:"name"`
	Company Company `xml:"company"`
}

type Company struct {
	XMLName  xml.Name `xml:"company"`
	Category string   `xml:"category,attr"`
	Name     string   `xml:"name"`
	Location string   `xml:"location"`
}

type XMLUtil struct {
}

func NewXMLUtil() *XMLUtil {
	return &XMLUtil{}
}

func (xmlUtil *XMLUtil) readXML(path string) {

	xmlFormatFile, error := os.Open("customers.xml")

	if error != nil {
		fmt.Println(error)
	}

	fmt.Println("customers.xml exists and opened")

	defer xmlFormatFile.Close()

	xmlBytes, _ := io.ReadAll(xmlFormatFile)

	var customers Customers

	xml.Unmarshal(xmlBytes, &customers)

	fmt.Println(customers)

	for i := 0; i < len(customers.Customers); i++ {
		fmt.Println("Customer Name is " + customers.Customers[i].Name)
		fmt.Println("Customer Address is " + customers.Customers[i].Address)
		fmt.Println("Customer's Company  " + customers.Customers[i].Company.Name)
		fmt.Println("Customer's Company Category is " + customers.Customers[i].Company.Category)
	}

}

func (xmlUtil *XMLUtil) writeXML(path string) {

	customer := Customer{
		Name:    "Jack Donner",
		Address: "Chicago",
		Company: Company{
			Name:     "Marigold",
			Location: "Denver",
			Category: "Silver"},
	}

	file, _ := xml.MarshalIndent(customer, "", " ")

	_ = os.WriteFile(path, file, 0644)

}

func main() {

	var xmlUtil *XMLUtil

	xmlUtil = NewXMLUtil()

	xmlUtil.readXML("customers.xml")

	xmlUtil.writeXML("customer.xml")

}
