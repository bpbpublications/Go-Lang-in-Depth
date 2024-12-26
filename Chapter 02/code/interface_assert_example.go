package main

import (
	"fmt"
)

func assert(interf interface{}) {
	varS := interf.(string)
	fmt.Println(varS)
}
func main() {
	var varInterface interface{} = "checking interface"
	assert(varInterface)
}
