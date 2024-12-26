
package main

import "fmt"

func main() {


	var channel chan int
	fmt.Println(" channel's Value: ", channel)
	fmt.Printf("channel's Type: %T ", channel)


	channel1 := make(chan int)
	fmt.Println("\n channel1's Value: ", channel1)
	fmt.Printf("channel1's Type: %T ", channel1)
}
