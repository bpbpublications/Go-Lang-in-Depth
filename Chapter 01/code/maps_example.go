package main

import "fmt"

func main() {

	var integer_keymap map[int]int

	if integer_keymap == nil {

		fmt.Println("It is not initalized")
	} else {

		fmt.Println("It is initialized")
	}

	string_keymap := map[int]string{

		9:  "Bat",
		23: "Ball",
		14: "Cap",
		19: "Pad",
		4:  "Shoe",
	}

	fmt.Println("String keymap: ", string_keymap)
}
