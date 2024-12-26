package main

import "fmt"

func main() {

	fibonacci_numbers := [10]int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}

	fmt.Println(fibonacci_numbers)

	var subset_fib []int

	subset_fib = fibonacci_numbers[2:6]

	fmt.Println(subset_fib)

}
