package main

import (
	"fmt"
	"time"
)

func execute(message string) {
	for i := 1; i <= 5; i++ {
		time.Sleep(50 * time.Millisecond)
		fmt.Println(message, ": ", i)
	}
}

func main() {

	fmt.Println("Go concurrency example")

	go execute("Go routine 1")
	go execute("Go routine 2")
	go execute("Go routine 3")

	time.Sleep(time.Second)

	fmt.Println("Example completed")
}
