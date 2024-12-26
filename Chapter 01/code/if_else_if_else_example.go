package main

import (
	"fmt"
)

func main() {
	var str = "Adobe"
	var str1 = "Twitter"
	var str2 = "Amazon"
	var defaultStr = "Google"
	boolA := false

	boolB := true

	boolC := true
	if boolA {
		fmt.Println(str)
	} else if boolB {
		fmt.Println(str1)
	} else if boolC {
		fmt.Println(str2)
	} else {
		fmt.Println(defaultStr)
	}
}
