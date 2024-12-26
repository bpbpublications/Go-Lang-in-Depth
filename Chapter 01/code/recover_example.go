package main

import "fmt"

func callPanic() {

	if a := recover(); a != nil {
		fmt.Println("RECOVER", a)
	}
}

func enterInput(lady *string, resort *string) {

	callPanic()

	if lady == nil {
		panic("Error: Lady name cannot be nil")
	}

	if resort == nil {
		panic("Error: Resort name cannot be nil")
	}

	fmt.Printf("Lady Name: %s \n Resort: %s\n", *lady, *resort)
	fmt.Printf("enterInput completed")
}

func main() {

	Name_lady := "Annie"
	enterInput(&Name_lady, nil)
	fmt.Printf("main function completed")
}
