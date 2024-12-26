package main

import "fmt"

type Displayer interface {
	display()
}

func main() {
	var varDisplay Displayer
	if varDisplay == nil {
		fmt.Printf("varDispaly is nil. It is of type %T  and value is %v\n", varDisplay, varDisplay)
	}
}
