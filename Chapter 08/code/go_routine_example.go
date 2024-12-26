
package main

import (
	"fmt"
	"time"
    "strconv"
)

func show(message string) {
	for i := 0; i < 7; i++ {
		time.Sleep(1 * time.Second)
        fmt.Println(strconv.Itoa(i) + ":"+ message)
	}
}

func main() {


	go show("go routine")


	show("running a method")
}
