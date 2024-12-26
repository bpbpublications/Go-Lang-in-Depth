package main

import "fmt"

type numVar int
type strVar string
type boolVar bool

func main() {

	var intVar numVar = 98

	var phraseVar strVar = "checking string"

	var falseVar boolVar = false

	fmt.Printf("%v\n", intVar)

	fmt.Printf("%T\n", intVar)

	fmt.Printf("%v\n", phraseVar)

	fmt.Printf("%T\n", phraseVar)

	fmt.Printf("%v\n", falseVar)

	fmt.Printf("%T\n", falseVar)

}
