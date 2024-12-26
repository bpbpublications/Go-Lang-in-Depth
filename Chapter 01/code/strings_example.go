package main

import "fmt"

func main() {

	hello_string := "Hello World"
	var string_var string
	string_var = "initialized later"
	fmt.Println("intiallizedString: ", hello_string)
	fmt.Println("declared and initialized later: ", string_var)
	var strings [5]string
	strings[0] = "George"
	strings[1] = "John"
	strings[2] = "Tom"
	strings[3] = "Vivan"
	strings[4] = "Marlin"
	fmt.Println("array of strings: ", strings)
}
