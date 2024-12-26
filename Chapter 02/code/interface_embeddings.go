package main

import "fmt"

type embeddedStruct struct {
	varSName string
}

func (eStruct embeddedStruct) display() string {
	return fmt.Sprintf("embedded Struct=%s", eStruct.varSName)
}

type Composite struct {
	embeddedStruct
	varName string
}

func main() {

	composite := Composite{
		embeddedStruct: embeddedStruct{
			varSName: "embedded User Defned Type",
		},
		varName: "Composite With Embedded Struct",
	}

	fmt.Printf("composite={struct Name: %s, composite Name: %s}\n", composite.varSName, composite.varName)

	fmt.Println("Struct name:", composite.embeddedStruct.varSName)

	fmt.Println("display:", composite.display())

	type displayer interface {
		display() string
	}

	var vDisplayer displayer = composite
	fmt.Println("displayer:", vDisplayer.display())
}
