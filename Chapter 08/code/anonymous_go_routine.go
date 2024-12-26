
package main

import (
	"fmt"
	"time"
)


func main() {

	fmt.Println("Starting the main function")


	go func() {

		fmt.Println("Calling the anonymous function")
	}()

	time.Sleep(1 * time.Second)
	fmt.Println("Leaving the main Function")
}
