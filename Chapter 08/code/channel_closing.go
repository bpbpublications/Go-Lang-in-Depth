
package main

import "fmt"

func newfunc(channel chan string) {

	for v := 0; v < 4; v++ {
		channel <- "running the channel"
	}
	close(channel)
}


func main() {


	channel := make(chan string)


	go newfunc(channel)


	for {
		result, message := <-channel
		if message == false {
			fmt.Println("Closing the channel ", message)
			break
		}
		fmt.Println("Opening the channel ", result, message)
	}
}
