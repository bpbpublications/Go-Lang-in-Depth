package main

import (
	"fmt"
	"sync"
)

var singlock = &sync.Mutex{}

type singleton struct {
}

var singletonInstance *singleton

func getSingleton() *singleton {
	if singletonInstance == nil {
		singlock.Lock()
		defer singlock.Unlock()
		if singletonInstance == nil {
			fmt.Println("Creating singleton first time")
			singletonInstance = &singleton{}
		} else {
			fmt.Println("Singleton instance is already created.")
		}
	} else {
		fmt.Println("Single instance already created.")
	}

	return singletonInstance
}
func main() {

	for i := 0; i < 30; i++ {
		go getSingleton()
	}

	fmt.Scanln()
}
