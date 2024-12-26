package main

import "fmt"

func main() {

	var arr [4]string

	arr[0] = "My"
	arr[1] = "Name"
	arr[2] = "Is"
	arr[3] = "Tom"

	fmt.Println(arr[0], arr[1], arr[2])
	fmt.Println(arr[3])

	fibonacci_numbers := [10]int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}

	fmt.Println(fibonacci_numbers)

}
