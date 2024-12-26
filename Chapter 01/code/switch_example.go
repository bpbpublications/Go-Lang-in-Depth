package main

import (
	"fmt"
	"time"
)

func main() {
	timeNow := time.Now()
	var day int = timeNow.Day()

	switch day {
	case 6, 11, 17:
		fmt.Println("Buy Groceries")
	case 5, 16, 27:
		fmt.Println("Play Golf")
	case 3:
		fmt.Println("Go to Church")
	default:
		fmt.Println("Read a magazine")
	}
}
