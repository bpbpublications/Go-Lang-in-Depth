package main

import (
	"fmt"
	"sync"
)

var GEncounter = 0

func executor(waitGroup *sync.WaitGroup) {
	GEncounter = GEncounter + 1

	waitGroup.Done()
}
func main() {

	var varWaitGroup sync.WaitGroup

	for i := 0; i < 1000; i++ {
		varWaitGroup.Add(1)
		go executor(&varWaitGroup)
	}

	varWaitGroup.Wait()
	fmt.Println("Value of encounter", GEncounter)
}
