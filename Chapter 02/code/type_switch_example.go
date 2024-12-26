package main

import (
	"fmt"
)

func identifyType(varInter interface{}) {
	switch varInter.(type) {
	case string:
		fmt.Printf("The type is a string and the value is %s\n", varInter.(string))
	case int:
		fmt.Printf("The type is int and the value is %d\n", varInter.(int))
	default:
		fmt.Printf("Ths is not known type\n")
	}
}
func main() {
	identifyType("This is a type check")
	identifyType(93)
	identifyType(231.44)
}
