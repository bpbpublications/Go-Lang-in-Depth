package main

import "fmt"

func main() {

	string_array := [5]string{"Tom", "George", "John", "Jim", "Bill"}

	fmt.Println("Elements of string array:")

	for i := 0; i < 4; i++ {
		fmt.Println(string_array[i])
	}

}
