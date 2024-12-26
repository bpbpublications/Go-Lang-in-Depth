package main

import (
	"fmt"
)

func identifyInterfaceType(interf interface{}) {
	fmt.Printf("Type of this interface is = %T, value of the interface is = %v\n", interf, interf)
}

func main() {
	varStr := "John Smith has come"
	identifyInterfaceType(varStr)
	varInt := 94
	identifyInterfaceType(varInt)
	varCustomer := struct {
		name string
	}{
		name: "Thomas Smith",
	}
	identifyInterfaceType(varCustomer)
}
