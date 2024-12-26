package main

import "fmt"

func shareToChannel(varChannel chan int) {

	for i := 0; i < 7; i++ {
		varChannel <- i
	}

	close(varChannel)

}

func main() {

	varintChannel := make(chan int)

	go shareToChannel(varintChannel)

	for i := 0; i < 7; i++ {
		fmt.Println(<-varintChannel)
	}
}
