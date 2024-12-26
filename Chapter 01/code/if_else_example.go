package main

import (
	"fmt"
)

func main() {
	var str = "Adobe"
	var defaultStr = "Google"
	boolA := false
	if boolA {
		fmt.Println(str)
	} else {
		fmt.Println(defaultStr)
	}
}
