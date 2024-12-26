
package main

import "fmt"

func exfunc(channel chan int) {

	fmt.Println(123 + <-channel)
}
func main() {
	fmt.Println("Main method has started")

	channel := make(chan int)

	go exfunc(channel)
	channel <- 23
	fmt.Println("Main method has ended")
}
