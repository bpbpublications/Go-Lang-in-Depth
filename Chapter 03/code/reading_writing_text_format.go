package main

import (
	"fmt"
	"log"
	"os"
)

func WriteTextFile(path string) {

	fmt.Println("Writing a Text Format file in Go lang")

	text_file, error := os.Create(path)

	if error != nil {
		log.Fatalf("could not create a text file: %s", error)
	}

	defer text_file.Close()

	length, error := text_file.WriteString("This is to test writing a text formatted file" +
		" Code demonstrates writing and reading a text file" +
		" Go lang.")

	if error != nil {
		log.Fatalf("could not create a text file: %s", error)
	}

	fmt.Println("Text File Name is %s", text_file.Name())
	fmt.Println("Text File Length is %d bytes", length)
}

func FindTextFile(textFile string) {

	fmt.Println("Reading a text format  in Go lang")
	text_data, error := os.ReadFile(textFile)
	if error != nil {
		log.Panicf("cannot reading test from file: %s", error)
	}
	fmt.Println("Text File Name: %s", textFile)
	fmt.Println("Text DataSize: %d bytes", len(text_data))
	fmt.Println("Text Data: %s", string(text_data))

}

func main() {

	path := "input_test_format.txt"
	WriteTextFile(path)
	FindTextFile(path)
}
